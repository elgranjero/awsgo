package textract

// GetLendingAnalysisSummary is generated as a reference stub.
// Executable command wiring lives under cmd/textract.go.
//
// Gets summarized results for the StartLendingAnalysis operation, which analyzes
// text in a lending document. The returned summary consists of information about
// documents grouped together by a common document type. Information like detected
// signatures, page numbers, and split documents is returned with respect to the
// type of grouped document.
//
// You start asynchronous text analysis by calling StartLendingAnalysis , which
// returns a job identifier ( JobId ). When the text analysis operation finishes,
// Amazon Textract publishes a completion status to the Amazon Simple Notification
// Service (Amazon SNS) topic that's registered in the initial call to
// StartLendingAnalysis .
//
// To get the results of the text analysis operation, first check that the status
// value published to the Amazon SNS topic is SUCCEEDED. If so, call
// GetLendingAnalysisSummary , and pass the job identifier ( JobId ) from the
// initial call to StartLendingAnalysis .
