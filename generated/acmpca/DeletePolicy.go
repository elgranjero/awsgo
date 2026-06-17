package acmpca

// DeletePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/acmpca.go.
//
// Deletes the resource-based policy attached to a private CA. Deletion will
// remove any access that the policy has granted. If there is no policy attached to
// the private CA, this action will return successful.
//
// If you delete a policy that was applied through Amazon Web Services Resource
// Access Manager (RAM), the CA will be removed from all shares in which it was
// included.
//
// The Certificate Manager Service Linked Role that the policy supports is not
// affected when you delete the policy.
//
// The current policy can be shown with [GetPolicy] and updated with [PutPolicy].
//
// About Policies
//
// - A policy grants access on a private CA to an Amazon Web Services customer
// account, to Amazon Web Services Organizations, or to an Amazon Web Services
// Organizations unit. Policies are under the control of a CA administrator. For
// more information, see [Using a Resource Based Policy with Amazon Web Services Private CA].
//
// - A policy permits a user of Certificate Manager (ACM) to issue ACM
// certificates signed by a CA in another account.
//
// - For ACM to manage automatic renewal of these certificates, the ACM user
// must configure a Service Linked Role (SLR). The SLR allows the ACM service to
// assume the identity of the user, subject to confirmation against the Amazon Web
// Services Private CA policy. For more information, see [Using a Service Linked Role with ACM].
//
// - Updates made in Amazon Web Services Resource Manager (RAM) are reflected in
// policies. For more information, see [Attach a Policy for Cross-Account Access].
//
// [PutPolicy]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_PutPolicy.html
// [Using a Resource Based Policy with Amazon Web Services Private CA]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-rbp.html
// [Using a Service Linked Role with ACM]: https://docs.aws.amazon.com/acm/latest/userguide/acm-slr.html
// [Attach a Policy for Cross-Account Access]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-ram.html
// [GetPolicy]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_GetPolicy.html
