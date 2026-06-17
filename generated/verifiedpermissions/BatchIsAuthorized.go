package verifiedpermissions

// BatchIsAuthorized is generated as a reference stub.
// Executable command wiring lives under cmd/verifiedpermissions.go.
//
// Makes a series of decisions about multiple authorization requests for one
// principal or resource. Each request contains the equivalent content of an
// IsAuthorized request: principal, action, resource, and context. Either the
// principal or the resource parameter must be identical across all requests. For
// example, Verified Permissions won't evaluate a pair of requests where bob views
// photo1 and alice views photo2 . Authorization of bob to view photo1 and photo2 ,
// or bob and alice to view photo1 , are valid batches.
//
// The request is evaluated against all policies in the specified policy store
// that match the entities that you declare. The result of the decisions is a
// series of Allow or Deny responses, along with the IDs of the policies that
// produced each decision.
//
// The entities of a BatchIsAuthorized API request can contain up to 100
// principals and up to 100 resources. The requests of a BatchIsAuthorized API
// request can contain up to 30 requests.
//
// The BatchIsAuthorized operation doesn't have its own IAM permission. To
// authorize this operation for Amazon Web Services principals, include the
// permission verifiedpermissions:IsAuthorized in their IAM policies.
