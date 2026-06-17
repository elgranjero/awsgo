package proton

// GetResourcesSummary is generated as a reference stub.
// Executable command wiring lives under cmd/proton.go.
//
// Get counts of Proton resources.
//
// For infrastructure-provisioning resources (environments, services, service
// instances, pipelines), the action returns staleness counts. A resource is stale
// when it's behind the recommended version of the Proton template that it uses and
// it needs an update to become current.
//
// The action returns staleness counts (counts of resources that are up-to-date,
// behind a template major version, or behind a template minor version), the total
// number of resources, and the number of resources that are in a failed state,
// grouped by resource type. Components, environments, and service templates return
// less information - see the components , environments , and serviceTemplates
// field descriptions.
//
// For context, the action also returns the total number of each type of Proton
// template in the Amazon Web Services account.
//
// For more information, see [Proton dashboard] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Proton dashboard]: https://docs.aws.amazon.com/proton/latest/userguide/monitoring-dashboard.html
