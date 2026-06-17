package oam

// CreateSink is generated as a reference stub.
// Executable command wiring lives under cmd/oam.go.
//
// Use this to create a sink in the current account, so that it can be used as a
// monitoring account in CloudWatch cross-account observability. A sink is a
// resource that represents an attachment point in a monitoring account. Source
// accounts can link to the sink to send observability data.
//
// After you create a sink, you must create a sink policy that allows source
// accounts to attach to it. For more information, see [PutSinkPolicy].
//
// Each account can contain one sink per Region. If you delete a sink, you can
// then create a new one in that Region.
//
// [PutSinkPolicy]: https://docs.aws.amazon.com/OAM/latest/APIReference/API_PutSinkPolicy.html
