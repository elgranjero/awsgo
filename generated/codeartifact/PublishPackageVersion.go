package codeartifact

// PublishPackageVersion is generated as a reference stub.
// Executable command wiring lives under cmd/codeartifact.go.
//
// Creates a new package version containing one or more assets (or files).
//
// The unfinished flag can be used to keep the package version in the Unfinished
// state until all of its assets have been uploaded (see [Package version status]in the CodeArtifact user
// guide). To set the package version’s status to Published , omit the unfinished
// flag when uploading the final asset, or set the status using [UpdatePackageVersionStatus]. Once a package
// version’s status is set to Published , it cannot change back to Unfinished .
//
// Only generic packages can be published using this API. For more information,
// see [Using generic packages]in the CodeArtifact User Guide.
//
// [Using generic packages]: https://docs.aws.amazon.com/codeartifact/latest/ug/using-generic.html
// [UpdatePackageVersionStatus]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_UpdatePackageVersionsStatus.html
// [Package version status]: https://docs.aws.amazon.com/codeartifact/latest/ug/packages-overview.html#package-version-status.html#package-version-status
