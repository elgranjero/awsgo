package fsx

// AssociateFileSystemAliases is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Use this action to associate one or more Domain Name Server (DNS) aliases with
// an existing Amazon FSx for Windows File Server file system. A file system can
// have a maximum of 50 DNS aliases associated with it at any one time. If you try
// to associate a DNS alias that is already associated with the file system, FSx
// takes no action on that alias in the request. For more information, see [Working with DNS Aliases]and [Walkthrough 5: Using DNS aliases to access your file system],
// including additional steps you must take to be able to access your file system
// using a DNS alias.
//
// The system response shows the DNS aliases that Amazon FSx is attempting to
// associate with the file system. Use the API operation to monitor the status of
// the aliases Amazon FSx is associating with the file system.
//
// [Walkthrough 5: Using DNS aliases to access your file system]: https://docs.aws.amazon.com/fsx/latest/WindowsGuide/walkthrough05-file-system-custom-CNAME.html
// [Working with DNS Aliases]: https://docs.aws.amazon.com/fsx/latest/WindowsGuide/managing-dns-aliases.html
