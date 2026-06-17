package oam

// CreateLink is generated as a reference stub.
// Executable command wiring lives under cmd/oam.go.
//
// Creates a link between a source account and a sink that you have created in a
// monitoring account. After the link is created, data is sent from the source
// account to the monitoring account. When you create a link, you can optionally
// specify filters that specify which metric namespaces and which log groups are
// shared from the source account to the monitoring account.
//
// Before you create a link, you must create a sink in the monitoring account and
// create a sink policy in that account. The sink policy must permit the source
// account to link to it. You can grant permission to source accounts by granting
// permission to an entire organization or to individual accounts.
//
// For more information, see [CreateSink] and [PutSinkPolicy].
//
// Each monitoring account can be linked to as many as 100,000 source accounts.
//
// Each source account can be linked to as many as five monitoring accounts.
//
// [CreateSink]: https://docs.aws.amazon.com/OAM/latest/APIReference/API_CreateSink.html
// [PutSinkPolicy]: https://docs.aws.amazon.com/OAM/latest/APIReference/API_PutSinkPolicy.html
