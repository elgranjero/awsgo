package aiops

// CreateInvestigationGroup is generated as a reference stub.
// Executable command wiring lives under cmd/aiops.go.
//
// Creates an investigation group in your account. Creating an investigation group
// is a one-time setup task for each Region in your account. It is a necessary task
// to be able to perform investigations.
//
// Settings in the investigation group help you centrally manage the common
// properties of your investigations, such as the following:
//
// - Who can access the investigations
//
// - Whether investigation data is encrypted with a customer managed Key
// Management Service key.
//
// - How long investigations and their data are retained by default.
//
// Currently, you can have one investigation group in each Region in your account.
// Each investigation in a Region is a part of the investigation group in that
// Region
//
// To create an investigation group and set up CloudWatch investigations, you must
// be signed in to an IAM principal that has either the AIOpsConsoleAdminPolicy or
// the AdministratorAccess IAM policy attached, or to an account that has similar
// permissions.
//
// You can configure CloudWatch alarms to start investigations and add events to
// investigations. If you create your investigation group with
// CreateInvestigationGroup and you want to enable alarms to do this, you must use
// PutInvestigationGroupPolicy to create a resource policy that grants this
// permission to CloudWatch alarms.
//
// For more information about configuring CloudWatch alarms, see [Using Amazon CloudWatch alarms]
//
// [Using Amazon CloudWatch alarms]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html
