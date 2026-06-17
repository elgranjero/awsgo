package rekognition

// DeleteProject is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Deletes a Amazon Rekognition project. To delete a project you must first delete
// all models or adapters associated with the project. To delete a model or
// adapter, see DeleteProjectVersion.
//
// DeleteProject is an asynchronous operation. To check if the project is deleted,
// call DescribeProjects. The project is deleted when the project no longer appears in the
// response. Be aware that deleting a given project will also delete any
// ProjectPolicies associated with that project.
//
// This operation requires permissions to perform the rekognition:DeleteProject
// action.
