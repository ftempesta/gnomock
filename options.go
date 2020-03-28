package localstack

// Option is an optional configuration of this Gnomock preset. Use available
// Options to configure the container
type Option func(*options)

// Service represents an AWS service that can be setup using localstack
type Service string

// These services are available in this Preset
const (
	APIGateway       Service = "apigateway"
	CloudFormation   Service = "cloudformation"
	CloudWatch       Service = "cloudwatch"
	CloudWatchLogs   Service = "logs"
	CloudWatchEvents Service = "events"
	DynamoDB         Service = "dynamodb"
	DynamoDBStreams  Service = "dynamodbstreams"
	EC2              Service = "ec2"
	ES               Service = "es"
	Firehose         Service = "firehose"
	IAM              Service = "iam"
	Kinesis          Service = "kinesis"
	KMS              Service = "kms"
	Lambda           Service = "lambda"
	Redshift         Service = "redshift"
	Route53          Service = "route53"
	S3               Service = "s3"
	SecretsManager   Service = "secretsmanager"
	SES              Service = "ses"
	SNS              Service = "sns"
	SQS              Service = "sqs"
	SSM              Service = "ssm"
	STS              Service = "sts"
	StepFunctions    Service = "stepfunctions"
)

// WithServices selects localstack services to spin up. It is OK to not select
// any services, but in such case the container will be useless
func WithServices(services ...Service) Option {
	return func(o *options) {
		for _, service := range services {
			o.services = append(o.services, string(service))
		}
	}
}

type options struct {
	services []string
}

func buildConfig(opts ...Option) *options {
	config := &options{}

	for _, opt := range opts {
		opt(config)
	}

	return config
}
