package ivs

// PutMetadata is generated as a reference stub.
// Executable command wiring lives under cmd/ivs.go.
//
// Inserts metadata into the active stream of the specified channel. At most 5
// requests per second per channel are allowed, each with a maximum 1 KB payload.
// (If 5 TPS is not sufficient for your needs, we recommend batching your data into
// a single PutMetadata call.) At most 155 requests per second per account are
// allowed. Also see [Embedding Metadata within a Video Stream]in the Amazon IVS User Guide.
//
// [Embedding Metadata within a Video Stream]: https://docs.aws.amazon.com/ivs/latest/userguide/metadata.html
