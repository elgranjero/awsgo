package codeartifact

// DisposePackageVersions is generated as a reference stub.
// Executable command wiring lives under cmd/codeartifact.go.
//
// Deletes the assets in package versions and sets the package versions' status
//
// to Disposed . A disposed package version cannot be restored in your repository
// because its assets are deleted.
//
// To view all disposed package versions in a repository, use [ListPackageVersions] and set the [status]
// parameter to Disposed .
//
// To view information about a disposed package version, use [DescribePackageVersion].
//
// [DescribePackageVersion]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_DescribePackageVersion.html
// [ListPackageVersions]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html
// [status]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html#API_ListPackageVersions_RequestSyntax
