package transcribe

// CreateCallAnalyticsCategory is generated as a reference stub.
// Executable command wiring lives under cmd/transcribe.go.
//
// Creates a new Call Analytics category.
//
// All categories are automatically applied to your Call Analytics transcriptions.
// Note that in order to apply categories to your transcriptions, you must create
// them before submitting your transcription request, as categories cannot be
// applied retroactively.
//
// When creating a new category, you can use the InputType parameter to label the
// category as a POST_CALL or a REAL_TIME category. POST_CALL categories can only
// be applied to post-call transcriptions and REAL_TIME categories can only be
// applied to real-time transcriptions. If you do not include InputType , your
// category is created as a POST_CALL category by default.
//
// Call Analytics categories are composed of rules. For each category, you must
// create between 1 and 20 rules. Rules can include these parameters: , , , and .
//
// To update an existing category, see .
//
// To learn more about Call Analytics categories, see [Creating categories for post-call transcriptions] and [Creating categories for real-time transcriptions].
//
// [Creating categories for post-call transcriptions]: https://docs.aws.amazon.com/transcribe/latest/dg/tca-categories-batch.html
// [Creating categories for real-time transcriptions]: https://docs.aws.amazon.com/transcribe/latest/dg/tca-categories-stream.html
