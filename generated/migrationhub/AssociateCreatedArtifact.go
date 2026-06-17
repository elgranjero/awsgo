package migrationhub

// AssociateCreatedArtifact is generated as a reference stub.
// Executable command wiring lives under cmd/migrationhub.go.
//
// Associates a created artifact of an AWS cloud resource, the target receiving
// the migration, with the migration task performed by a migration tool. This API
// has the following traits:
//
// - Migration tools can call the AssociateCreatedArtifact operation to indicate
// which AWS artifact is associated with a migration task.
//
// - The created artifact name must be provided in ARN (Amazon Resource Name)
// format which will contain information about type and region; for example:
// arn:aws:ec2:us-east-1:488216288981:image/ami-6d0ba87b .
//
// - Examples of the AWS resource behind the created artifact are, AMI's, EC2
// instance, or DMS endpoint, etc.
