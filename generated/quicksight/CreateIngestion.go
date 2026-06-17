package quicksight

// CreateIngestion is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Creates and starts a new SPICE ingestion for a dataset. You can manually
// refresh datasets in an Enterprise edition account 32 times in a 24-hour period.
// You can manually refresh datasets in a Standard edition account 8 times in a
// 24-hour period. Each 24-hour period is measured starting 24 hours before the
// current date and time.
//
// Any ingestions operating on tagged datasets inherit the same tags automatically
// for use in access control. For an example, see [How do I create an IAM policy to control access to Amazon EC2 resources using tags?]in the Amazon Web Services
// Knowledge Center. Tags are visible on the tagged dataset, but not on the
// ingestion resource.
//
// [How do I create an IAM policy to control access to Amazon EC2 resources using tags?]: http://aws.amazon.com/premiumsupport/knowledge-center/iam-ec2-resource-tags/
