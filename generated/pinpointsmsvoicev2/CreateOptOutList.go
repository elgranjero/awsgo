package pinpointsmsvoicev2

// CreateOptOutList is generated as a reference stub.
// Executable command wiring lives under cmd/pinpointsmsvoicev2.go.
//
// Creates a new opt-out list.
//
// If the opt-out list name already exists, an error is returned.
//
// An opt-out list is a list of phone numbers that are opted out, meaning you
// can't send SMS or voice messages to them. If end user replies with the keyword
// "STOP," an entry for the phone number is added to the opt-out list. In addition
// to STOP, your recipients can use any supported opt-out keyword, such as CANCEL
// or OPTOUT. For a list of supported opt-out keywords, see [SMS opt out]in the End User
// Messaging SMS User Guide.
//
// [SMS opt out]: https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-manage.html#channels-sms-manage-optout
