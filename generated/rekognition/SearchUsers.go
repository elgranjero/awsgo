package rekognition

// SearchUsers is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Searches for UserIDs within a collection based on a FaceId or UserId . This API
// can be used to find the closest UserID (with a highest similarity) to associate
// a face. The request must be provided with either FaceId or UserId . The
// operation returns an array of UserID that match the FaceId or UserId , ordered
// by similarity score with the highest similarity first.
