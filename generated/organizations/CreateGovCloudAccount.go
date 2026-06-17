package organizations

// CreateGovCloudAccount is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// This action is available if all of the following are true:
//
// - You're authorized to create accounts in the Amazon Web Services GovCloud
// (US) Region. For more information on the Amazon Web Services GovCloud (US)
// Region, see the [Amazon Web Services GovCloud User Guide.]
//
// - You already have an account in the Amazon Web Services GovCloud (US) Region
// that is paired with a management account of an organization in the commercial
// Region.
//
// - You call this action from the management account of your organization in
// the commercial Region.
//
// - You have the organizations:CreateGovCloudAccount permission.
//
// Organizations automatically creates the required service-linked role named
// AWSServiceRoleForOrganizations . For more information, see [Organizations and service-linked roles] in the
// Organizations User Guide.
//
// Amazon Web Services automatically enables CloudTrail for Amazon Web Services
// GovCloud (US) accounts, but you should also do the following:
//
// - Verify that CloudTrail is enabled to store logs.
//
// - Create an Amazon S3 bucket for CloudTrail log storage.
//
// For more information, see [Verifying CloudTrail Is Enabled]in the Amazon Web Services GovCloud User Guide.
//
// If the request includes tags, then the requester must have the
// organizations:TagResource permission. The tags are attached to the commercial
// account associated with the GovCloud account, rather than the GovCloud account
// itself. To add tags to the GovCloud account, call the TagResourceoperation in the GovCloud
// Region after the new GovCloud account exists.
//
// You call this action from the management account of your organization in the
// commercial Region to create a standalone Amazon Web Services account in the
// Amazon Web Services GovCloud (US) Region. After the account is created, the
// management account of an organization in the Amazon Web Services GovCloud (US)
// Region can invite it to that organization. For more information on inviting
// standalone accounts in the Amazon Web Services GovCloud (US) to join an
// organization, see [Organizations]in the Amazon Web Services GovCloud User Guide.
//
// Calling CreateGovCloudAccount is an asynchronous request that Amazon Web
// Services performs in the background. Because CreateGovCloudAccount operates
// asynchronously, it can return a successful completion message even though
// account initialization might still be in progress. You might need to wait a few
// minutes before you can successfully access the account. To check the status of
// the request, do one of the following:
//
// - Use the OperationId response element from this operation to provide as a
// parameter to the DescribeCreateAccountStatusoperation.
//
// - Check the CloudTrail log for the CreateAccountResult event. For information
// on using CloudTrail with Organizations, see [Logging and monitoring in Organizations]in the Organizations User Guide.
//
// When you call the CreateGovCloudAccount action, you create two accounts: a
// standalone account in the Amazon Web Services GovCloud (US) Region and an
// associated account in the commercial Region for billing and support purposes.
// The account in the commercial Region is automatically a member of the
// organization whose credentials made the request. Both accounts are associated
// with the same email address.
//
// A role is created in the new account in the commercial Region that allows the
// management account in the organization in the commercial Region to assume it. An
// Amazon Web Services GovCloud (US) account is then created and associated with
// the commercial account that you just created. A role is also created in the new
// Amazon Web Services GovCloud (US) account that can be assumed by the Amazon Web
// Services GovCloud (US) account that is associated with the management account of
// the commercial organization. For more information and to view a diagram that
// explains how account access works, see [Organizations]in the Amazon Web Services GovCloud User
// Guide.
//
// For more information about creating accounts, see [Creating a member account in your organization] in the Organizations User
// Guide.
//
// - When you create an account in an organization using the Organizations
// console, API, or CLI commands, the information required for the account to
// operate as a standalone account is not automatically collected. This includes a
// payment method and signing the end user license agreement (EULA). If you must
// remove an account from your organization later, you can do so only after you
// provide the missing information. For more information, see [Considerations before removing an account from an organization]in the
// Organizations User Guide.
//
// - If you get an exception that indicates that you exceeded your account
// limits for the organization, contact [Amazon Web Services Support].
//
// - If you get an exception that indicates that the operation failed because
// your organization is still initializing, wait one hour and then try again. If
// the error persists, contact [Amazon Web Services Support].
//
// - Using CreateGovCloudAccount to create multiple temporary accounts isn't
// recommended. You can only close an account from the Amazon Web Services Billing
// and Cost Management console, and you must be signed in as the root user. For
// information on the requirements and process for closing an account, see [Closing a member account in your organization]in
// the Organizations User Guide.
//
// When you create a member account with this operation, you can choose whether to
// create the account with the IAM User and Role Access to Billing Information
// switch enabled. If you enable it, IAM users and roles that have appropriate
// permissions can view billing information for the account. If you disable it,
// only the account root user can access billing information. For information about
// how to disable this switch for an account, see [Granting access to your billing information and tools].
//
// [Granting access to your billing information and tools]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/grantaccess.html
// [Verifying CloudTrail Is Enabled]: https://docs.aws.amazon.com/govcloud-us/latest/UserGuide/verifying-cloudtrail.html
// [Amazon Web Services Support]: https://console.aws.amazon.com/support/home#/
// [Logging and monitoring in Organizations]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_security_incident-response.html
// [Organizations]: https://docs.aws.amazon.com/govcloud-us/latest/UserGuide/govcloud-organizations.html
// [Organizations and service-linked roles]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html#orgs_integrate_services-using_slrs
// [Creating a member account in your organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_create.html
// [Considerations before removing an account from an organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_account-before-remove.html
// [Amazon Web Services GovCloud User Guide.]: https://docs.aws.amazon.com/govcloud-us/latest/UserGuide/welcome.html
// [Closing a member account in your organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_close.html
