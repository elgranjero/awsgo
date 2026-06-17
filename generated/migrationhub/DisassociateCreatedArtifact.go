package migrationhub

// DisassociateCreatedArtifact is generated as a reference stub.
// Executable command wiring lives under cmd/migrationhub.go.
//
// Disassociates a created artifact of an AWS resource with a migration task
// performed by a migration tool that was previously associated. This API has the
// following traits:
//
// - A migration user can call the DisassociateCreatedArtifacts operation to
// disassociate a created AWS Artifact from a migration task.
//
// - The created artifact name must be provided in ARN (Amazon Resource Name)
// format which will contain information about type and region; for example:
// arn:aws:ec2:us-east-1:488216288981:image/ami-6d0ba87b .
//
// - Examples of the AWS resource behind the created artifact are, AMI's, EC2
// instance, or RDS instance, etc.
