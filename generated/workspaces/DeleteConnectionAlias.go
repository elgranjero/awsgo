package workspaces

// DeleteConnectionAlias is generated as a reference stub.
// Executable command wiring lives under cmd/workspaces.go.
//
// Deletes the specified connection alias. For more information, see [Cross-Region Redirection for Amazon WorkSpaces].
//
// If you will no longer be using a fully qualified domain name (FQDN) as the
// registration code for your WorkSpaces users, you must take certain precautions
// to prevent potential security issues. For more information, see [Security Considerations if You Stop Using Cross-Region Redirection].
//
// To delete a connection alias that has been shared, the shared account must
// first disassociate the connection alias from any directories it has been
// associated with. Then you must unshare the connection alias from the account it
// has been shared with. You can delete a connection alias only after it is no
// longer shared with any accounts or associated with any directories.
//
// [Security Considerations if You Stop Using Cross-Region Redirection]: https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html#cross-region-redirection-security-considerations
// [Cross-Region Redirection for Amazon WorkSpaces]: https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html
