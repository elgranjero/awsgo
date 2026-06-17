package directoryservice

// StartADAssessment is generated as a reference stub.
// Executable command wiring lives under cmd/directoryservice.go.
//
// Initiates a directory assessment to validate your self-managed AD environment
// for hybrid domain join. The assessment checks compatibility and connectivity of
// the self-managed AD environment.
//
// A directory assessment is automatically created when you create a hybrid
// directory. There are two types of assessments: CUSTOMER and SYSTEM . Your Amazon
// Web Services account has a limit of 100 CUSTOMER directory assessments.
//
// The assessment process typically takes 30 minutes or more to complete. The
// assessment process is asynchronous and you can monitor it with
// DescribeADAssessment .
//
// The InstanceIds must have a one-to-one correspondence with CustomerDnsIps ,
// meaning that if the IP address for instance i-10243410 is 10.24.34.100 and the
// IP address for instance i-10243420 is 10.24.34.200, then the input arrays must
// maintain the same order relationship, either [10.24.34.100, 10.24.34.200] paired
// with [i-10243410, i-10243420] or [10.24.34.200, 10.24.34.100] paired with
// [i-10243420, i-10243410].
//
// Note: You must provide exactly one DirectoryId or AssessmentConfiguration .
