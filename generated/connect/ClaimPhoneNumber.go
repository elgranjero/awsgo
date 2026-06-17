package connect

// ClaimPhoneNumber is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Claims an available phone number to your Amazon Connect instance or traffic
// distribution group. You can call this API only in the same Amazon Web Services
// Region where the Amazon Connect instance or traffic distribution group was
// created.
//
// For more information about how to use this operation, see [Claim a phone number in your country] and [Claim phone numbers to traffic distribution groups] in the Amazon
// Connect Administrator Guide.
//
// You can call the [SearchAvailablePhoneNumbers] API for available phone numbers that you can claim. Call the [DescribePhoneNumber]
// API to verify the status of a previous [ClaimPhoneNumber]operation.
//
// If you plan to claim and release numbers frequently, contact us for a service
// quota exception. Otherwise, it is possible you will be blocked from claiming and
// releasing any more numbers until up to 180 days past the oldest number released
// has expired.
//
// By default you can claim and release up to 200% of your maximum number of
// active phone numbers. If you claim and release phone numbers using the UI or API
// during a rolling 180 day cycle that exceeds 200% of your phone number service
// level quota, you will be blocked from claiming any more numbers until 180 days
// past the oldest number released has expired.
//
// For example, if you already have 99 claimed numbers and a service level quota
// of 99 phone numbers, and in any 180 day period you release 99, claim 99, and
// then release 99, you will have exceeded the 200% limit. At that point you are
// blocked from claiming any more numbers until you open an Amazon Web Services
// support ticket.
//
// [Claim phone numbers to traffic distribution groups]: https://docs.aws.amazon.com/connect/latest/adminguide/claim-phone-numbers-traffic-distribution-groups.html
// [Claim a phone number in your country]: https://docs.aws.amazon.com/connect/latest/adminguide/claim-phone-number.html
// [SearchAvailablePhoneNumbers]: https://docs.aws.amazon.com/connect/latest/APIReference/API_SearchAvailablePhoneNumbers.html
// [DescribePhoneNumber]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribePhoneNumber.html
// [ClaimPhoneNumber]: https://docs.aws.amazon.com/connect/latest/APIReference/API_ClaimPhoneNumber.html
