package appflow

// CancelFlowExecutions is generated as a reference stub.
// Executable command wiring lives under cmd/appflow.go.
//
// Cancels active runs for a flow.
//
// You can cancel all of the active runs for a flow, or you can cancel specific
// runs by providing their IDs.
//
// You can cancel a flow run only when the run is in progress. You can't cancel a
// run that has already completed or failed. You also can't cancel a run that's
// scheduled to occur but hasn't started yet. To prevent a scheduled run, you can
// deactivate the flow with the StopFlow action.
//
// You cannot resume a run after you cancel it.
//
// When you send your request, the status for each run becomes CancelStarted . When
// the cancellation completes, the status becomes Canceled .
//
// When you cancel a run, you still incur charges for any data that the run
// already processed before the cancellation. If the run had already written some
// data to the flow destination, then that data remains in the destination. If you
// configured the flow to use a batch API (such as the Salesforce Bulk API 2.0),
// then the run will finish reading or writing its entire batch of data after the
// cancellation. For these operations, the data processing charges for Amazon
// AppFlow apply. For the pricing information, see [Amazon AppFlow pricing].
//
// [Amazon AppFlow pricing]: http://aws.amazon.com/appflow/pricing/
