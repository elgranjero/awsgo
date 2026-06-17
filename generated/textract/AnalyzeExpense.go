package textract

// AnalyzeExpense is generated as a reference stub.
// Executable command wiring lives under cmd/textract.go.
//
// AnalyzeExpense synchronously analyzes an input document for financially related
// relationships between text.
//
// Information is returned as ExpenseDocuments and seperated as follows:
//
// - LineItemGroups - A data set containing LineItems which store information
// about the lines of text, such as an item purchased and its price on a receipt.
//
// - SummaryFields - Contains all other information a receipt, such as header
// information or the vendors name.
