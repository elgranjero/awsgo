package transfer

// CreateConnector is generated as a reference stub.
// Executable command wiring lives under cmd/transfer.go.
//
// Creates the connector, which captures the parameters for a connection for the
// AS2 or SFTP protocol. For AS2, the connector is required for sending files to an
// externally hosted AS2 server. For SFTP, the connector is required when sending
// files to an SFTP server or receiving files from an SFTP server. For more details
// about connectors, see [Configure AS2 connectors]and [Create SFTP connectors].
//
// You must specify exactly one configuration object: either for AS2 ( As2Config )
// or SFTP ( SftpConfig ).
//
// [Configure AS2 connectors]: https://docs.aws.amazon.com/transfer/latest/userguide/configure-as2-connector.html
// [Create SFTP connectors]: https://docs.aws.amazon.com/transfer/latest/userguide/configure-sftp-connector.html
