package eks

// DeleteCapability is generated as a reference stub.
// Executable command wiring lives under cmd/eks.go.
//
// Deletes a managed capability from your Amazon EKS cluster. When you delete a
// capability, Amazon EKS removes the capability infrastructure but retains all
// resources that were managed by the capability.
//
// Before deleting a capability, you should delete all Kubernetes resources that
// were created by the capability. After the capability is deleted, these resources
// become difficult to manage because the controller that managed them is no longer
// available. To delete resources before removing the capability, use kubectl
// delete or remove them through your GitOps workflow.
