package fms

// DisassociateAdminAccount is generated as a reference stub.
// Executable command wiring lives under cmd/fms.go.
//
// Disassociates an Firewall Manager administrator account. To set a different
// account as an Firewall Manager administrator, submit a PutAdminAccountrequest. To set an
// account as a default administrator account, you must submit an AssociateAdminAccountrequest.
//
// Disassociation of the default administrator account follows the first in, last
// out principle. If you are the default administrator, all Firewall Manager
// administrators within the organization must first disassociate their accounts
// before you can disassociate your account.
