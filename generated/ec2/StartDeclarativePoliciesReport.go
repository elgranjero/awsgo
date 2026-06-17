package ec2

// StartDeclarativePoliciesReport is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Generates an account status report. The report is generated asynchronously, and
// can take several hours to complete.
//
// The report provides the current status of all attributes supported by
// declarative policies for the accounts within the specified scope. The scope is
// determined by the specified TargetId , which can represent an individual
// account, or all the accounts that fall under the specified organizational unit
// (OU) or root (the entire Amazon Web Services Organization).
//
// The report is saved to your specified S3 bucket, using the following path
// structure (with the capitalized placeholders representing your specific values):
//
// s3://AMZN-S3-DEMO-BUCKET/YOUR-OPTIONAL-S3-PREFIX/ec2_TARGETID_REPORTID_YYYYMMDDTHHMMZ.csv
//
// Prerequisites for generating a report
//
// - The StartDeclarativePoliciesReport API can only be called by the management
// account or delegated administrators for the organization.
//
// - An S3 bucket must be available before generating the report (you can create
// a new one or use an existing one), it must be in the same Region where the
// report generation request is made, and it must have an appropriate bucket
// policy. For a sample S3 policy, see Sample Amazon S3 policy under [Examples].
//
// - Trusted access must be enabled for the service for which the declarative
// policy will enforce a baseline configuration. If you use the Amazon Web Services
// Organizations console, this is done automatically when you enable declarative
// policies. The API uses the following service principal to identify the EC2
// service: ec2.amazonaws.com . For more information on how to enable trusted
// access with the Amazon Web Services CLI and Amazon Web Services SDKs, see [Using Organizations with other Amazon Web Services services]in
// the Amazon Web Services Organizations User Guide.
//
// - Only one report per organization can be generated at a time. Attempting to
// generate a report while another is in progress will result in an error.
//
// For more information, including the required IAM permissions to run this API,
// see [Generating the account status report for declarative policies]in the Amazon Web Services Organizations User Guide.
//
// [Generating the account status report for declarative policies]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_declarative_status-report.html
// [Using Organizations with other Amazon Web Services services]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html
// [Examples]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_StartDeclarativePoliciesReport.html#API_StartDeclarativePoliciesReport_Examples
