package pinpointsmsvoicev2

// CreatePool is generated as a reference stub.
// Executable command wiring lives under cmd/pinpointsmsvoicev2.go.
//
// Creates a new pool and associates the specified origination identity to the
// pool. A pool can include one or more phone numbers and SenderIds that are
// associated with your Amazon Web Services account.
//
// The new pool inherits its configuration from the specified origination
// identity. This includes keywords, message type, opt-out list, two-way
// configuration, and self-managed opt-out configuration. Deletion protection isn't
// inherited from the origination identity and defaults to false.
//
// If the origination identity is a phone number and is already associated with
// another pool, an error is returned. A sender ID can be associated with multiple
// pools.
