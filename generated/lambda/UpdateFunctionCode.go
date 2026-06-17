package lambda

// UpdateFunctionCode is generated as a reference stub.
// Executable command wiring lives under cmd/lambda.go.
//
// Updates a Lambda function's code. If code signing is enabled for the function,
// the code package must be signed by a trusted publisher. For more information,
// see [Configuring code signing for Lambda].
//
// If the function's package type is Image , then you must specify the code package
// in ImageUri as the URI of a [container image] in the Amazon ECR registry.
//
// If the function's package type is Zip , then you must specify the deployment
// package as a [.zip file archive]. Enter the Amazon S3 bucket and key of the code .zip file
// location. You can also provide the function code inline using the ZipFile field.
//
// The code in the deployment package must be compatible with the target
// instruction set architecture of the function ( x86-64 or arm64 ).
//
// The function's code is locked when you publish a version. You can't modify the
// code of a published version, only the unpublished version.
//
// For a function defined as a container image, Lambda resolves the image tag to
// an image digest. In Amazon ECR, if you update the image tag to a new image,
// Lambda does not automatically update the function.
//
// [.zip file archive]: https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html#gettingstarted-package-zip
// [Configuring code signing for Lambda]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html
// [container image]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html
