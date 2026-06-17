package organizations

// CreateOrganizationalUnit is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Creates an organizational unit (OU) within a root or parent OU. An OU is a
// container for accounts that enables you to organize your accounts to apply
// policies according to your business requirements. The number of levels deep that
// you can nest OUs is dependent upon the policy types enabled for that root. For
// service control policies, the limit is five.
//
// For more information about OUs, see [Managing organizational units (OUs)] in the Organizations User Guide.
//
// If the request includes tags, then the requester must have the
// organizations:TagResource permission.
//
// You can only call this operation from the management account.
//
// [Managing organizational units (OUs)]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_ous.html
