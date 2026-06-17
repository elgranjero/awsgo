package proton

// CreateRepository is generated as a reference stub.
// Executable command wiring lives under cmd/proton.go.
//
// Create and register a link to a repository. Proton uses the link to repeatedly
// access the repository, to either push to it (self-managed provisioning) or pull
// from it (template sync). You can share a linked repository across multiple
// resources (like environments using self-managed provisioning, or synced
// templates). When you create a repository link, Proton creates a [service-linked role]for you.
//
// For more information, see [Self-managed provisioning], [Template bundles], and [Template sync configurations] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Self-managed provisioning]: https://docs.aws.amazon.com/proton/latest/userguide/ag-works-prov-methods.html#ag-works-prov-methods-self
// [Template sync configurations]: https://docs.aws.amazon.com/proton/latest/userguide/ag-template-sync-configs.html
// [Template bundles]: https://docs.aws.amazon.com/proton/latest/userguide/ag-template-authoring.html#ag-template-bundles
// [service-linked role]: https://docs.aws.amazon.com/proton/latest/userguide/using-service-linked-roles.html
