package ssm

// PutComplianceItems is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// Registers a compliance type and other compliance details on a designated
// resource. This operation lets you register custom compliance details with a
// resource. This call overwrites existing compliance information on the resource,
// so you must provide a full list of compliance items each time that you send the
// request.
//
// ComplianceType can be one of the following:
//
// - ExecutionId: The execution ID when the patch, association, or custom
// compliance item was applied.
//
// - ExecutionType: Specify patch, association, or Custom: string .
//
// - ExecutionTime. The time the patch, association, or custom compliance item
// was applied to the managed node.
//
// For State Manager associations, this represents the time when compliance status
//
// was captured by the Systems Manager service during its internal compliance
// aggregation workflow, not necessarily when the association was executed on the
// managed node. State Manager updates compliance information for all associations
// on an instance whenever any association executes, which may result in multiple
// associations showing the same execution time.
//
// - Id: The patch, association, or custom compliance ID.
//
// - Title: A title.
//
// - Status: The status of the compliance item. For example, approved for
// patches, or Failed for associations.
//
// - Severity: A patch severity. For example, Critical .
//
// - DocumentName: An SSM document name. For example, AWS-RunPatchBaseline .
//
// - DocumentVersion: An SSM document version number. For example, 4.
//
// - Classification: A patch classification. For example, security updates .
//
// - PatchBaselineId: A patch baseline ID.
//
// - PatchSeverity: A patch severity. For example, Critical .
//
// - PatchState: A patch state. For example, InstancesWithFailedPatches .
//
// - PatchGroup: The name of a patch group.
//
// - InstalledTime: The time the association, patch, or custom compliance item
// was applied to the resource. Specify the time by using the following format:
// yyyy-MM-dd'T'HH:mm:ss'Z'
