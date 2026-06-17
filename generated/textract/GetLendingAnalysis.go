package textract

// GetLendingAnalysis is generated as a reference stub.
// Executable command wiring lives under cmd/textract.go.
//
// Gets the results for an Amazon Textract asynchronous operation that analyzes
// text in a lending document.
//
// You start asynchronous text analysis by calling StartLendingAnalysis , which
// returns a job identifier ( JobId ). When the text analysis operation finishes,
// Amazon Textract publishes a completion status to the Amazon Simple Notification
// Service (Amazon SNS) topic that's registered in the initial call to
// StartLendingAnalysis .
//
// To get the results of the text analysis operation, first check that the status
// value published to the Amazon SNS topic is SUCCEEDED. If so, call
// GetLendingAnalysis, and pass the job identifier ( JobId ) from the initial call
// to StartLendingAnalysis .
