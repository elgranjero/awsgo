package securitylake

// DeleteAwsLogSource is generated as a reference stub.
// Executable command wiring lives under cmd/securitylake.go.
//
// Removes a natively supported Amazon Web Services service as an Amazon Security
// Lake source. You can remove a source for one or more Regions. When you remove
// the source, Security Lake stops collecting data from that source in the
// specified Regions and accounts, and subscribers can no longer consume new data
// from the source. However, subscribers can still consume data that Security Lake
// collected from the source before removal.
//
// You can choose any source type in any Amazon Web Services Region for either
// accounts that are part of a trusted organization or standalone accounts.
