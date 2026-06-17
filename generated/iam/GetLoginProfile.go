package iam

// GetLoginProfile is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Retrieves the user name for the specified IAM user. A login profile is created
// when you create a password for the user to access the Amazon Web Services
// Management Console. If the user does not exist or does not have a password, the
// operation returns a 404 ( NoSuchEntity ) error.
//
// If you create an IAM user with access to the console, the CreateDate reflects
// the date you created the initial password for the user.
//
// If you create an IAM user with programmatic access, and then later add a
// password for the user to access the Amazon Web Services Management Console, the
// CreateDate reflects the initial password creation date. A user with programmatic
// access does not have a login profile unless you create a password for the user
// to access the Amazon Web Services Management Console.
