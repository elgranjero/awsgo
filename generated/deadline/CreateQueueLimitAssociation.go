package deadline

// CreateQueueLimitAssociation is generated as a reference stub.
// Executable command wiring lives under cmd/deadline.go.
//
// Associates a limit with a particular queue. After the limit is associated, all
// workers for jobs that specify the limit associated with the queue are subject to
// the limit. You can't associate two limits with the same amountRequirementName
// to the same queue.
