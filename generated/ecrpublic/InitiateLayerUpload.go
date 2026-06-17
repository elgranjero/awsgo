package ecrpublic

// InitiateLayerUpload is generated as a reference stub.
// Executable command wiring lives under cmd/ecrpublic.go.
//
// Notifies Amazon ECR that you intend to upload an image layer.
//
// When an image is pushed, the InitiateLayerUpload API is called once for each
// image layer that hasn't already been uploaded. Whether an image layer uploads is
// determined by the BatchCheckLayerAvailability API action.
//
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
