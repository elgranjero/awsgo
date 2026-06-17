package voiceid

// OptOutSpeaker is generated as a reference stub.
// Executable command wiring lives under cmd/voiceid.go.
//
// Opts out a speaker from Voice ID. A speaker can be opted out regardless of
// whether or not they already exist in Voice ID. If they don't yet exist, a new
// speaker is created in an opted out state. If they already exist, their existing
// status is overridden and they are opted out. Enrollment and evaluation
// authentication requests are rejected for opted out speakers, and opted out
// speakers have no voice embeddings stored in Voice ID.
