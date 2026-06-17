package configservice

// GetDiscoveredResourceCounts is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Returns the resource types, the number of each resource type, and the total
// number of resources that Config is recording in this region for your Amazon Web
// Services account.
//
// Example
//
// - Config is recording three resource types in the US East (Ohio) Region for
// your account: 25 EC2 instances, 20 IAM users, and 15 S3 buckets.
//
// - You make a call to the GetDiscoveredResourceCounts action and specify that
// you want all resource types.
//
// - Config returns the following:
//
// - The resource types (EC2 instances, IAM users, and S3 buckets).
//
// - The number of each resource type (25, 20, and 15).
//
// - The total number of all resources (60).
//
// The response is paginated. By default, Config lists 100 ResourceCount objects on each page.
// You can customize this number with the limit parameter. The response includes a
// nextToken string. To get the next page of results, run the request again and
// specify the string for the nextToken parameter.
//
// If you make a call to the GetDiscoveredResourceCounts action, you might not immediately receive resource
// counts in the following situations:
//
// - You are a new Config customer.
//
// - You just enabled resource recording.
//
// It might take a few minutes for Config to record and count your resources. Wait
// a few minutes and then retry the GetDiscoveredResourceCountsaction.
