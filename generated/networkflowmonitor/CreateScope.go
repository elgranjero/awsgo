package networkflowmonitor

// CreateScope is generated as a reference stub.
// Executable command wiring lives under cmd/networkflowmonitor.go.
//
// In Network Flow Monitor, you specify a scope for the service to generate
// metrics for. By using the scope, Network Flow Monitor can generate a topology of
// all the resources to measure performance metrics for. When you create a scope,
// you enable permissions for Network Flow Monitor.
//
// A scope is a Region-account pair or multiple Region-account pairs. Network Flow
// Monitor uses your scope to determine all the resources (the topology) where
// Network Flow Monitor will gather network flow performance metrics for you. To
// provide performance metrics, Network Flow Monitor uses the data that is sent by
// the Network Flow Monitor agents you install on the resources.
//
// To define the Region-account pairs for your scope, the Network Flow Monitor API
// uses the following constucts, which allow for future flexibility in defining
// scopes:
//
// - Targets, which are arrays of targetResources.
//
// - Target resources, which are Region-targetIdentifier pairs.
//
// - Target identifiers, made up of a targetID (currently always an account ID)
// and a targetType (currently always an account).
