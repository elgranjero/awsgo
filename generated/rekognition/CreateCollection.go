package rekognition

// CreateCollection is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Creates a collection in an AWS Region. You can add faces to the collection
// using the IndexFacesoperation.
//
// For example, you might create collections, one for each of your application
// users. A user can then index faces using the IndexFaces operation and persist
// results in a specific collection. Then, a user can search the collection for
// faces in the user-specific container.
//
// When you create a collection, it is associated with the latest version of the
// face model version.
//
// Collection names are case-sensitive.
//
// This operation requires permissions to perform the rekognition:CreateCollection
// action. If you want to tag your collection, you also require permission to
// perform the rekognition:TagResource operation.
