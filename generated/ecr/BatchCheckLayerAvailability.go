package ecr

// BatchCheckLayerAvailability is generated as a reference stub.
// Executable command wiring lives under cmd/ecr.go.
//
// Checks the availability of one or more image layers in a repository.
//
// When an image is pushed to a repository, each image layer is checked to verify
// if it has been uploaded before. If it has been uploaded, then the image layer is
// skipped.
//
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
