package acm

// PutAccountConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/acm.go.
//
// Adds or modifies account-level configurations in ACM.
//
// The supported configuration option is DaysBeforeExpiry . This option specifies
// the number of days prior to certificate expiration when ACM starts generating
// EventBridge events. ACM sends one event per day per certificate until the
// certificate expires. By default, accounts receive events starting 45 days before
// certificate expiration.
