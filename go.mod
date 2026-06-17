module github.com/shabarkin/aws-enumerator

go 1.24

require (
	github.com/aws/aws-sdk-go-v2 v1.42.0
	github.com/aws/aws-sdk-go-v2/config v1.32.25
	github.com/aws/aws-sdk-go-v2/service/acm v1.40.0
	github.com/aws/aws-sdk-go-v2/service/amplify v1.39.4
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.40.6
	github.com/aws/aws-sdk-go-v2/service/appmesh v1.36.4
	github.com/aws/aws-sdk-go-v2/service/appsync v1.54.4
	github.com/aws/aws-sdk-go-v2/service/athena v1.58.4
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.67.4
	github.com/aws/aws-sdk-go-v2/service/backup v1.57.6
	github.com/aws/aws-sdk-go-v2/service/batch v1.65.6
	github.com/aws/aws-sdk-go-v2/service/chime v1.42.4
	github.com/aws/aws-sdk-go-v2/service/cloud9 v1.34.6
	github.com/aws/aws-sdk-go-v2/service/clouddirectory v1.31.6
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.72.1
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.65.2
	github.com/aws/aws-sdk-go-v2/service/cloudhsm v1.30.5
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.35.4
	github.com/aws/aws-sdk-go-v2/service/cloudsearch v1.33.4
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.56.4
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.69.4
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.34.4
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.36.4
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.47.4
	github.com/aws/aws-sdk-go-v2/service/codestar v1.23.4
	github.com/aws/aws-sdk-go-v2/service/comprehend v1.41.6
	github.com/aws/aws-sdk-go-v2/service/datapipeline v1.31.5
	github.com/aws/aws-sdk-go-v2/service/datasync v1.59.7
	github.com/aws/aws-sdk-go-v2/service/dax v1.30.2
	github.com/aws/aws-sdk-go-v2/service/devicefarm v1.39.5
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.40.0
	github.com/aws/aws-sdk-go-v2/service/dlm v1.37.6
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.59.0
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.307.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.58.4
	github.com/aws/aws-sdk-go-v2/service/ecs v1.83.0
	github.com/aws/aws-sdk-go-v2/service/eks v1.86.0
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.54.3
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.35.4
	github.com/aws/aws-sdk-go-v2/service/elastictranscoder v1.33.0
	github.com/aws/aws-sdk-go-v2/service/firehose v1.44.0
	github.com/aws/aws-sdk-go-v2/service/fms v1.45.6
	github.com/aws/aws-sdk-go-v2/service/fsx v1.66.6
	github.com/aws/aws-sdk-go-v2/service/gamelift v1.55.4
	github.com/aws/aws-sdk-go-v2/service/globalaccelerator v1.36.6
	github.com/aws/aws-sdk-go-v2/service/glue v1.144.0
	github.com/aws/aws-sdk-go-v2/service/greengrass v1.33.4
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.79.3
	github.com/aws/aws-sdk-go-v2/service/health v1.38.5
	github.com/aws/aws-sdk-go-v2/service/iam v1.54.5
	github.com/aws/aws-sdk-go-v2/service/inspector v1.31.4
	github.com/aws/aws-sdk-go-v2/service/iot v1.75.4
	github.com/aws/aws-sdk-go-v2/service/iotanalytics v1.32.0
	github.com/aws/aws-sdk-go-v2/service/kafka v1.52.6
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.44.2
	github.com/aws/aws-sdk-go-v2/service/kinesisanalytics v1.31.4
	github.com/aws/aws-sdk-go-v2/service/kinesisvideo v1.34.5
	github.com/aws/aws-sdk-go-v2/service/kms v1.53.4
	github.com/aws/aws-sdk-go-v2/service/lambda v1.92.3
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.56.1
	github.com/aws/aws-sdk-go-v2/service/machinelearning v1.35.4
	github.com/aws/aws-sdk-go-v2/service/macie v1.19.2
	github.com/aws/aws-sdk-go-v2/service/mediaconnect v1.50.1
	github.com/aws/aws-sdk-go-v2/service/mediaconvert v1.93.1
	github.com/aws/aws-sdk-go-v2/service/medialive v1.99.0
	github.com/aws/aws-sdk-go-v2/service/mediapackage v1.40.1
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.30.3
	github.com/aws/aws-sdk-go-v2/service/mediatailor v1.59.6
	github.com/aws/aws-sdk-go-v2/service/mobile v1.21.3
	github.com/aws/aws-sdk-go-v2/service/mq v1.35.2
	github.com/aws/aws-sdk-go-v2/service/opsworks v1.31.0
	github.com/aws/aws-sdk-go-v2/service/organizations v1.51.10
	github.com/aws/aws-sdk-go-v2/service/pinpoint v1.40.3
	github.com/aws/aws-sdk-go-v2/service/polly v1.58.3
	github.com/aws/aws-sdk-go-v2/service/pricing v1.42.7
	github.com/aws/aws-sdk-go-v2/service/ram v1.37.3
	github.com/aws/aws-sdk-go-v2/service/rds v1.119.3
	github.com/aws/aws-sdk-go-v2/service/redshift v1.63.3
	github.com/aws/aws-sdk-go-v2/service/rekognition v1.52.3
	github.com/aws/aws-sdk-go-v2/service/robomaker v1.36.2
	github.com/aws/aws-sdk-go-v2/service/route53 v1.63.3
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.36.3
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.46.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.104.0
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.255.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.42.3
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.71.7
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.40.4
	github.com/aws/aws-sdk-go-v2/service/shield v1.35.3
	github.com/aws/aws-sdk-go-v2/service/signer v1.33.6
	github.com/aws/aws-sdk-go-v2/service/sms v1.29.0
	github.com/aws/aws-sdk-go-v2/service/snowball v1.37.3
	github.com/aws/aws-sdk-go-v2/service/sns v1.40.1
	github.com/aws/aws-sdk-go-v2/service/sqs v1.44.0
	github.com/aws/aws-sdk-go-v2/service/ssm v1.69.3
	github.com/aws/aws-sdk-go-v2/service/storagegateway v1.44.3
	github.com/aws/aws-sdk-go-v2/service/sts v1.43.3
	github.com/aws/aws-sdk-go-v2/service/support v1.32.0
	github.com/aws/aws-sdk-go-v2/service/transcribe v1.56.3
	github.com/aws/aws-sdk-go-v2/service/transfer v1.73.3
	github.com/aws/aws-sdk-go-v2/service/translate v1.34.6
	github.com/aws/aws-sdk-go-v2/service/waf v1.31.5
	github.com/aws/aws-sdk-go-v2/service/workdocs v1.31.3
	github.com/aws/aws-sdk-go-v2/service/worklink v1.23.2
	github.com/aws/aws-sdk-go-v2/service/workmail v1.37.6
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.70.0
	github.com/aws/aws-sdk-go-v2/service/xray v1.37.3
	github.com/fatih/color v1.10.0
	github.com/joho/godotenv v1.3.0
	github.com/wayneashleyberry/terminal-dimensions v1.0.0
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.7.13 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.19.24 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.18.29 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.29 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.29 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.4.30 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.9.22 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.12.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.29 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.19.29 // indirect
	github.com/aws/aws-sdk-go-v2/service/signin v1.2.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.31.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.36.6 // indirect
	github.com/aws/smithy-go v1.27.1 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	golang.org/x/sys v0.0.0-20200223170610-d5e6a3e2c0ae // indirect
)
