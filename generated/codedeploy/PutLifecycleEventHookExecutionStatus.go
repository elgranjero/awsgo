package codedeploy

// PutLifecycleEventHookExecutionStatus is generated as a reference stub.
// Executable command wiring lives under cmd/codedeploy.go.
//
// Sets the result of a Lambda validation function. The function validates
//
// lifecycle hooks during a deployment that uses the Lambda or Amazon ECS compute
// platform. For Lambda deployments, the available lifecycle hooks are
// BeforeAllowTraffic and AfterAllowTraffic . For Amazon ECS deployments, the
// available lifecycle hooks are BeforeInstall , AfterInstall ,
// AfterAllowTestTraffic , BeforeAllowTraffic , and AfterAllowTraffic . Lambda
// validation functions return Succeeded or Failed . For more information, see [AppSpec 'hooks' Section for an Lambda Deployment]
// and [AppSpec 'hooks' Section for an Amazon ECS Deployment].
//
// [AppSpec 'hooks' Section for an Amazon ECS Deployment]: https://docs.aws.amazon.com/codedeploy/latest/userguide/reference-appspec-file-structure-hooks.html#appspec-hooks-ecs
// [AppSpec 'hooks' Section for an Lambda Deployment]: https://docs.aws.amazon.com/codedeploy/latest/userguide/reference-appspec-file-structure-hooks.html#appspec-hooks-lambda
