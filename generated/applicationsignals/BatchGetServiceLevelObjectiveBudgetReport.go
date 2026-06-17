package applicationsignals

// BatchGetServiceLevelObjectiveBudgetReport is generated as a reference stub.
// Executable command wiring lives under cmd/applicationsignals.go.
//
// Use this operation to retrieve one or more service level objective (SLO) budget
// reports.
//
// An error budget is the amount of time or requests in an unhealthy state that
// your service can accumulate during an interval before your overall SLO budget
// health is breached and the SLO is considered to be unmet. For example, an SLO
// with a threshold of 99.95% and a monthly interval translates to an error budget
// of 21.9 minutes of downtime in a 30-day month.
//
// Budget reports include a health indicator, the attainment value, and remaining
// budget.
//
// For more information about SLO error budgets, see [SLO concepts].
//
// [SLO concepts]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-ServiceLevelObjectives.html#CloudWatch-ServiceLevelObjectives-concepts
