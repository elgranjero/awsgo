package connect

// DescribeContactFlow is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Describes the specified flow.
//
// You can also create and update flows using the [Amazon Connect Flow language].
//
// Use the $SAVED alias in the request to describe the SAVED content of a Flow.
// For example, arn:aws:.../contact-flow/{id}:$SAVED . After a flow is published,
// $SAVED needs to be supplied to view saved content that has not been published.
//
// Use arn:aws:.../contact-flow/{id}:{version} to retrieve the content of a
// specific flow version.
//
// In the response, Status indicates the flow status as either SAVED or PUBLISHED .
// The PUBLISHED status will initiate validation on the content. SAVED does not
// initiate validation of the content. SAVED | PUBLISHED
//
// [Amazon Connect Flow language]: https://docs.aws.amazon.com/connect/latest/APIReference/flow-language.html
