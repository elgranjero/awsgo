package fms

// PutNotificationChannel is generated as a reference stub.
// Executable command wiring lives under cmd/fms.go.
//
// Designates the IAM role and Amazon Simple Notification Service (SNS) topic that
// Firewall Manager uses to record SNS logs.
//
// To perform this action outside of the console, you must first configure the SNS
// topic's access policy to allow the SnsRoleName to publish SNS logs. If the
// SnsRoleName provided is a role other than the AWSServiceRoleForFMS
// service-linked role, this role must have a trust relationship configured to
// allow the Firewall Manager service principal fms.amazonaws.com to assume this
// role. For information about configuring an SNS access policy, see [Service roles for Firewall Manager]in the
// Firewall Manager Developer Guide.
//
// [Service roles for Firewall Manager]: https://docs.aws.amazon.com/waf/latest/developerguide/fms-security_iam_service-with-iam.html#fms-security_iam_service-with-iam-roles-service
