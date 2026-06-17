package networkfirewall

// UpdateLoggingConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/networkfirewall.go.
//
// Sets the logging configuration for the specified firewall.
//
// To change the logging configuration, retrieve the LoggingConfiguration by calling DescribeLoggingConfiguration, then change it
// and provide the modified object to this update call. You must change the logging
// configuration one LogDestinationConfigat a time inside the retrieved LoggingConfiguration object.
//
// You can perform only one of the following actions in any call to
// UpdateLoggingConfiguration :
//
// - Create a new log destination object by adding a single LogDestinationConfig
// array element to LogDestinationConfigs .
//
// - Delete a log destination object by removing a single LogDestinationConfig
// array element from LogDestinationConfigs .
//
// - Change the LogDestination setting in a single LogDestinationConfig array
// element.
//
// You can't change the LogDestinationType or LogType in a LogDestinationConfig .
// To change these settings, delete the existing LogDestinationConfig object and
// create a new one, using two separate calls to this update operation.
