package acmpca

// PutPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/acmpca.go.
//
// Attaches a resource-based policy to a private CA.
//
// A policy can also be applied by sharing a private CA through Amazon Web
// Services Resource Access Manager (RAM). For more information, see [Attach a Policy for Cross-Account Access].
//
// The policy can be displayed with [GetPolicy] and removed with [DeletePolicy].
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
// [Using a Resource Based Policy with Amazon Web Services Private CA]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-rbp.html
// [Using a Service Linked Role with ACM]: https://docs.aws.amazon.com/acm/latest/userguide/acm-slr.html
// [Attach a Policy for Cross-Account Access]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-ram.html
// [DeletePolicy]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_DeletePolicy.html
// [GetPolicy]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_GetPolicy.html
