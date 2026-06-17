package ecrpublic

// CompleteLayerUpload is generated as a reference stub.
// Executable command wiring lives under cmd/ecrpublic.go.
//
// Informs Amazon ECR that the image layer upload is complete for a specified
// public registry, repository name, and upload ID. You can optionally provide a
// sha256 digest of the image layer for data validation purposes.
//
// When an image is pushed, the CompleteLayerUpload API is called once for each
// new image layer to verify that the upload is complete.
//
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
