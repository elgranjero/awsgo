package pcaconnectorscep

// CreateChallenge is generated as a reference stub.
// Executable command wiring lives under cmd/pcaconnectorscep.go.
//
// For general-purpose connectors. Creates a challenge password for the specified
// connector. The SCEP protocol uses a challenge password to authenticate a request
// before issuing a certificate from a certificate authority (CA). Your SCEP
// clients include the challenge password as part of their certificate request to
// Connector for SCEP. To retrieve the connector Amazon Resource Names (ARNs) for
// the connectors in your account, call [ListConnectors].
//
// To create additional challenge passwords for the connector, call CreateChallenge
// again. We recommend frequently rotating your challenge passwords.
//
// [ListConnectors]: https://docs.aws.amazon.com/pca-connector-scep/latest/APIReference/API_ListConnectors.html
