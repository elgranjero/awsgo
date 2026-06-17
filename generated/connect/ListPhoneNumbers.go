package connect

// ListPhoneNumbers is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Provides information about the phone numbers for the specified Amazon Connect
// instance.
//
// For more information about phone numbers, see [Set Up Phone Numbers for Your Contact Center] in the Amazon Connect
// Administrator Guide.
//
// - We recommend using [ListPhoneNumbersV2]to return phone number types. ListPhoneNumbers doesn't
// support number types UIFN , SHARED , THIRD_PARTY_TF , and THIRD_PARTY_DID .
// While it returns numbers of those types, it incorrectly lists them as
// TOLL_FREE or DID .
//
// - The phone number Arn value that is returned from each of the items in the [PhoneNumberSummaryList]
// cannot be used to tag phone number resources. It will fail with a
// ResourceNotFoundException . Instead, use the [ListPhoneNumbersV2]API. It returns the new phone
// number ARN that can be used to tag phone number resources.
//
// [ListPhoneNumbersV2]: https://docs.aws.amazon.com/connect/latest/APIReference/API_ListPhoneNumbersV2.html
// [Set Up Phone Numbers for Your Contact Center]: https://docs.aws.amazon.com/connect/latest/adminguide/contact-center-phone-number.html
// [PhoneNumberSummaryList]: https://docs.aws.amazon.com/connect/latest/APIReference/API_ListPhoneNumbers.html#connect-ListPhoneNumbers-response-PhoneNumberSummaryList
