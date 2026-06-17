package workspaces

// UpdateConnectionAliasPermission is generated as a reference stub.
// Executable command wiring lives under cmd/workspaces.go.
//
// Shares or unshares a connection alias with one account by specifying whether
// that account has permission to associate the connection alias with a directory.
// If the association permission is granted, the connection alias is shared with
// that account. If the association permission is revoked, the connection alias is
// unshared with the account. For more information, see [Cross-Region Redirection for Amazon WorkSpaces].
//
// - Before performing this operation, call [DescribeConnectionAliases]to make sure that the current state
// of the connection alias is CREATED .
//
// - To delete a connection alias that has been shared, the shared account must
// first disassociate the connection alias from any directories it has been
// associated with. Then you must unshare the connection alias from the account it
// has been shared with. You can delete a connection alias only after it is no
// longer shared with any accounts or associated with any directories.
//
// [DescribeConnectionAliases]: https://docs.aws.amazon.com/workspaces/latest/api/API_DescribeConnectionAliases.html
// [Cross-Region Redirection for Amazon WorkSpaces]: https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html
