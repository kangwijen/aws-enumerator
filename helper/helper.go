package helper

import (
	"flag"
	"fmt"
	"os"

	"github.com/shabarkin/aws-enumerator/servicemaster"
	"github.com/shabarkin/aws-enumerator/servicestructs"
	"github.com/shabarkin/aws-enumerator/utils"
)

func changeSpeedForTime(speed string) int {
	switch speed {
	case "fast":
		return 2000
	case "slow":
		return 4000
	case "normal":
		return 3000
	default:
		fmt.Fprintf(os.Stderr, "%s %s\n", utils.Red("Error:"), utils.Yellow("speed must be fast, normal, or slow"))
		os.Exit(2)
		return 3000
	}
}

func SetEnumerationPipeline(services, speed, endpointURL *string) {
	utils.LoadEnv()

	if *endpointURL != "" {
		os.Setenv("AWS_ENDPOINT_URL", *endpointURL)
	}

	if servicemaster.CheckAWSCredentials() {
		servicemaster.ServiceCall(
			servicestructs.GetServices(),
			utils.ProcessServiceArgument(*services),
			changeSpeedForTime(*speed),
		)
	}
}

func ValidateCredFlags() {
	if *AWS_region == "" || *AWS_access_key_id == "" || *AWS_secret_access_key == "" {
		fmt.Fprintf(os.Stderr, "%s %s\n", utils.Red("Error:"), utils.Yellow("region, access-key-id, and secret-access-key are required"))
		fmt.Fprintf(os.Stderr, "%s\n", utils.Yellow("Run: ./aws-enumerator cred -h"))
		os.Exit(2)
	}
}

func DumpInfo(services_dump *string, print_apicalls *bool, filter *string, errors_dump *bool) {
	if *services_dump == "all" {
		for _, service := range utils.ServiceNames() {
			utils.AnalyseService(service, *print_apicalls, *filter, *errors_dump)
		}
		fmt.Println()
	} else {
		for _, service := range utils.ProcessServiceArgument(*services_dump) {
			utils.AnalyseService(service, *print_apicalls, *filter, *errors_dump)
		}
		fmt.Println()
	}
}

var Cloudrider_help string = `Usage:
aws-enumerator [command]

Available Commands:
  cred         Write AWS credentials to .env for later authentication
  enum         Enumerate AWS services with configured credentials
  dump         Display gathered enumeration results from enum-results/

Flags:
  -h           Help for aws-enumerator
`

var Cloudrider_cred_help string = `Usage:
aws-enumerator cred [flags]

Flags:
  -region              AWS region (env: AWS_REGION)
  -access-key-id       AWS access key ID (env: AWS_ACCESS_KEY_ID)
  -secret-access-key   AWS secret access key (env: AWS_SECRET_ACCESS_KEY)
  -session-token       AWS session token (env: AWS_SESSION_TOKEN)
  -endpoint-url        Custom AWS API endpoint (env: AWS_ENDPOINT_URL)

Legacy flags (-aws_region, -aws_access_key_id, etc.) are still accepted.

Example:
  ./aws-enumerator cred -region us-west-2 -access-key-id AKIA... -secret-access-key ...
  ./aws-enumerator cred -region us-east-1 -access-key-id AKIA... -secret-access-key ... -endpoint-url http://lab.example.com
`

var Cloudrider_enum_help string = `Usage:
aws-enumerator enum [flags]

Flags:
  -services            Comma-separated service list or all (default: all)
  -speed               fast, normal, or slow (default: normal)
  -endpoint-url        Override AWS_ENDPOINT_URL for this run

Environment:
  AWS_ENDPOINT_URL              Global custom endpoint (LocalStack, lab APIs)
  AWS_ENDPOINT_URL_<SERVICE>    Per-service override (for example AWS_ENDPOINT_URL_S3)

Example:
  ./aws-enumerator enum -services iam,sts,s3,ec2 -speed normal
  ./aws-enumerator enum -services all
  ./aws-enumerator enum -services iam,s3 -endpoint-url http://lab.example.com
`

var Cloudrider_dump_help string = `Usage:
aws-enumerator dump [flags]

Flags:
  -services     Service name, comma-separated list, or all (default: all)
  -filter       Filter API call names by prefix
  -print        Print matching API call payloads
  -errors       Read error output instead of successful results

Example:
  ./aws-enumerator dump -services iam
  ./aws-enumerator dump -services iam -filter List -print
  ./aws-enumerator dump -services all -filter Get -print -errors
`

// Command line flags:

var Cred *flag.FlagSet = flag.NewFlagSet("cred", flag.ExitOnError)

var (
	credRegion          string
	credAccessKeyID     string
	credSecretAccessKey string
	credSessionToken    string
	credEndpointURL     string
)

var AWS_region *string = &credRegion
var AWS_access_key_id *string = &credAccessKeyID
var AWS_secret_access_key *string = &credSecretAccessKey
var AWS_session_token *string = &credSessionToken
var AWS_endpoint_url *string = &credEndpointURL

var Enum *flag.FlagSet = flag.NewFlagSet("enum", flag.ExitOnError)
var Services_enum *string = Enum.String("services", "all", "Comma-separated services or all")
var Speed *string = Enum.String("speed", "normal", "Enumeration throttle: fast, normal, or slow")

var (
	enumEndpointURL string
)

var Enum_endpoint_url *string = &enumEndpointURL

var Dump *flag.FlagSet = flag.NewFlagSet("dump", flag.ExitOnError)
var Services_dump *string = Dump.String("services", "all", "Comma-separated services or all")
var Errors_dump *bool = Dump.Bool("errors", false, "Dump API errors instead of results")
var Print *bool = Dump.Bool("print", false, "Print API call payloads")
var Filter *string = Dump.String("filter", "", "Filter API calls by name prefix")

func init() {
	Cred.StringVar(&credRegion, "region", "", "AWS region (env: AWS_REGION)")
	Cred.StringVar(&credRegion, "aws_region", "", "Deprecated: use -region")
	Cred.StringVar(&credAccessKeyID, "access-key-id", "", "AWS access key ID (env: AWS_ACCESS_KEY_ID)")
	Cred.StringVar(&credAccessKeyID, "aws_access_key_id", "", "Deprecated: use -access-key-id")
	Cred.StringVar(&credSecretAccessKey, "secret-access-key", "", "AWS secret access key (env: AWS_SECRET_ACCESS_KEY)")
	Cred.StringVar(&credSecretAccessKey, "aws_secret_access_key", "", "Deprecated: use -secret-access-key")
	Cred.StringVar(&credSessionToken, "session-token", "", "AWS session token (env: AWS_SESSION_TOKEN)")
	Cred.StringVar(&credSessionToken, "aws_session_token", "", "Deprecated: use -session-token")
	Cred.StringVar(&credEndpointURL, "endpoint-url", "", "Custom AWS API endpoint (env: AWS_ENDPOINT_URL)")
	Cred.StringVar(&credEndpointURL, "aws_endpoint_url", "", "Deprecated: use -endpoint-url")

	Enum.StringVar(&enumEndpointURL, "endpoint-url", "", "Override AWS_ENDPOINT_URL for this run")
	Enum.StringVar(&enumEndpointURL, "aws_endpoint_url", "", "Deprecated: use -endpoint-url")
}
