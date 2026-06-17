package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemakerruntime"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// sagemakerruntimeCmd represents the sagemakerruntime command
var _sagemakerruntimeCmd = &cobra.Command{
	Use:   "sagemakerruntime",
	Short: "AWS sagemakerruntime CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := sagemakerruntime.NewFromConfig(cfg)
		if _sagemakerruntimeInvokeEndpoint {
			sagemakerruntime_InvokeEndpoint(cfg, client)
			return
		}
		if _sagemakerruntimeInvokeEndpointAsync {
			sagemakerruntime_InvokeEndpointAsync(cfg, client)
			return
		}
		if _sagemakerruntimeInvokeEndpointWithResponseStream {
			sagemakerruntime_InvokeEndpointWithResponseStream(cfg, client)
			return
		}

	},
}

var (
	_sagemakerruntimeInvokeEndpoint                   bool
	_sagemakerruntimeInvokeEndpointAsync              bool
	_sagemakerruntimeInvokeEndpointWithResponseStream bool

	_sagemakerruntimeAccept                   string
	_sagemakerruntimeBody                     string
	_sagemakerruntimeContentType              string
	_sagemakerruntimeCustomAttributes         string
	_sagemakerruntimeEnableExplanations       string
	_sagemakerruntimeEndpointName             string
	_sagemakerruntimeFilename                 string
	_sagemakerruntimeInferenceComponentName   string
	_sagemakerruntimeInferenceId              string
	_sagemakerruntimeInputLocation            string
	_sagemakerruntimeInvocationTimeoutSeconds string
	_sagemakerruntimeRequestTTLSeconds        string
	_sagemakerruntimeS3OutputPathExtension    string
	_sagemakerruntimeSessionId                string
	_sagemakerruntimeTargetContainerHostname  string
	_sagemakerruntimeTargetModel              string
	_sagemakerruntimeTargetVariant            string
)

// After you deploy a model into production using Amazon SageMaker AI hosting
// services, your client applications use this API to get inferences from the model
// hosted at the specified endpoint.
//
// For an overview of Amazon SageMaker AI, see [How It Works].
//
// Amazon SageMaker AI strips all POST headers except those supported by the API.
// Amazon SageMaker AI might add additional headers. You should not rely on the
// behavior of headers outside those enumerated in the request syntax.
//
// Calls to InvokeEndpoint are authenticated by using Amazon Web Services
// Signature Version 4. For information, see [Authenticating Requests (Amazon Web Services Signature Version 4)]in the Amazon S3 API Reference.
//
// A customer's model containers must respond to requests within 60 seconds. The
// model itself can have a maximum processing time of 60 seconds before responding
// to invocations. If your model is going to take 50-60 seconds of processing time,
// the SDK socket timeout should be set to be 70 seconds.
//
// Endpoints are scoped to an individual account, and are not public. The URL does
// not contain the account ID, but Amazon SageMaker AI determines the account ID
// from the authentication token that is supplied by the caller.
//
// [How It Works]: https://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works.html
// [Authenticating Requests (Amazon Web Services Signature Version 4)]: https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html
func sagemakerruntime_InvokeEndpoint(cfg aws.Config, client *sagemakerruntime.Client) {
	input := &sagemakerruntime.InvokeEndpointInput{
		// Body: []byte, // Required
		// EndpointName: *string, // Required
	}

	if len(_sagemakerruntimeBody) > 0 {
		if err := assignInputField(input, "Body", _sagemakerruntimeBody); err != nil {
			log.Errorf("invalid --body: %s", err.Error())
			return
		}
	}
	if len(_sagemakerruntimeEndpointName) > 0 {
		input.EndpointName = aws.String(_sagemakerruntimeEndpointName)
	}
	if len(_sagemakerruntimeAccept) > 0 {
		input.Accept = aws.String(_sagemakerruntimeAccept)
	}
	if len(_sagemakerruntimeContentType) > 0 {
		input.ContentType = aws.String(_sagemakerruntimeContentType)
	}
	if len(_sagemakerruntimeCustomAttributes) > 0 {
		input.CustomAttributes = aws.String(_sagemakerruntimeCustomAttributes)
	}
	if len(_sagemakerruntimeEnableExplanations) > 0 {
		input.EnableExplanations = aws.String(_sagemakerruntimeEnableExplanations)
	}
	if len(_sagemakerruntimeInferenceComponentName) > 0 {
		input.InferenceComponentName = aws.String(_sagemakerruntimeInferenceComponentName)
	}
	if len(_sagemakerruntimeInferenceId) > 0 {
		input.InferenceId = aws.String(_sagemakerruntimeInferenceId)
	}
	if len(_sagemakerruntimeSessionId) > 0 {
		input.SessionId = aws.String(_sagemakerruntimeSessionId)
	}
	if len(_sagemakerruntimeTargetContainerHostname) > 0 {
		input.TargetContainerHostname = aws.String(_sagemakerruntimeTargetContainerHostname)
	}
	if len(_sagemakerruntimeTargetModel) > 0 {
		input.TargetModel = aws.String(_sagemakerruntimeTargetModel)
	}
	if len(_sagemakerruntimeTargetVariant) > 0 {
		input.TargetVariant = aws.String(_sagemakerruntimeTargetVariant)
	}

	if resp, err := client.InvokeEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// After you deploy a model into production using Amazon SageMaker AI hosting
// services, your client applications use this API to get inferences from the model
// hosted at the specified endpoint in an asynchronous manner.
//
// Inference requests sent to this API are enqueued for asynchronous processing.
// The processing of the inference request may or may not complete before you
// receive a response from this API. The response from this API will not contain
// the result of the inference request but contain information about where you can
// locate it.
//
// Amazon SageMaker AI strips all POST headers except those supported by the API.
// Amazon SageMaker AI might add additional headers. You should not rely on the
// behavior of headers outside those enumerated in the request syntax.
//
// Calls to InvokeEndpointAsync are authenticated by using Amazon Web Services
// Signature Version 4. For information, see [Authenticating Requests (Amazon Web Services Signature Version 4)]in the Amazon S3 API Reference.
//
// [Authenticating Requests (Amazon Web Services Signature Version 4)]: https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html
func sagemakerruntime_InvokeEndpointAsync(cfg aws.Config, client *sagemakerruntime.Client) {
	input := &sagemakerruntime.InvokeEndpointAsyncInput{
		// EndpointName: *string, // Required
		// InputLocation: *string, // Required
	}

	if len(_sagemakerruntimeEndpointName) > 0 {
		input.EndpointName = aws.String(_sagemakerruntimeEndpointName)
	}
	if len(_sagemakerruntimeInputLocation) > 0 {
		input.InputLocation = aws.String(_sagemakerruntimeInputLocation)
	}
	if len(_sagemakerruntimeAccept) > 0 {
		input.Accept = aws.String(_sagemakerruntimeAccept)
	}
	if len(_sagemakerruntimeContentType) > 0 {
		input.ContentType = aws.String(_sagemakerruntimeContentType)
	}
	if len(_sagemakerruntimeCustomAttributes) > 0 {
		input.CustomAttributes = aws.String(_sagemakerruntimeCustomAttributes)
	}
	if len(_sagemakerruntimeFilename) > 0 {
		input.Filename = aws.String(_sagemakerruntimeFilename)
	}
	if len(_sagemakerruntimeInferenceId) > 0 {
		input.InferenceId = aws.String(_sagemakerruntimeInferenceId)
	}
	if len(_sagemakerruntimeInvocationTimeoutSeconds) > 0 {
		if err := assignInputField(input, "InvocationTimeoutSeconds", _sagemakerruntimeInvocationTimeoutSeconds); err != nil {
			log.Errorf("invalid --invocation-timeout-seconds: %s", err.Error())
			return
		}
	}
	if len(_sagemakerruntimeRequestTTLSeconds) > 0 {
		if err := assignInputField(input, "RequestTTLSeconds", _sagemakerruntimeRequestTTLSeconds); err != nil {
			log.Errorf("invalid --request-ttl-seconds: %s", err.Error())
			return
		}
	}
	if len(_sagemakerruntimeS3OutputPathExtension) > 0 {
		input.S3OutputPathExtension = aws.String(_sagemakerruntimeS3OutputPathExtension)
	}

	if resp, err := client.InvokeEndpointAsync(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Invokes a model at the specified endpoint to return the inference response as a
// stream. The inference stream provides the response payload incrementally as a
// series of parts. Before you can get an inference stream, you must have access to
// a model that's deployed using Amazon SageMaker AI hosting services, and the
// container for that model must support inference streaming.
//
// For more information that can help you use this API, see the following sections
// in the Amazon SageMaker AI Developer Guide:
//
// - For information about how to add streaming support to a model, see [How Containers Serve Requests].
//
// - For information about how to process the streaming response, see [Invoke real-time endpoints].
//
// Before you can use this operation, your IAM permissions must allow the
// sagemaker:InvokeEndpoint action. For more information about Amazon SageMaker AI
// actions for IAM policies, see [Actions, resources, and condition keys for Amazon SageMaker AI]in the IAM Service Authorization Reference.
//
// Amazon SageMaker AI strips all POST headers except those supported by the API.
// Amazon SageMaker AI might add additional headers. You should not rely on the
// behavior of headers outside those enumerated in the request syntax.
//
// Calls to InvokeEndpointWithResponseStream are authenticated by using Amazon Web
// Services Signature Version 4. For information, see [Authenticating Requests (Amazon Web Services Signature Version 4)]in the Amazon S3 API
// Reference.
//
// [How Containers Serve Requests]: https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-inference-code.html#your-algorithms-inference-code-how-containe-serves-requests
// [Invoke real-time endpoints]: https://docs.aws.amazon.com/sagemaker/latest/dg/realtime-endpoints-test-endpoints.html
// [Actions, resources, and condition keys for Amazon SageMaker AI]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazonsagemaker.html
// [Authenticating Requests (Amazon Web Services Signature Version 4)]: https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html
func sagemakerruntime_InvokeEndpointWithResponseStream(cfg aws.Config, client *sagemakerruntime.Client) {
	input := &sagemakerruntime.InvokeEndpointWithResponseStreamInput{
		// Body: []byte, // Required
		// EndpointName: *string, // Required
	}

	if len(_sagemakerruntimeBody) > 0 {
		if err := assignInputField(input, "Body", _sagemakerruntimeBody); err != nil {
			log.Errorf("invalid --body: %s", err.Error())
			return
		}
	}
	if len(_sagemakerruntimeEndpointName) > 0 {
		input.EndpointName = aws.String(_sagemakerruntimeEndpointName)
	}
	if len(_sagemakerruntimeAccept) > 0 {
		input.Accept = aws.String(_sagemakerruntimeAccept)
	}
	if len(_sagemakerruntimeContentType) > 0 {
		input.ContentType = aws.String(_sagemakerruntimeContentType)
	}
	if len(_sagemakerruntimeCustomAttributes) > 0 {
		input.CustomAttributes = aws.String(_sagemakerruntimeCustomAttributes)
	}
	if len(_sagemakerruntimeInferenceComponentName) > 0 {
		input.InferenceComponentName = aws.String(_sagemakerruntimeInferenceComponentName)
	}
	if len(_sagemakerruntimeInferenceId) > 0 {
		input.InferenceId = aws.String(_sagemakerruntimeInferenceId)
	}
	if len(_sagemakerruntimeSessionId) > 0 {
		input.SessionId = aws.String(_sagemakerruntimeSessionId)
	}
	if len(_sagemakerruntimeTargetContainerHostname) > 0 {
		input.TargetContainerHostname = aws.String(_sagemakerruntimeTargetContainerHostname)
	}
	if len(_sagemakerruntimeTargetVariant) > 0 {
		input.TargetVariant = aws.String(_sagemakerruntimeTargetVariant)
	}

	if resp, err := client.InvokeEndpointWithResponseStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_sagemakerruntimeCmd)
	_sagemakerruntimeCmd.Flags().SortFlags = false

	_sagemakerruntimeCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_sagemakerruntimeCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_sagemakerruntimeCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_sagemakerruntimeCmd.Flags().StringVarP(&_sagemakerruntimeAccept, "accept", "", "", "Accept")
	_sagemakerruntimeCmd.Flags().StringVarP(&_sagemakerruntimeBody, "body", "", "", "Body")
	_sagemakerruntimeCmd.Flags().StringVarP(&_sagemakerruntimeContentType, "content-type", "", "", "Content Type")
	_sagemakerruntimeCmd.Flags().StringVarP(&_sagemakerruntimeCustomAttributes, "custom-attributes", "", "", "Custom Attributes")
	_sagemakerruntimeCmd.Flags().StringVarP(&_sagemakerruntimeEnableExplanations, "enable-explanations", "", "", "Enable Explanations")
	_sagemakerruntimeCmd.Flags().StringVarP(&_sagemakerruntimeEndpointName, "endpoint-name", "", "", "Endpoint Name")
	_sagemakerruntimeCmd.Flags().StringVarP(&_sagemakerruntimeFilename, "filename", "", "", "Filename")
	_sagemakerruntimeCmd.Flags().StringVarP(&_sagemakerruntimeInferenceComponentName, "inference-component-name", "", "", "Inference Component Name")
	_sagemakerruntimeCmd.Flags().StringVarP(&_sagemakerruntimeInferenceId, "inference-id", "", "", "Inference ID")
	_sagemakerruntimeCmd.Flags().StringVarP(&_sagemakerruntimeInputLocation, "input-location", "", "", "Input Location")
	_sagemakerruntimeCmd.Flags().StringVarP(&_sagemakerruntimeInvocationTimeoutSeconds, "invocation-timeout-seconds", "", "", "Invocation Timeout Seconds")
	_sagemakerruntimeCmd.Flags().StringVarP(&_sagemakerruntimeRequestTTLSeconds, "request-ttl-seconds", "", "", "Request TTL Seconds")
	_sagemakerruntimeCmd.Flags().StringVarP(&_sagemakerruntimeS3OutputPathExtension, "s3-output-path-extension", "", "", "S3 Output Path Extension")
	_sagemakerruntimeCmd.Flags().StringVarP(&_sagemakerruntimeSessionId, "session-id", "", "", "Session ID")
	_sagemakerruntimeCmd.Flags().StringVarP(&_sagemakerruntimeTargetContainerHostname, "target-container-hostname", "", "", "Target Container Hostname")
	_sagemakerruntimeCmd.Flags().StringVarP(&_sagemakerruntimeTargetModel, "target-model", "", "", "Target Model")
	_sagemakerruntimeCmd.Flags().StringVarP(&_sagemakerruntimeTargetVariant, "target-variant", "", "", "Target Variant")

	_sagemakerruntimeCmd.Flags().BoolVarP(&_sagemakerruntimeInvokeEndpoint, "invoke-endpoint", "", false, "Invoke Endpoint")
	_sagemakerruntimeCmd.Flags().BoolVarP(&_sagemakerruntimeInvokeEndpointAsync, "invoke-endpoint-async", "", false, "Invoke Endpoint Async")
	_sagemakerruntimeCmd.Flags().BoolVarP(&_sagemakerruntimeInvokeEndpointWithResponseStream, "invoke-endpoint-with-response-stream", "", false, "Invoke Endpoint With Response Stream")

}
