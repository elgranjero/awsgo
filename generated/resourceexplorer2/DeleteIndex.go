package resourceexplorer2

// DeleteIndex is generated as a reference stub.
// Executable command wiring lives under cmd/resourceexplorer2.go.
//
// Deletes the specified index and turns off Amazon Web Services Resource Explorer
// in the specified Amazon Web Services Region. When you delete an index, Resource
// Explorer stops discovering and indexing resources in that Region. Resource
// Explorer also deletes all views in that Region. These actions occur as
// asynchronous background tasks. You can check to see when the actions are
// complete by using the GetIndexoperation and checking the Status response value.
//
// If the index you delete is the aggregator index for the Amazon Web Services
// account, you must wait 24 hours before you can promote another local index to be
// the aggregator index for the account. Users can't perform account-wide searches
// using Resource Explorer until another aggregator index is configured.
