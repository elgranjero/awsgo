package imagebuilder

// PutContainerRecipePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/imagebuilder.go.
//
// Applies a policy to a container image. We recommend that you call the RAM API
// CreateResourceShare
// (https://docs.aws.amazon.com//ram/latest/APIReference/API_CreateResourceShare.html)
// to share resources. If you call the Image Builder API PutContainerImagePolicy ,
// you must also call the RAM API PromoteResourceShareCreatedFromPolicy
// (https://docs.aws.amazon.com//ram/latest/APIReference/API_PromoteResourceShareCreatedFromPolicy.html)
// in order for the resource to be visible to all principals with whom the resource
// is shared.
