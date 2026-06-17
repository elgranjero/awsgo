package eks

// ListInsights is generated as a reference stub.
// Executable command wiring lives under cmd/eks.go.
//
// Returns a list of all insights checked for against the specified cluster. You
// can filter which insights are returned by category, associated Kubernetes
// version, and status. The default filter lists all categories and every status.
//
// The following lists the available categories:
//
// - UPGRADE_READINESS : Amazon EKS identifies issues that could impact your
// ability to upgrade to new versions of Kubernetes. These are called upgrade
// insights.
//
// - MISCONFIGURATION : Amazon EKS identifies misconfiguration in your EKS Hybrid
// Nodes setup that could impair functionality of your cluster or workloads. These
// are called configuration insights.
