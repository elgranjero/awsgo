package codeartifact

// ListPackageVersionDependencies is generated as a reference stub.
// Executable command wiring lives under cmd/codeartifact.go.
//
// Returns the direct dependencies for a package version. The dependencies are
//
// returned as [PackageDependency]objects. CodeArtifact extracts the dependencies for a package
// version from the metadata file for the package format (for example, the
// package.json file for npm packages and the pom.xml file for Maven). Any package
// version dependencies that are not listed in the configuration file are not
// returned.
//
// [PackageDependency]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageDependency.html
