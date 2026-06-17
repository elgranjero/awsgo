package chime

// CreateMeetingDialOut is generated as a reference stub.
// Executable command wiring lives under cmd/chime.go.
//
// Uses the join token and call metadata in a meeting request (From number, To
// number, and so forth) to initiate an outbound call to a public switched
// telephone network (PSTN) and join them into a Chime meeting. Also ensures that
// the From number belongs to the customer.
//
// To play welcome audio or implement an interactive voice response (IVR), use the
// CreateSipMediaApplicationCall action with the corresponding SIP media
// application ID.
//
// This API is not available in a dedicated namespace.
