package verifiedpermissions

// IsAuthorizedWithToken is generated as a reference stub.
// Executable command wiring lives under cmd/verifiedpermissions.go.
//
// Makes an authorization decision about a service request described in the
// parameters. The principal in this request comes from an external identity source
// in the form of an identity token formatted as a [JSON web token (JWT)]. The information in the
// parameters can also define additional context that Verified Permissions can
// include in the evaluation. The request is evaluated against all matching
// policies in the specified policy store. The result of the decision is either
// Allow or Deny , along with a list of the policies that resulted in the decision.
//
// Verified Permissions validates each token that is specified in a request by
// checking its expiration date and its signature.
//
// Tokens from an identity source user continue to be usable until they expire.
// Token revocation and resource deletion have no effect on the validity of a token
// in your policy store
//
// [JSON web token (JWT)]: https://wikipedia.org/wiki/JSON_Web_Token
