package eks

// CreateCapability is generated as a reference stub.
// Executable command wiring lives under cmd/eks.go.
//
// Creates a managed capability resource for an Amazon EKS cluster.
//
// Capabilities provide fully managed capabilities to build and scale with
// Kubernetes. When you create a capability, Amazon EKSprovisions and manages the
// infrastructure required to run the capability outside of your cluster. This
// approach reduces operational overhead and preserves cluster resources.
//
// You can only create one Capability of each type on a given Amazon EKS cluster.
// Valid types are Argo CD for declarative GitOps deployment, Amazon Web Services
// Controllers for Kubernetes (ACK) for resource management, and Kube Resource
// Orchestrator (KRO) for Kubernetes custom resource orchestration.
//
// For more information, see [EKS Capabilities] in the Amazon EKS User Guide.
//
// [EKS Capabilities]: https://docs.aws.amazon.com/eks/latest/userguide/capabilities.html
