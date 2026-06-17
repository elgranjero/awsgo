package organizations

// ListHandshakesForOrganization is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Lists the recent handshakes that you have sent.
//
// You can view CANCELED , ACCEPTED , DECLINED , or EXPIRED handshakes in API
// responses for 30 days before they are deleted.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
//
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
