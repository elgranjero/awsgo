package codeguruprofiler

// GetProfile is generated as a reference stub.
// Executable command wiring lives under cmd/codeguruprofiler.go.
//
// Gets the aggregated profile of a profiling group for a specified time range.
//
// Amazon CodeGuru Profiler collects posted agent profiles for a profiling group
// into aggregated profiles.
//
// Because aggregated profiles expire over time GetProfile is not idempotent.
//
// Specify the time range for the requested aggregated profile using 1 or 2 of the
// following parameters: startTime , endTime , period . The maximum time range
// allowed is 7 days. If you specify all 3 parameters, an exception is thrown. If
// you specify only period , the latest aggregated profile is returned.
//
// Aggregated profiles are available with aggregation periods of 5 minutes, 1
// hour, and 1 day, aligned to UTC. The aggregation period of an aggregated profile
// determines how long it is retained. For more information, see [AggregatedProfileTime]
// AggregatedProfileTime . The aggregated profile's aggregation period determines
// how long
//
// it is retained by CodeGuru Profiler.
//
// - If the aggregation period is 5 minutes, the aggregated profile is retained
// for 15 days.
//
// - If the aggregation period is 1 hour, the aggregated profile is retained for
// 60 days.
//
// - If the aggregation period is 1 day, the aggregated profile is retained for
// 3 years.
//
// There are two use cases for calling GetProfile .
//
// - If you want to return an aggregated profile that already exists, use [ListProfileTimes]
// ListProfileTimes to view the time ranges of existing aggregated profiles. Use
// them in a GetProfile request to return a specific, existing aggregated
// profile.
//
// - If you want to return an aggregated profile for a time range that doesn't
// align with an existing aggregated profile, then CodeGuru Profiler makes a best
// effort to combine existing aggregated profiles from the requested time range and
// return them as one aggregated profile.
//
// If aggregated profiles do not exist for the full time range requested, then
//
// aggregated profiles for a smaller time range are returned. For example, if the
// requested time range is from 00:00 to 00:20, and the existing aggregated
// profiles are from 00:15 and 00:25, then the aggregated profiles from 00:15 to
// 00:20 are returned.
//
// [ListProfileTimes]: https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ListProfileTimes.html
// [AggregatedProfileTime]: https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_AggregatedProfileTime.html
