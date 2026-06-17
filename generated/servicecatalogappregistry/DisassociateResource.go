package servicecatalogappregistry

// DisassociateResource is generated as a reference stub.
// Executable command wiring lives under cmd/servicecatalogappregistry.go.
//
// Disassociates a resource from application. Both the resource and the
//
// application can be specified either by ID or name.
//
// # Minimum permissions
//
// You must have the following permissions to remove a resource that's been
// associated with an application using the APPLY_APPLICATION_TAG option for [AssociateResource].
//
// - tag:GetResources
//
// - tag:UntagResources
//
// You must also have the following permissions if you don't use the
// AWSServiceCatalogAppRegistryFullAccess policy. For more information, see [AWSServiceCatalogAppRegistryFullAccess] in
// the AppRegistry Administrator Guide.
//
// - resource-groups:DisassociateResource
//
// - cloudformation:UpdateStack
//
// - cloudformation:DescribeStacks
//
// In addition, you must have the tagging permission defined by the Amazon Web
// Services service that creates the resource. For more information, see [UntagResources]in the
// Resource Groups Tagging API Reference.
//
// [UntagResources]: https://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/API_UntTagResources.html
// [AWSServiceCatalogAppRegistryFullAccess]: https://docs.aws.amazon.com/servicecatalog/latest/arguide/full.html
// [AssociateResource]: https://docs.aws.amazon.com/servicecatalog/latest/dg/API_app-registry_AssociateResource.html
