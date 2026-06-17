package ssm

// DeleteAssociation is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// Disassociates the specified Amazon Web Services Systems Manager document (SSM
// document) from the specified managed node. If you created the association by
// using the Targets parameter, then you must delete the association by using the
// association ID.
//
// When you disassociate a document from a managed node, it doesn't change the
// configuration of the node. To change the configuration state of a managed node
// after you disassociate a document, you must create a new document with the
// desired configuration and associate it with the node.
