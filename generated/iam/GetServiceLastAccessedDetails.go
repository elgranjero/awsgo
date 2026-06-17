package iam

// GetServiceLastAccessedDetails is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Retrieves a service last accessed report that was created using the
// GenerateServiceLastAccessedDetails operation. You can use the JobId parameter
// in GetServiceLastAccessedDetails to retrieve the status of your report job.
// When the report is complete, you can retrieve the generated report. The report
// includes a list of Amazon Web Services services that the resource (user, group,
// role, or managed policy) can access.
//
// Service last accessed data does not use other policy types when determining
// whether a resource could access a service. These other policy types include
// resource-based policies, access control lists, Organizations policies, IAM
// permissions boundaries, and STS assume role policies. It only applies
// permissions policy logic. For more about the evaluation of policy types, see [Evaluating policies]in
// the IAM User Guide.
//
// For each service that the resource could access using permissions policies, the
// operation returns details about the most recent access attempt. If there was no
// attempt, the service is listed without details about the most recent attempt to
// access the service. If the operation fails, the GetServiceLastAccessedDetails
// operation returns the reason that it failed.
//
// The GetServiceLastAccessedDetails operation returns a list of services. This
// list includes the number of entities that have attempted to access the service
// and the date and time of the last attempt. It also returns the ARN of the
// following entity, depending on the resource ARN that you used to generate the
// report:
//
// - User – Returns the user ARN that you used to generate the report
//
// - Group – Returns the ARN of the group member (user) that last attempted to
// access the service
//
// - Role – Returns the role ARN that you used to generate the report
//
// - Policy – Returns the ARN of the user or role that last used the policy to
// attempt to access the service
//
// By default, the list is sorted by service namespace.
//
// If you specified ACTION_LEVEL granularity when you generated the report, this
// operation returns service and action last accessed data. This includes the most
// recent access attempt for each tracked action within a service. Otherwise, this
// operation returns only service data.
//
// For more information about service and action last accessed data, see [Reducing permissions using service last accessed data] in the
// IAM User Guide.
//
// [Reducing permissions using service last accessed data]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html
// [Evaluating policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-basics
