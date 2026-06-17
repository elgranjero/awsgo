package verifiedpermissions

// BatchIsAuthorizedWithToken is generated as a reference stub.
// Executable command wiring lives under cmd/verifiedpermissions.go.
//
// Makes a series of decisions about multiple authorization requests for one
// token. The principal in this request comes from an external identity source in
// the form of an identity or access token, formatted as a [JSON web token (JWT)]. The information in
// the parameters can also define additional context that Verified Permissions can
// include in the evaluations.
//
// The request is evaluated against all policies in the specified policy store
// that match the entities that you provide in the entities declaration and in the
// token. The result of the decisions is a series of Allow or Deny responses,
// along with the IDs of the policies that produced each decision.
//
// The entities of a BatchIsAuthorizedWithToken API request can contain up to 100
// resources and up to 99 user groups. The requests of a BatchIsAuthorizedWithToken
// API request can contain up to 30 requests.
//
// The BatchIsAuthorizedWithToken operation doesn't have its own IAM permission.
// To authorize this operation for Amazon Web Services principals, include the
// permission verifiedpermissions:IsAuthorizedWithToken in their IAM policies.
//
// [JSON web token (JWT)]: https://wikipedia.org/wiki/JSON_Web_Token
