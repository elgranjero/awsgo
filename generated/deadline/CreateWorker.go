package deadline

// CreateWorker is generated as a reference stub.
// Executable command wiring lives under cmd/deadline.go.
//
// Creates a worker. A worker tells your instance how much processing power
// (vCPU), and memory (GiB) you’ll need to assemble the digital assets held within
// a particular instance. You can specify certain instance types to use, or let the
// worker know which instances types to exclude.
//
// Deadline Cloud limits the number of workers to less than or equal to the
// fleet's maximum worker count. The service maintains eventual consistency for the
// worker count. If you make multiple rapid calls to CreateWorker before the field
// updates, you might exceed your fleet's maximum worker count. For example, if
// your maxWorkerCount is 10 and you currently have 9 workers, making two quick
// CreateWorker calls might successfully create 2 workers instead of 1, resulting
// in 11 total workers.
