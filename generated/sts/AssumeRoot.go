package sts

// AssumeRoot is generated as a reference stub.
// Executable command wiring lives under cmd/sts.go.
//
// Returns a set of short term credentials you can use to perform privileged tasks
// on a member account in your organization. You must use credentials from an
// Organizations management account or a delegated administrator account for IAM to
// call AssumeRoot . You cannot use root user credentials to make this call.
//
// Before you can launch a privileged session, you must have centralized root
// access in your organization. For steps to enable this feature, see [Centralize root access for member accounts]in the IAM
// User Guide.
//
// The STS global endpoint is not supported for AssumeRoot. You must send this
// request to a Regional STS endpoint. For more information, see [Endpoints].
//
// You can track AssumeRoot in CloudTrail logs to determine what actions were
// performed in a session. For more information, see [Track privileged tasks in CloudTrail]in the IAM User Guide.
//
// When granting access to privileged tasks you should only grant the necessary
// permissions required to perform that task. For more information, see [Security best practices in IAM]. In
// addition, you can use [service control policies](SCPs) to manage and limit permissions in your
// organization. See [General examples]in the Organizations User Guide for more information on SCPs.
//
// [Endpoints]: https://docs.aws.amazon.com/STS/latest/APIReference/welcome.html#sts-endpoints
// [Security best practices in IAM]: https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html
// [Track privileged tasks in CloudTrail]: https://docs.aws.amazon.com/IAM/latest/UserGuide/cloudtrail-track-privileged-tasks.html
// [General examples]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps_examples_general.html
// [service control policies]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps.html
// [Centralize root access for member accounts]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_root-enable-root-access.html
