package ecr

// GetDownloadUrlForLayer is generated as a reference stub.
// Executable command wiring lives under cmd/ecr.go.
//
// Retrieves the pre-signed Amazon S3 download URL corresponding to an image
// layer. You can only get URLs for image layers that are referenced in an image.
//
// When an image is pulled, the GetDownloadUrlForLayer API is called once per
// image layer that is not already cached.
//
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
