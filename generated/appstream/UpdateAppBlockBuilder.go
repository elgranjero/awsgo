package appstream

// UpdateAppBlockBuilder is generated as a reference stub.
// Executable command wiring lives under cmd/appstream.go.
//
// Updates an app block builder.
//
// If the app block builder is in the STARTING or STOPPING state, you can't update
// it. If the app block builder is in the RUNNING state, you can only update the
// DisplayName and Description. If the app block builder is in the STOPPED state,
// you can update any attribute except the Name.
