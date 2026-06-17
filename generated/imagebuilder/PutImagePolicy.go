package imagebuilder

// PutImagePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/imagebuilder.go.
//
// Applies a policy to an image. We recommend that you call the RAM API [CreateResourceShare] to share
// resources. If you call the Image Builder API PutImagePolicy , you must also call
// the RAM API [PromoteResourceShareCreatedFromPolicy]in order for the resource to be visible to all principals with whom
// the resource is shared.
//
// [PromoteResourceShareCreatedFromPolicy]: https://docs.aws.amazon.com/ram/latest/APIReference/API_PromoteResourceShareCreatedFromPolicy.html
// [CreateResourceShare]: https://docs.aws.amazon.com/ram/latest/APIReference/API_CreateResourceShare.html
