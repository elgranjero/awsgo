package quicksight

// DeleteAnalysis is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Deletes an analysis from Amazon Quick Sight. You can optionally include a
// recovery window during which you can restore the analysis. If you don't specify
// a recovery window value, the operation defaults to 30 days. Amazon Quick Sight
// attaches a DeletionTime stamp to the response that specifies the end of the
// recovery window. At the end of the recovery window, Amazon Quick Sight deletes
// the analysis permanently.
//
// At any time before recovery window ends, you can use the RestoreAnalysis API
// operation to remove the DeletionTime stamp and cancel the deletion of the
// analysis. The analysis remains visible in the API until it's deleted, so you can
// describe it but you can't make a template from it.
//
// An analysis that's scheduled for deletion isn't accessible in the Amazon Quick
// Sight console. To access it in the console, restore it. Deleting an analysis
// doesn't delete the dashboards that you publish from it.
