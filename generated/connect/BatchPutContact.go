package connect

// BatchPutContact is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Only the Amazon Connect outbound campaigns service principal is allowed to
// assume a role in your account and call this API.
//
// Allows you to create a batch of contacts in Amazon Connect. The outbound
// campaigns capability ingests dial requests via the [PutDialRequestBatch]API. It then uses
// BatchPutContact to create contacts corresponding to those dial requests. If
// agents are available, the dial requests are dialed out, which results in a voice
// call. The resulting voice call uses the same contactId that was created by
// BatchPutContact.
//
// [PutDialRequestBatch]: https://docs.aws.amazon.com/connect-outbound/latest/APIReference/API_PutDialRequestBatch.html
