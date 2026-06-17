package iotsitewise

// DescribeGatewayCapabilityConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/iotsitewise.go.
//
// Each gateway capability defines data sources for a gateway. This is the
// namespace of the gateway capability.
//
// . The namespace follows the format service:capability:version , where:
//
// - service - The service providing the capability, or iotsitewise .
//
// - capability - The specific capability type. Options include: opcuacollector
// for the OPC UA data source collector, or publisher for data publisher
// capability.
//
// - version - The version number of the capability. Option include 2 for Classic
// streams, V2 gateways, and 3 for MQTT-enabled, V3 gateways.
//
// After updating a capability configuration, the sync status becomes OUT_OF_SYNC
// until the gateway processes the configuration.Use
// DescribeGatewayCapabilityConfiguration to check the sync status and verify the
// configuration was applied.
//
// A gateway can have multiple capability configurations with different namespaces.
