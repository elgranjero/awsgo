package rbin

// CreateRule is generated as a reference stub.
// Executable command wiring lives under cmd/rbin.go.
//
// Creates a Recycle Bin retention rule. You can create two types of retention
// rules:
//
// - Tag-level retention rules - These retention rules use resource tags to
// identify the resources to protect. For each retention rule, you specify one or
// more tag key and value pairs. Resources (of the specified type) that have at
// least one of these tag key and value pairs are automatically retained in the
// Recycle Bin upon deletion. Use this type of retention rule to protect specific
// resources in your account based on their tags.
//
// - Region-level retention rules - These retention rules, by default, apply to
// all of the resources (of the specified type) in the Region, even if the
// resources are not tagged. However, you can specify exclusion tags to exclude
// resources that have specific tags. Use this type of retention rule to protect
// all resources of a specific type in a Region.
//
// For more information, see [Create Recycle Bin retention rules] in the Amazon EBS User Guide.
//
// [Create Recycle Bin retention rules]: https://docs.aws.amazon.com/ebs/latest/userguide/recycle-bin.html
