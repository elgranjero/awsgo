package lambda

// CreateFunction is generated as a reference stub.
// Executable command wiring lives under cmd/lambda.go.
//
// Creates a Lambda function. To create a function, you need a [deployment package] and an [execution role]. The
// deployment package is a .zip file archive or container image that contains your
// function code. The execution role grants the function permission to use Amazon
// Web Services services, such as Amazon CloudWatch Logs for log streaming and
// X-Ray for request tracing.
//
// If the deployment package is a [container image], then you set the package type to Image . For a
// container image, the code property must include the URI of a container image in
// the Amazon ECR registry. You do not need to specify the handler and runtime
// properties.
//
// If the deployment package is a [.zip file archive], then you set the package type to Zip . For a
// .zip file archive, the code property specifies the location of the .zip file.
// You must also specify the handler and runtime properties. The code in the
// deployment package must be compatible with the target instruction set
// architecture of the function ( x86-64 or arm64 ). If you do not specify the
// architecture, then the default value is x86-64 .
//
// When you create a function, Lambda provisions an instance of the function and
// its supporting resources. If your function connects to a VPC, this process can
// take a minute or so. During this time, you can't invoke or modify the function.
// The State , StateReason , and StateReasonCode fields in the response from GetFunctionConfiguration
// indicate when the function is ready to invoke. For more information, see [Lambda function states].
//
// A function has an unpublished version, and can have published versions and
// aliases. The unpublished version changes when you update your function's code
// and configuration. A published version is a snapshot of your function code and
// configuration that can't be changed. An alias is a named resource that maps to a
// version, and can be changed to map to a different version. Use the Publish
// parameter to create version 1 of your function from its initial configuration.
//
// The other parameters let you configure version-specific and function-level
// settings. You can modify version-specific settings later with UpdateFunctionConfiguration. Function-level
// settings apply to both the unpublished and published versions of the function,
// and include tags (TagResource ) and per-function concurrency limits (PutFunctionConcurrency ).
//
// You can use code signing if your deployment package is a .zip file archive. To
// enable code signing for this function, specify the ARN of a code-signing
// configuration. When a user attempts to deploy a code package with UpdateFunctionCode, Lambda
// checks that the code package has a valid signature from a trusted publisher. The
// code-signing configuration includes set of signing profiles, which define the
// trusted publishers for this function.
//
// If another Amazon Web Services account or an Amazon Web Services service
// invokes your function, use AddPermissionto grant permission by creating a resource-based
// Identity and Access Management (IAM) policy. You can grant permissions at the
// function level, on a version, or on an alias.
//
// To invoke your function directly, use Invoke. To invoke your function in response to
// events in other Amazon Web Services services, create an event source mapping (CreateEventSourceMapping
// ), or configure a function trigger in the other service. For more information,
// see [Invoking Lambda functions].
//
// [Invoking Lambda functions]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-invocation.html
// [Lambda function states]: https://docs.aws.amazon.com/lambda/latest/dg/functions-states.html
// [.zip file archive]: https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html#gettingstarted-package-zip
// [execution role]: https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html#lambda-intro-execution-role
// [container image]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html
// [deployment package]: https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html
