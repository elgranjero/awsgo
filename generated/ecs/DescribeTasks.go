package ecs

// DescribeTasks is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Describes a specified task or tasks.
//
// Currently, stopped tasks appear in the returned results for at least one hour.
//
// If you have tasks with tags, and then delete the cluster, the tagged tasks are
// returned in the response. If you create a new cluster with the same name as the
// deleted cluster, the tagged tasks are not included in the response.
