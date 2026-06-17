package ec2

// ModifyDefaultCreditSpecification is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Modifies the default credit option for CPU usage of burstable performance
// instances. The default credit option is set at the account level per Amazon Web
// Services Region, and is specified per instance family. All new burstable
// performance instances in the account launch using the default credit option.
//
// ModifyDefaultCreditSpecification is an asynchronous operation, which works at
// an Amazon Web Services Region level and modifies the credit option for each
// Availability Zone. All zones in a Region are updated within five minutes. But if
// instances are launched during this operation, they might not get the new credit
// option until the zone is updated. To verify whether the update has occurred, you
// can call GetDefaultCreditSpecification and check DefaultCreditSpecification for
// updates.
//
// For more information, see [Burstable performance instances] in the Amazon EC2 User Guide.
//
// [Burstable performance instances]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html
