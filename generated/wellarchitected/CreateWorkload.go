package wellarchitected

// CreateWorkload is generated as a reference stub.
// Executable command wiring lives under cmd/wellarchitected.go.
//
// Create a new workload.
//
// The owner of a workload can share the workload with other Amazon Web Services
// accounts, users, an organization, and organizational units (OUs) in the same
// Amazon Web Services Region. Only the owner of a workload can delete it.
//
// For more information, see [Defining a Workload] in the Well-Architected Tool User Guide.
//
// Either AwsRegions , NonAwsRegions , or both must be specified when creating a
// workload.
//
// You also must specify ReviewOwner , even though the parameter is listed as not
// being required in the following section.
//
// When creating a workload using a review template, you must have the following
// IAM permissions:
//
// - wellarchitected:GetReviewTemplate
//
// - wellarchitected:GetReviewTemplateAnswer
//
// - wellarchitected:ListReviewTemplateAnswers
//
// - wellarchitected:GetReviewTemplateLensReview
//
// [Defining a Workload]: https://docs.aws.amazon.com/wellarchitected/latest/userguide/define-workload.html
