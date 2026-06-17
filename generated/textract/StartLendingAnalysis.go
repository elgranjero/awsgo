package textract

// StartLendingAnalysis is generated as a reference stub.
// Executable command wiring lives under cmd/textract.go.
//
// Starts the classification and analysis of an input document.
// StartLendingAnalysis initiates the classification and analysis of a packet of
// lending documents. StartLendingAnalysis operates on a document file located in
// an Amazon S3 bucket.
//
// StartLendingAnalysis can analyze text in documents that are in one of the
// following formats: JPEG, PNG, TIFF, PDF. Use DocumentLocation to specify the
// bucket name and the file name of the document.
//
// StartLendingAnalysis returns a job identifier ( JobId ) that you use to get the
// results of the operation. When the text analysis is finished, Amazon Textract
// publishes a completion status to the Amazon Simple Notification Service (Amazon
// SNS) topic that you specify in NotificationChannel . To get the results of the
// text analysis operation, first check that the status value published to the
// Amazon SNS topic is SUCCEEDED. If the status is SUCCEEDED you can call either
// GetLendingAnalysis or GetLendingAnalysisSummary and provide the JobId to obtain
// the results of the analysis.
//
// If using OutputConfig to specify an Amazon S3 bucket, the output will be
// contained within the specified prefix in a directory labeled with the job-id. In
// the directory there are 3 sub-directories:
//
// - detailedResponse (contains the GetLendingAnalysis response)
//
// - summaryResponse (for the GetLendingAnalysisSummary response)
//
// - splitDocuments (documents split across logical boundaries)
