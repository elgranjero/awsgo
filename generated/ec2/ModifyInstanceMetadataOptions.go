package ec2

// ModifyInstanceMetadataOptions is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Modify the instance metadata parameters on a running or stopped instance. When
// you modify the parameters on a stopped instance, they are applied when the
// instance is started. When you modify the parameters on a running instance, the
// API responds with a state of “pending”. After the parameter modifications are
// successfully applied to the instance, the state of the modifications changes
// from “pending” to “applied” in subsequent describe-instances API calls. For more
// information, see [Instance metadata and user data]in the Amazon EC2 User Guide.
//
// [Instance metadata and user data]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html
