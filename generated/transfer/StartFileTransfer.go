package transfer

// StartFileTransfer is generated as a reference stub.
// Executable command wiring lives under cmd/transfer.go.
//
// Begins a file transfer between local Amazon Web Services storage and a remote
// AS2 or SFTP server.
//
// - For an AS2 connector, you specify the ConnectorId and one or more
// SendFilePaths to identify the files you want to transfer.
//
// - For an SFTP connector, the file transfer can be either outbound or inbound.
// In both cases, you specify the ConnectorId . Depending on the direction of the
// transfer, you also specify the following items:
//
// - If you are transferring file from a partner's SFTP server to Amazon Web
// Services storage, you specify one or more RetrieveFilePaths to identify the
// files you want to transfer, and a LocalDirectoryPath to specify the
// destination folder.
//
// - If you are transferring file to a partner's SFTP server from Amazon Web
// Services storage, you specify one or more SendFilePaths to identify the files
// you want to transfer, and a RemoteDirectoryPath to specify the destination
// folder.
