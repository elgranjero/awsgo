package resourceexplorer2

// CreateIndex is generated as a reference stub.
// Executable command wiring lives under cmd/resourceexplorer2.go.
//
// Turns on Amazon Web Services Resource Explorer in the Amazon Web Services
// Region in which you called this operation by creating an index. Resource
// Explorer begins discovering the resources in this Region and stores the details
// about the resources in the index so that they can be queried by using the Search
// operation. You can create only one index in a Region.
//
// This operation creates only a local index. To promote the local index in one
// Amazon Web Services Region into the aggregator index for the Amazon Web Services
// account, use the UpdateIndexTypeoperation. For more information, see [Turning on cross-Region search by creating an aggregator index] in the Amazon Web
// Services Resource Explorer User Guide.
//
// For more details about what happens when you turn on Resource Explorer in an
// Amazon Web Services Region, see [Turn on Resource Explorer to index your resources in an Amazon Web Services Region]in the Amazon Web Services Resource Explorer
// User Guide.
//
// If this is the first Amazon Web Services Region in which you've created an
// index for Resource Explorer, then this operation also [creates a service-linked role]in your Amazon Web
// Services account that allows Resource Explorer to enumerate your resources to
// populate the index.
//
// - Action: resource-explorer-2:CreateIndex
//
// Resource: The ARN of the index (as it will exist after the operation completes)
//
// in the Amazon Web Services Region and account in which you're trying to create
// the index. Use the wildcard character ( * ) at the end of the string to match
// the eventual UUID. For example, the following Resource element restricts the
// role or user to creating an index in only the us-east-2 Region of the
// specified account.
//
// "Resource": "arn:aws:resource-explorer-2:us-west-2:<account-id>:index/*"
//
// Alternatively, you can use "Resource": "*" to allow the role or user to create
//
// an index in any Region.
//
// - Action: iam:CreateServiceLinkedRole
//
// Resource: No specific resource (*).
//
// This permission is required only the first time you create an index to turn on
//
// Resource Explorer in the account. Resource Explorer uses this to create the [service-linked role needed to index the resources in your account].
// Resource Explorer uses the same service-linked role for all additional indexes
// you create afterwards.
//
// [Turning on cross-Region search by creating an aggregator index]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-aggregator-region.html
// [creates a service-linked role]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/security_iam_service-linked-roles.html
// [Turn on Resource Explorer to index your resources in an Amazon Web Services Region]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-service-activate.html
//
// [service-linked role needed to index the resources in your account]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/security_iam_service-linked-roles.html
