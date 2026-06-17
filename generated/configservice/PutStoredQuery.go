package configservice

// PutStoredQuery is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Saves a new query or updates an existing saved query. The QueryName must be
// unique for a single Amazon Web Services account and a single Amazon Web Services
// Region. You can create upto 300 queries in a single Amazon Web Services account
// and a single Amazon Web Services Region.
//
// # Tags are added at creation and cannot be updated
//
// PutStoredQuery is an idempotent API. Subsequent requests won’t create a
// duplicate resource if one was already created. If a following request has
// different tags values, Config will ignore these differences and treat it as an
// idempotent request of the previous. In this case, tags will not be updated,
// even if they are different.
