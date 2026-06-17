package servicecatalogappregistry

// AssociateResource is generated as a reference stub.
// Executable command wiring lives under cmd/servicecatalogappregistry.go.
//
// Associates a resource with an application. The resource can be specified by
//
// its ARN or name. The application can be specified by ARN, ID, or name.
//
// # Minimum permissions
//
// You must have the following permissions to associate a resource using the
// OPTIONS parameter set to APPLY_APPLICATION_TAG .
//
// - tag:GetResources
//
// - tag:TagResources
//
// You must also have these additional permissions if you don't use the
// AWSServiceCatalogAppRegistryFullAccess policy. For more information, see [AWSServiceCatalogAppRegistryFullAccess] in
// the AppRegistry Administrator Guide.
//
// - resource-groups:AssociateResource
//
// - cloudformation:UpdateStack
//
// - cloudformation:DescribeStacks
//
// In addition, you must have the tagging permission defined by the Amazon Web
// Services service that creates the resource. For more information, see [TagResources]in the
// Resource Groups Tagging API Reference.
//
// [TagResources]: https://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/API_TagResources.html
// [AWSServiceCatalogAppRegistryFullAccess]: https://docs.aws.amazon.com/servicecatalog/latest/arguide/full.html
