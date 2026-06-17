package rekognition

// AssociateFaces is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Associates one or more faces with an existing UserID. Takes an array of FaceIds
// . Each FaceId that are present in the FaceIds list is associated with the
// provided UserID. The number of FaceIds that can be used as input in a single
// request is limited to 100.
//
// Note that the total number of faces that can be associated with a single UserID
// is also limited to 100. Once a UserID has 100 faces associated with it, no
// additional faces can be added. If more API calls are made after the limit is
// reached, a ServiceQuotaExceededException will result.
//
// The UserMatchThreshold parameter specifies the minimum user match confidence
// required for the face to be associated with a UserID that has at least one
// FaceID already associated. This ensures that the FaceIds are associated with
// the right UserID. The value ranges from 0-100 and default value is 75.
//
// If successful, an array of AssociatedFace objects containing the associated
// FaceIds is returned. If a given face is already associated with the given UserID
// , it will be ignored and will not be returned in the response. If a given face
// is already associated to a different UserID , isn't found in the collection,
// doesn’t meet the UserMatchThreshold , or there are already 100 faces associated
// with the UserID , it will be returned as part of an array of
// UnsuccessfulFaceAssociations.
//
// The UserStatus reflects the status of an operation which updates a UserID
// representation with a list of given faces. The UserStatus can be:
//
// - ACTIVE - All associations or disassociations of FaceID(s) for a UserID are
// complete.
//
// - CREATED - A UserID has been created, but has no FaceID(s) associated with
// it.
//
// - UPDATING - A UserID is being updated and there are current associations or
// disassociations of FaceID(s) taking place.
