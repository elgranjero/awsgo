package connect

// ReleasePhoneNumber is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Releases a phone number previously claimed to an Amazon Connect instance or
// traffic distribution group. You can call this API only in the Amazon Web
// Services Region where the number was claimed.
//
// To release phone numbers from a traffic distribution group, use the
// ReleasePhoneNumber API, not the Amazon Connect admin website.
//
// After releasing a phone number, the phone number enters into a cooldown period
// for up to 180 days. It cannot be searched for or claimed again until the period
// has ended. If you accidentally release a phone number, contact Amazon Web
// Services Support.
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
