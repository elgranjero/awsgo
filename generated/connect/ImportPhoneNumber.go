package connect

// ImportPhoneNumber is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Imports a claimed phone number from an external service, such as Amazon Web
// Services End User Messaging, into an Amazon Connect instance. You can call this
// API only in the same Amazon Web Services Region where the Amazon Connect
// instance was created.
//
// Call the [DescribePhoneNumber] API to verify the status of a previous ImportPhoneNumber operation.
//
// If you plan to claim or import numbers and then release numbers frequently,
// contact us for a service quota exception. Otherwise, it is possible you will be
// blocked from claiming and releasing any more numbers until up to 180 days past
// the oldest number released has expired.
//
// By default you can claim or import and then release up to 200% of your maximum
// number of active phone numbers. If you claim or import and then release phone
// numbers using the UI or API during a rolling 180 day cycle that exceeds 200% of
// your phone number service level quota, you will be blocked from claiming or
// importing any more numbers until 180 days past the oldest number released has
// expired.
//
// For example, if you already have 99 claimed or imported numbers and a service
// level quota of 99 phone numbers, and in any 180 day period you release 99, claim
// 99, and then release 99, you will have exceeded the 200% limit. At that point
// you are blocked from claiming any more numbers until you open an Amazon Web
// Services Support ticket.
//
// [DescribePhoneNumber]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribePhoneNumber.html
