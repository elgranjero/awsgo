package iam

// GenerateServiceLastAccessedDetails is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Generates a report that includes details about when an IAM resource (user,
// group, role, or policy) was last used in an attempt to access Amazon Web
// Services services. Recent activity usually appears within four hours. IAM
// reports activity for at least the last 400 days, or less if your Region began
// supporting this feature within the last year. For more information, see [Regions where data is tracked]. For
// more information about services and actions for which action last accessed
// information is displayed, see [IAM action last accessed information services and actions].
//
// The service last accessed data includes all attempts to access an Amazon Web
// Services API, not just the successful ones. This includes all attempts that were
// made using the Amazon Web Services Management Console, the Amazon Web Services
// API through any of the SDKs, or any of the command line tools. An unexpected
// entry in the service last accessed data does not mean that your account has been
// compromised, because the request might have been denied. Refer to your
// CloudTrail logs as the authoritative source for information about all API calls
// and whether they were successful or denied access. For more information, see [Logging IAM events with CloudTrail]in
// the IAM User Guide.
//
// The GenerateServiceLastAccessedDetails operation returns a JobId . Use this
// parameter in the following operations to retrieve the following details from
// your report:
//
// [GetServiceLastAccessedDetails]
// - – Use this operation for users, groups, roles, or policies to list every
// Amazon Web Services service that the resource could access using permissions
// policies. For each service, the response includes information about the most
// recent access attempt.
//
// The JobId returned by GenerateServiceLastAccessedDetail must be used by the same
//
// role within a session, or by the same user when used to call
// GetServiceLastAccessedDetail .
//
// [GetServiceLastAccessedDetailsWithEntities]
// - – Use this operation for groups and policies to list information about the
// associated entities (users or roles) that attempted to access a specific Amazon
// Web Services service.
//
// To check the status of the GenerateServiceLastAccessedDetails request, use the
// JobId parameter in the same operations and test the JobStatus response
// parameter.
//
// For additional information about the permissions policies that allow an
// identity (user, group, or role) to access specific services, use the [ListPoliciesGrantingServiceAccess]operation.
//
// Service last accessed data does not use other policy types when determining
// whether a resource could access a service. These other policy types include
// resource-based policies, access control lists, Organizations policies, IAM
// permissions boundaries, and STS assume role policies. It only applies
// permissions policy logic. For more about the evaluation of policy types, see [Evaluating policies]in
// the IAM User Guide.
//
// For more information about service and action last accessed data, see [Reducing permissions using service last accessed data] in the
// IAM User Guide.
//
// [Logging IAM events with CloudTrail]: https://docs.aws.amazon.com/IAM/latest/UserGuide/cloudtrail-integration.html
// [GetServiceLastAccessedDetails]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetServiceLastAccessedDetails.html
// [ListPoliciesGrantingServiceAccess]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListPoliciesGrantingServiceAccess.html
// [Reducing permissions using service last accessed data]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html
// [Regions where data is tracked]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html#access-advisor_tracking-period
// [Evaluating policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-basics
// [GetServiceLastAccessedDetailsWithEntities]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetServiceLastAccessedDetailsWithEntities.html
// [IAM action last accessed information services and actions]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor-action-last-accessed.html
