package inspectorscan

// ScanSbom is generated as a reference stub.
// Executable command wiring lives under cmd/inspectorscan.go.
//
// Scans a provided CycloneDX 1.5 SBOM and reports on any vulnerabilities
// discovered in that SBOM. You can generate compatible SBOMs for your resources
// using the [Amazon Inspector SBOM generator].
//
// The output of this action reports NVD and CVSS scores when NVD and CVSS scores
// are available. Because the output reports both scores, you might notice a
// discrepency between them. However, you can triage the severity of either score
// depending on the vendor of your choosing.
//
// [Amazon Inspector SBOM generator]: https://docs.aws.amazon.com/inspector/latest/user/sbom-generator.html
