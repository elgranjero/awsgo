package codeartifact

// GetAssociatedPackageGroup is generated as a reference stub.
// Executable command wiring lives under cmd/codeartifact.go.
//
// Returns the most closely associated package group to the specified package.
// This API does not require that the package exist in any repository in the
// domain. As such, GetAssociatedPackageGroup can be used to see which package
// group's origin configuration applies to a package before that package is in a
// repository. This can be helpful to check if public packages are blocked without
// ingesting them.
//
// For information package group association and matching, see [Package group definition syntax and matching behavior] in the
// CodeArtifact User Guide.
//
// [Package group definition syntax and matching behavior]: https://docs.aws.amazon.com/codeartifact/latest/ug/package-group-definition-syntax-matching-behavior.html
