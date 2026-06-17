package deadline

// CreateLimit is generated as a reference stub.
// Executable command wiring lives under cmd/deadline.go.
//
// Creates a limit that manages the distribution of shared resources, such as
// floating licenses. A limit can throttle work assignments, help manage workloads,
// and track current usage. Before you use a limit, you must associate the limit
// with one or more queues.
//
// You must add the amountRequirementName to a step in a job template to declare
// the limit requirement.
