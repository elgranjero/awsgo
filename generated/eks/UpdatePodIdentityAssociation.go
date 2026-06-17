package eks

// UpdatePodIdentityAssociation is generated as a reference stub.
// Executable command wiring lives under cmd/eks.go.
//
// Updates a EKS Pod Identity association. In an update, you can change the IAM
// role, the target IAM role, or disableSessionTags . You must change at least one
// of these in an update. An association can't be moved between clusters,
// namespaces, or service accounts. If you need to edit the namespace or service
// account, you need to delete the association and then create a new association
// with your desired settings.
//
// Similar to Amazon Web Services IAM behavior, EKS Pod Identity associations are
// eventually consistent, and may take several seconds to be effective after the
// initial API call returns successfully. You must design your applications to
// account for these potential delays. We recommend that you don’t include
// association create/updates in the critical, high-availability code paths of your
// application. Instead, make changes in a separate initialization or setup routine
// that you run less frequently.
//
// You can set a target IAM role in the same or a different account for advanced
// scenarios. With a target role, EKS Pod Identity automatically performs two role
// assumptions in sequence: first assuming the role in the association that is in
// this account, then using those credentials to assume the target IAM role. This
// process provides your Pod with temporary credentials that have the permissions
// defined in the target role, allowing secure access to resources in another
// Amazon Web Services account.
