package marketplacecatalog

// StartChangeSet is generated as a reference stub.
// Executable command wiring lives under cmd/marketplacecatalog.go.
//
// Allows you to request changes for your entities. Within a single ChangeSet , you
// can't start the same change type against the same entity multiple times.
// Additionally, when a ChangeSet is running, all the entities targeted by the
// different changes are locked until the change set has completed (either
// succeeded, cancelled, or failed). If you try to start a change set containing a
// change against an entity that is already locked, you will receive a
// ResourceInUseException error.
//
// For example, you can't start the ChangeSet described in the [example] later in this
// topic because it contains two changes to run the same change type ( AddRevisions
// ) against the same entity ( entity-id(at)1 ).
//
// For more information about working with change sets, see [Working with change sets]. For information
// about change types for single-AMI products, see [Working with single-AMI products]. Also, for more information
// about change types available for container-based products, see [Working with container products].
//
// To download "DetailsDocument" shapes, see [Python] and [Java] shapes on GitHub.
//
// [Java]: https://github.com/awslabs/aws-marketplace-catalog-api-shapes-for-java/tree/main
// [Working with single-AMI products]: https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/ami-products.html#working-with-single-AMI-products
// [Working with change sets]: https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/welcome.html#working-with-change-sets
// [Working with container products]: https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/container-products.html#working-with-container-products
// [example]: https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/API_StartChangeSet.html#API_StartChangeSet_Examples
// [Python]: https://github.com/awslabs/aws-marketplace-catalog-api-shapes-for-python
