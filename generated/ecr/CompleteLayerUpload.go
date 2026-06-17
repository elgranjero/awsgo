package ecr

// CompleteLayerUpload is generated as a reference stub.
// Executable command wiring lives under cmd/ecr.go.
//
// Informs Amazon ECR that the image layer upload has completed for a specified
// registry, repository name, and upload ID. You can optionally provide a sha256
// digest of the image layer for data validation purposes.
//
// When an image is pushed, the CompleteLayerUpload API is called once per each
// new image layer to verify that the upload has completed.
//
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
