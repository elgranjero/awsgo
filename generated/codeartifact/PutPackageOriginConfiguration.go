package codeartifact

// PutPackageOriginConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/codeartifact.go.
//
// Sets the package origin configuration for a package.
//
// The package origin configuration determines how new versions of a package can
// be added to a repository. You can allow or block direct publishing of new
// package versions, or ingestion and retaining of new package versions from an
// external connection or upstream source. For more information about package
// origin controls and configuration, see [Editing package origin controls]in the CodeArtifact User Guide.
//
// PutPackageOriginConfiguration can be called on a package that doesn't yet exist
// in the repository. When called on a package that does not exist, a package is
// created in the repository with no versions and the requested restrictions are
// set on the package. This can be used to preemptively block ingesting or
// retaining any versions from external connections or upstream repositories, or to
// block publishing any versions of the package into the repository before
// connecting any package managers or publishers to the repository.
//
// [Editing package origin controls]: https://docs.aws.amazon.com/codeartifact/latest/ug/package-origin-controls.html
