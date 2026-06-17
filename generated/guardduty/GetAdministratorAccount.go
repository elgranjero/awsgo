package guardduty

// GetAdministratorAccount is generated as a reference stub.
// Executable command wiring lives under cmd/guardduty.go.
//
// Provides the details of the GuardDuty administrator account associated with the
// current GuardDuty member account.
//
// Based on the type of account that runs this API, the following list shows how
// the API behavior varies:
//
// - When the GuardDuty administrator account runs this API, it will return
// success ( HTTP 200 ) but no content.
//
// - When a member account runs this API, it will return the details of the
// GuardDuty administrator account that is associated with this calling member
// account.
//
// - When an individual account (not associated with an organization) runs this
// API, it will return success ( HTTP 200 ) but no content.
