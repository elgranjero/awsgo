package ec2

// BundleInstance is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Bundles an Amazon instance store-backed Windows instance.
//
// During bundling, only the root device volume (C:\) is bundled. Data on other
// instance store volumes is not preserved.
//
// This action is not applicable for Linux/Unix instances or Windows instances
// that are backed by Amazon EBS.
