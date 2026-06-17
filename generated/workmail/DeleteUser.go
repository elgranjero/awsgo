package workmail

// DeleteUser is generated as a reference stub.
// Executable command wiring lives under cmd/workmail.go.
//
// Deletes a user from WorkMail and all subsequent systems. Before you can delete
// a user, the user state must be DISABLED . Use the DescribeUser action to confirm the user
// state.
//
// Deleting a user is permanent and cannot be undone. WorkMail archives user
// mailboxes for 30 days before they are permanently removed.
