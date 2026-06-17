package chime

// DeleteAccount is generated as a reference stub.
// Executable command wiring lives under cmd/chime.go.
//
// Deletes the specified Amazon Chime account. You must suspend all users before
// deleting Team account. You can use the BatchSuspendUser action to dodo.
//
// For EnterpriseLWA and EnterpriseAD accounts, you must release the claimed
// domains for your Amazon Chime account before deletion. As soon as you release
// the domain, all users under that account are suspended.
//
// Deleted accounts appear in your Disabled accounts list for 90 days. To restore
// deleted account from your Disabled accounts list, you must contact AWS Support.
//
// After 90 days, deleted accounts are permanently removed from your Disabled
// accounts list.
