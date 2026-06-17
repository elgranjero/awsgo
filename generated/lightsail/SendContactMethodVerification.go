package lightsail

// SendContactMethodVerification is generated as a reference stub.
// Executable command wiring lives under cmd/lightsail.go.
//
// Sends a verification request to an email contact method to ensure it's owned by
// the requester. SMS contact methods don't need to be verified.
//
// A contact method is used to send you notifications about your Amazon Lightsail
// resources. You can add one email address and one mobile phone number contact
// method in each Amazon Web Services Region. However, SMS text messaging is not
// supported in some Amazon Web Services Regions, and SMS text messages cannot be
// sent to some countries/regions. For more information, see [Notifications in Amazon Lightsail].
//
// A verification request is sent to the contact method when you initially create
// it. Use this action to send another verification request if a previous
// verification request was deleted, or has expired.
//
// Notifications are not sent to an email contact method until after it is
// verified, and confirmed as valid.
//
// [Notifications in Amazon Lightsail]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-notifications
