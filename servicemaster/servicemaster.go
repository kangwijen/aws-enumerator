package servicemaster

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/shabarkin/aws-enumerator/utils"
)

type ServiceMaster struct {
	Svc     interface{}
	SvcName string

	ApiCalls           []map[string]interface{}
	json_result_struct map[string][]string
	json_error_struct  map[string][]string

	api_call_result_channel chan string
	api_call_error_channel  chan string

	result_counter int
	error_counter  int
}

func (svc *ServiceMaster) ServiceEnumerator() {
	defer wg.Done()

	svc.initialize()

	for i := 0; i < len(svc.ApiCalls); i++ {
		go svc.apicall_wrapper(i)
	}

	svc.control_node()
	svc.save_result_to_file()
}

func (svc *ServiceMaster) initialize() {
	svc.error_counter = 0
	svc.result_counter = 0

	delete(svc.json_error_struct, "errors")
	svc.json_error_struct = make(map[string][]string)

	delete(svc.json_result_struct, svc.SvcName)
	svc.json_result_struct = make(map[string][]string)

	svc.api_call_result_channel = make(chan string, len(svc.ApiCalls))
	svc.api_call_error_channel = make(chan string, len(svc.ApiCalls))
}

func (svc *ServiceMaster) control_node() {
	for {
		if svc.result_counter >= len(svc.ApiCalls) {
			close(svc.api_call_result_channel)
			close(svc.api_call_error_channel)
			fmt.Println(utils.Green("Message: "), utils.Yellow("Successful"), utils.Yellow(strings.ToUpper(svc.SvcName))+utils.Yellow(":"), utils.Green(svc.result_counter-svc.error_counter), utils.Yellow("/"), utils.Red(svc.result_counter))
			break
		}

		select {
		case msg := <-svc.api_call_result_channel:
			svc.json_result_struct[svc.SvcName] = append(svc.json_result_struct[svc.SvcName], msg)
			svc.result_counter++

		case err := <-svc.api_call_error_channel:
			svc.json_error_struct[svc.SvcName] = append(svc.json_error_struct[svc.SvcName], err)
			svc.result_counter++
			svc.error_counter++
		}
	}
}

func (svc *ServiceMaster) apicall_wrapper(it int) {
	defer func() {
		if r := recover(); r != nil {
			svc.api_call_error_channel <- utils.PackResponse(map[string]string{
				"panic": fmt.Sprintf("%v", r),
			})
		}
	}()

	apicallValue, ok := svc.ApiCalls[it]["apicall"].(string)
	if !ok {
		svc.api_call_error_channel <- utils.PackResponse(map[string]string{
			"apicall": "invalid apicall descriptor",
		})
		return
	}

	method := reflect.ValueOf(svc.Svc).MethodByName(apicallValue)
	if !method.IsValid() {
		svc.api_call_error_channel <- utils.PackResponse(map[string]string{
			apicallValue: "method not found on service client",
		})
		return
	}

	results := method.Call([]reflect.Value{
		reflect.ValueOf(context.TODO()),
		reflect.ValueOf(svc.ApiCalls[it]["input_obj"]),
	})

	response := results[0].Interface()
	if !results[1].IsNil() {
		callErr, ok := results[1].Interface().(error)
		if !ok {
			svc.api_call_error_channel <- utils.PackResponse(map[string]string{
				apicallValue: "non-error failure value returned",
			})
			return
		}
		svc.api_call_error_channel <- utils.PackResponse(map[string]string{apicallValue: callErr.Error()})
		return
	}

	svc.api_call_result_channel <- utils.PackResponse(map[string]interface{}{apicallValue: response})
}

func (svc *ServiceMaster) save_result_to_file() {
	if err := os.MkdirAll(utils.FILEPATH, 0755); err != nil {
		log.Fatalln(utils.Red(err))
	}

	fileResults := utils.PackResponse(svc.json_result_struct)
	resultPath := filepath.Join(utils.FILEPATH, svc.SvcName+".json")
	if err := os.WriteFile(resultPath, []byte(fileResults), 0644); err != nil {
		log.Fatalln(utils.Red(err))
	}

	if err := os.MkdirAll(utils.ERROR_FILEPATH, 0755); err != nil {
		log.Fatalln(utils.Red(err))
	}

	fileErrors := utils.PackResponse(svc.json_error_struct)
	errorPath := filepath.Join(utils.ERROR_FILEPATH, svc.SvcName+"_errors.json")
	if err := os.WriteFile(errorPath, []byte(fileErrors), 0644); err != nil {
		log.Fatalln(utils.Red(err))
	}
}

func CheckAWSCredentials() bool {
	cfg, err := utils.LoadAWSConfig(context.TODO())
	if err != nil {
		fmt.Println(utils.Red("Error:"), utils.Yellow("Unable to load SDK config,"))
		fmt.Println(utils.Green("Fix:"), utils.Yellow("Provide AWS credentials via .env, environment variables, or ~/.aws/credentials"))
		fmt.Println(utils.Red("Trace:"), utils.Yellow(err))
		os.Exit(1)
	}

	stsSvc := sts.NewFromConfig(cfg)
	_, awsErr := stsSvc.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if awsErr != nil {
		fmt.Println(utils.Red("Error:"), utils.Yellow("AWS Credentials are not valid"))
		fmt.Println(utils.Green("Fix:"), utils.Yellow("Provide AWS Credentials, use `./aws-enumerator cred -h` command"))
		fmt.Println(utils.Red("Trace:"), utils.Yellow(awsErr))
		os.Exit(1)
	}
	return true
}

func sleep_delay(i, speed int) {
	if (i+1)%5 == 0 {
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}

var wg sync.WaitGroup

func ServiceCall(AllAWSServices []ServiceMaster, wanted_services []string, speed int) {
	start := time.Now()
	if utils.Find(wanted_services, "all") {
		for i := range AllAWSServices {
			wg.Add(1)
			go AllAWSServices[i].ServiceEnumerator()
			sleep_delay(i, speed)
		}
		wg.Wait()
	} else {
		for aws_i := range AllAWSServices {
			for str_i := range wanted_services {
				if AllAWSServices[aws_i].SvcName == wanted_services[str_i] {
					wg.Add(1)
					go AllAWSServices[aws_i].ServiceEnumerator()
					break
				}
			}
		}
		wg.Wait()
	}
	fmt.Println(utils.Green("Time:"), time.Since(start))
}
