package directoryservice

// ResetUserPassword is generated as a reference stub.
// Executable command wiring lives under cmd/directoryservice.go.
//
// Resets the password for any user in your Managed Microsoft AD or Simple AD
// directory. Disabled users will become enabled and can be authenticated following
// the API call.
//
// You can reset the password for any user in your directory with the following
// exceptions:
//
// - For Simple AD, you cannot reset the password for any user that is a member
// of either the Domain Admins or Enterprise Admins group except for the
// administrator user.
//
// - For Managed Microsoft AD, you can only reset the password for a user that
// is in an OU based off of the NetBIOS name that you typed when you created your
// directory. For example, you cannot reset the password for a user in the Amazon
// Web Services Reserved OU. For more information about the OU structure for an
// Managed Microsoft AD directory, see [What Gets Created]in the Directory Service Administration
// Guide.
//
// [What Gets Created]: https://docs.aws.amazon.com/directoryservice/latest/admin-guide/ms_ad_getting_started_what_gets_created.html
