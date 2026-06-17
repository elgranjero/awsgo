package rekognition

// DisassociateFaces is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Removes the association between a Face supplied in an array of FaceIds and the
// User. If the User is not present already, then a ResourceNotFound exception is
// thrown. If successful, an array of faces that are disassociated from the User is
// returned. If a given face is already disassociated from the given UserID, it
// will be ignored and not be returned in the response. If a given face is already
// associated with a different User or not found in the collection it will be
// returned as part of UnsuccessfulDisassociations . You can remove 1 - 100 face
// IDs from a user at one time.
