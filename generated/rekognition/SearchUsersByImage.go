package rekognition

// SearchUsersByImage is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Searches for UserIDs using a supplied image. It first detects the largest face
// in the image, and then searches a specified collection for matching UserIDs.
//
// The operation returns an array of UserIDs that match the face in the supplied
// image, ordered by similarity score with the highest similarity first. It also
// returns a bounding box for the face found in the input image.
//
// Information about faces detected in the supplied image, but not used for the
// search, is returned in an array of UnsearchedFace objects. If no valid face is
// detected in the image, the response will contain an empty UserMatches list and
// no SearchedFace object.
