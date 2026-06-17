package groundstation

// CancelContact is generated as a reference stub.
// Executable command wiring lives under cmd/groundstation.go.
//
// Cancels or stops a contact with a specified contact ID based on its position in
// the [contact lifecycle].
//
// For contacts that:
//
// - Have yet to start, the contact will be cancelled.
//
// - Have started but have yet to finish, the contact will be stopped.
//
// [contact lifecycle]: https://docs.aws.amazon.com/ground-station/latest/ug/contacts.lifecycle.html
