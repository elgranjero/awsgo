package shield

// DescribeAttackStatistics is generated as a reference stub.
// Executable command wiring lives under cmd/shield.go.
//
// Provides information about the number and type of attacks Shield has detected
// in the last year for all resources that belong to your account, regardless of
// whether you've defined Shield protections for them. This operation is available
// to Shield customers as well as to Shield Advanced customers.
//
// The operation returns data for the time range of midnight UTC, one year ago, to
// midnight UTC, today. For example, if the current time is 2020-10-26 15:39:32 PDT
// , equal to 2020-10-26 22:39:32 UTC , then the time range for the attack data
// returned is from 2019-10-26 00:00:00 UTC to 2020-10-26 00:00:00 UTC .
//
// The time range indicates the period covered by the attack statistics data items.
