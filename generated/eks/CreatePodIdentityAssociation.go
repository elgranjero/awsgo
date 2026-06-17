package eks

// CreatePodIdentityAssociation is generated as a reference stub.
// Executable command wiring lives under cmd/eks.go.
//
// Creates an EKS Pod Identity association between a service account in an Amazon
// EKS cluster and an IAM role with EKS Pod Identity. Use EKS Pod Identity to give
// temporary IAM credentials to Pods and the credentials are rotated automatically.
//
// Amazon EKS Pod Identity associations provide the ability to manage credentials
// for your applications, similar to the way that Amazon EC2 instance profiles
// provide credentials to Amazon EC2 instances.
//
// If a Pod uses a service account that has an association, Amazon EKS sets
// environment variables in the containers of the Pod. The environment variables
// configure the Amazon Web Services SDKs, including the Command Line Interface, to
// use the EKS Pod Identity credentials.
//
// EKS Pod Identity is a simpler method than IAM roles for service accounts, as
// this method doesn't use OIDC identity providers. Additionally, you can configure
// a role for EKS Pod Identity once, and reuse it across clusters.
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
