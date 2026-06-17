package rekognition

// CreateUser is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Creates a new User within a collection specified by CollectionId . Takes UserId
// as a parameter, which is a user provided ID which should be unique within the
// collection. The provided UserId will alias the system generated UUID to make
// the UserId more user friendly.
//
// Uses a ClientToken , an idempotency token that ensures a call to CreateUser
// completes only once. If the value is not supplied, the AWS SDK generates an
// idempotency token for the requests. This prevents retries after a network error
// results from making multiple CreateUser calls.
