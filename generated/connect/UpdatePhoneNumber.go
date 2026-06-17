package connect

// UpdatePhoneNumber is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Updates your claimed phone number from its current Amazon Connect instance or
// traffic distribution group to another Amazon Connect instance or traffic
// distribution group in the same Amazon Web Services Region.
//
// After using this API, you must verify that the phone number is attached to the
// correct flow in the target instance or traffic distribution group. You need to
// do this because the API switches only the phone number to a new instance or
// traffic distribution group. It doesn't migrate the flow configuration of the
// phone number, too.
//
// You can call [DescribePhoneNumber] API to verify the status of a previous [UpdatePhoneNumber] operation.
//
// [UpdatePhoneNumber]: https://docs.aws.amazon.com/connect/latest/APIReference/API_UpdatePhoneNumber.html
// [DescribePhoneNumber]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribePhoneNumber.html
