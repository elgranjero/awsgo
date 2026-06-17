package organizations

// CreateAccount is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Creates an Amazon Web Services account that is automatically a member of the
// organization whose credentials made the request. This is an asynchronous request
// that Amazon Web Services performs in the background. Because CreateAccount
// operates asynchronously, it can return a successful completion message even
// though account initialization might still be in progress. You might need to wait
// a few minutes before you can successfully access the account. To check the
// status of the request, do one of the following:
//
// - Use the Id value of the CreateAccountStatus response element from this
// operation to provide as a parameter to the DescribeCreateAccountStatusoperation.
//
// - Check the CloudTrail log for the CreateAccountResult event. For information
// on using CloudTrail with Organizations, see [Logging and monitoring in Organizations]in the Organizations User Guide.
//
// The user who calls the API to create an account must have the
// organizations:CreateAccount permission. If you enabled all features in the
// organization, Organizations creates the required service-linked role named
// AWSServiceRoleForOrganizations . For more information, see [Organizations and service-linked roles] in the
// Organizations User Guide.
//
// If the request includes tags, then the requester must have the
// organizations:TagResource permission.
//
// Organizations preconfigures the new member account with a role (named
// OrganizationAccountAccessRole by default) that grants users in the management
// account administrator permissions in the new member account. Principals in the
// management account can assume the role. Organizations clones the company name
// and address information for the new account from the organization's management
// account.
//
// You can only call this operation from the management account.
//
// For more information about creating accounts, see [Creating a member account in your organization] in the Organizations User
// Guide.
//
// - When you create an account in an organization using the Organizations
// console, API, or CLI commands, the information required for the account to
// operate as a standalone account, such as a payment method is not automatically
// collected. If you must remove an account from your organization later, you can
// do so only after you provide the missing information. For more information, see [Considerations before removing an account from an organization]
// in the Organizations User Guide.
//
// - If you get an exception that indicates that you exceeded your account
// limits for the organization, contact [Amazon Web Services Support].
//
// - If you get an exception that indicates that the operation failed because
// your organization is still initializing, wait one hour and then try again. If
// the error persists, contact [Amazon Web Services Support].
//
// - It isn't recommended to use CreateAccount to create multiple temporary
// accounts, and using the CreateAccount API to close accounts is subject to a
// 30-day usage quota. For information on the requirements and process for closing
// an account, see [Closing a member account in your organization]in the Organizations User Guide.
//
// When you create a member account with this operation, you can choose whether to
// create the account with the IAM User and Role Access to Billing Information
// switch enabled. If you enable it, IAM users and roles that have appropriate
// permissions can view billing information for the account. If you disable it,
// only the account root user can access billing information. For information about
// how to disable this switch for an account, see [Granting access to your billing information and tools].
//
// [Granting access to your billing information and tools]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/control-access-billing.html#grantaccess
// [Amazon Web Services Support]: https://console.aws.amazon.com/support/home#/
// [Logging and monitoring in Organizations]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_security_incident-response.html#orgs_cloudtrail-integration
// [Organizations and service-linked roles]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html#orgs_integrate_services-using_slrs
// [Creating a member account in your organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_create.html
// [Considerations before removing an account from an organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_account-before-remove.html
// [Closing a member account in your organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_close.html
