package elasticloadbalancingv2

// ModifyListener is generated as a reference stub.
// Executable command wiring lives under cmd/elasticloadbalancingv2.go.
//
// Replaces the specified properties of the specified listener. Any properties
// that you do not specify remain unchanged.
//
// Changing the protocol from HTTPS to HTTP, or from TLS to TCP, removes the
// security policy and default certificate properties. If you change the protocol
// from HTTP to HTTPS, or from TCP to TLS, you must add the security policy and
// default certificate properties.
//
// To add an item to a list, remove an item from a list, or update an item in a
// list, you must provide the entire list. For example, to add an action, specify a
// list with the current actions plus the new action.
