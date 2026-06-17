package applicationsignals

// CreateServiceLevelObjective is generated as a reference stub.
// Executable command wiring lives under cmd/applicationsignals.go.
//
// Creates a service level objective (SLO), which can help you ensure that your
// critical business operations are meeting customer expectations. Use SLOs to set
// and track specific target levels for the reliability and availability of your
// applications and services. SLOs use service level indicators (SLIs) to calculate
// whether the application is performing at the level that you want.
//
// Create an SLO to set a target for a service or operation’s availability or
// latency. CloudWatch measures this target frequently you can find whether it has
// been breached.
//
// The target performance quality that is defined for an SLO is the attainment
// goal.
//
// You can set SLO targets for your applications that are discovered by
// Application Signals, using critical metrics such as latency and availability.
// You can also set SLOs against any CloudWatch metric or math expression that
// produces a time series.
//
// You can't create an SLO for a service operation that was discovered by
// Application Signals until after that operation has reported standard metrics to
// Application Signals.
//
// When you create an SLO, you specify whether it is a period-based SLO or a
// request-based SLO. Each type of SLO has a different way of evaluating your
// application's performance against its attainment goal.
//
// - A period-based SLO uses defined periods of time within a specified total
// time interval. For each period of time, Application Signals determines whether
// the application met its goal. The attainment rate is calculated as the number
// of good periods/number of total periods .
//
// For example, for a period-based SLO, meeting an attainment goal of 99.9% means
//
// that within your interval, your application must meet its performance goal
// during at least 99.9% of the time periods.
//
// - A request-based SLO doesn't use pre-defined periods of time. Instead, the
// SLO measures number of good requests/number of total requests during the
// interval. At any time, you can find the ratio of good requests to total requests
// for the interval up to the time stamp that you specify, and measure that ratio
// against the goal set in your SLO.
//
// After you have created an SLO, you can retrieve error budget reports for it. An
// error budget is the amount of time or amount of requests that your application
// can be non-compliant with the SLO's goal, and still have your application meet
// the goal.
//
// - For a period-based SLO, the error budget starts at a number defined by the
// highest number of periods that can fail to meet the threshold, while still
// meeting the overall goal. The remaining error budget decreases with every failed
// period that is recorded. The error budget within one interval can never
// increase.
//
// For example, an SLO with a threshold that 99.95% of requests must be completed
//
// under 2000ms every month translates to an error budget of 21.9 minutes of
// downtime per month.
//
// - For a request-based SLO, the remaining error budget is dynamic and can
// increase or decrease, depending on the ratio of good requests to total requests.
//
// For more information about SLOs, see [Service level objectives (SLOs)].
//
// When you perform a CreateServiceLevelObjective operation, Application Signals
// creates the AWSServiceRoleForCloudWatchApplicationSignals service-linked role,
// if it doesn't already exist in your account. This service- linked role has the
// following permissions:
//
// - xray:GetServiceGraph
//
// - logs:StartQuery
//
// - logs:GetQueryResults
//
// - cloudwatch:GetMetricData
//
// - cloudwatch:ListMetrics
//
// - tag:GetResources
//
// - autoscaling:DescribeAutoScalingGroups
//
// [Service level objectives (SLOs)]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-ServiceLevelObjectives.html
