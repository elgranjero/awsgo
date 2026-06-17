package codeartifact

// DeletePackageVersions is generated as a reference stub.
// Executable command wiring lives under cmd/codeartifact.go.
//
// Deletes one or more versions of a package. A deleted package version cannot be
//
// restored in your repository. If you want to remove a package version from your
// repository and be able to restore it later, set its status to Archived .
// Archived packages cannot be downloaded from a repository and don't show up with
// list package APIs (for example, [ListPackageVersions]), but you can restore them using [UpdatePackageVersionsStatus].
//
// [ListPackageVersions]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html
// [UpdatePackageVersionsStatus]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_UpdatePackageVersionsStatus.html
