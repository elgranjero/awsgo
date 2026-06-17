package rekognition

// DeleteProjectVersion is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Deletes a Rekognition project model or project version, like a Amazon
// Rekognition Custom Labels model or a custom adapter.
//
// You can't delete a project version if it is running or if it is training. To
// check the status of a project version, use the Status field returned from DescribeProjectVersions. To
// stop a project version call StopProjectVersion. If the project version is training, wait until it
// finishes.
//
// This operation requires permissions to perform the
// rekognition:DeleteProjectVersion action.
