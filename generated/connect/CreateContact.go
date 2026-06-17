package connect

// CreateContact is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Only the VOICE, EMAIL, and TASK channels are supported.
//
// - For VOICE: The supported initiation method is TRANSFER . The contacts
// created with this initiation method have a subtype connect:ExternalAudio .
//
// - For EMAIL: The supported initiation methods are OUTBOUND , AGENT_REPLY , and
// FLOW .
//
// - For TASK: The supported initiation method is API . Contacts created with
// this API have a sub-type of connect:ExternalTask .
//
// Creates a new VOICE, EMAIL, or TASK contact.
//
// After a contact is created, you can move it to the desired state by using the
// InitiateAs parameter. While you can use API to create task contacts that are in
// the COMPLETED state, you must contact Amazon Web Services Support before using
// it for bulk import use cases. Bulk import causes your requests to be throttled
// or fail if your CreateContact limits aren't high enough.
