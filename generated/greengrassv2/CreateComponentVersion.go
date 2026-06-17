package greengrassv2

// CreateComponentVersion is generated as a reference stub.
// Executable command wiring lives under cmd/greengrassv2.go.
//
// Creates a component. Components are software that run on Greengrass core
// devices. After you develop and test a component on your core device, you can use
// this operation to upload your component to IoT Greengrass. Then, you can deploy
// the component to other core devices.
//
// You can use this operation to do the following:
//
// - Create components from recipes
//
// Create a component from a recipe, which is a file that defines the component's
//
// metadata, parameters, dependencies, lifecycle, artifacts, and platform
// capability. For more information, see [IoT Greengrass component recipe reference]in the IoT Greengrass V2 Developer
// Guide.
//
// To create a component from a recipe, specify inlineRecipe when you call this
//
// operation.
//
// - Create components from Lambda functions
//
// Create a component from an Lambda function that runs on IoT Greengrass. This
//
// creates a recipe and artifacts from the Lambda function's deployment package.
// You can use this operation to migrate Lambda functions from IoT Greengrass V1 to
// IoT Greengrass V2.
//
// This function accepts Lambda functions in all supported versions of Python,
//
// Node.js, and Java runtimes. IoT Greengrass doesn't apply any additional
// restrictions on deprecated Lambda runtime versions.
//
// To create a component from a Lambda function, specify lambdaFunction when you
//
// call this operation.
//
// IoT Greengrass currently supports Lambda functions on only Linux core devices.
//
// [IoT Greengrass component recipe reference]: https://docs.aws.amazon.com/greengrass/v2/developerguide/component-recipe-reference.html
