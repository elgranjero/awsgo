package ecrpublic

// PutImage is generated as a reference stub.
// Executable command wiring lives under cmd/ecrpublic.go.
//
// Creates or updates the image manifest and tags that are associated with an
// image.
//
// When an image is pushed and all new image layers have been uploaded, the
// PutImage API is called once to create or update the image manifest and the tags
// that are associated with the image.
//
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
