package lexmodelsv2

// ListIntentPaths is generated as a reference stub.
// Executable command wiring lives under cmd/lexmodelsv2.go.
//
// Retrieves summary statistics for a path of intents that users take over
// sessions with your bot. The following fields are required:
//
// - startDateTime and endDateTime – Define a time range for which you want to
// retrieve results.
//
// - intentPath – Define an order of intents for which you want to retrieve
// metrics. Separate intents in the path with a forward slash. For example,
// populate the intentPath field with /BookCar/BookHotel to see details about how
// many times users invoked the BookCar and BookHotel intents in that order.
//
// Use the optional filters field to filter the results.
