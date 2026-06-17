package glue

// RegisterConnectionType is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Registers a custom connection type in Glue based on the configuration provided.
// This operation enables customers to configure custom connectors for any data
// source with REST-based APIs, eliminating the need for building custom Lambda
// connectors.
//
// The registered connection type stores details about how requests and responses
// are interpreted by REST sources, including connection properties, authentication
// configuration, and REST configuration with entity definitions. Once registered,
// customers can create connections using this connection type and work with them
// the same way as natively supported Glue connectors.
//
// Supports multiple authentication types including Basic, OAuth2 (Client
// Credentials, JWT Bearer, Authorization Code), and Custom Auth configurations.
