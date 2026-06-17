package iotthingsgraph

// SearchThings is generated as a reference stub.
// Executable command wiring lives under cmd/iotthingsgraph.go.
//
// Searches for things associated with the specified entity. You can search by
// both device and device model.
//
// For example, if two different devices, camera1 and camera2, implement the
// camera device model, the user can associate thing1 to camera1 and thing2 to
// camera2. SearchThings(camera2) will return only thing2, but SearchThings(camera)
// will return both thing1 and thing2.
//
// This action searches for exact matches and doesn't perform partial text
// matching.
//
// Deprecated: since: 2022-08-30
