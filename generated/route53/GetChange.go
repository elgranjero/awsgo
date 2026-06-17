package route53

// GetChange is generated as a reference stub.
// Executable command wiring lives under cmd/route53.go.
//
// Returns the current status of a change batch request. The status is one of the
// following values:
//
// - PENDING indicates that the changes in this request have not propagated to
// all Amazon Route 53 DNS servers managing the hosted zone. This is the initial
// status of all change batch requests.
//
// - INSYNC indicates that the changes have propagated to all Route 53 DNS
// servers managing the hosted zone.
