package workspaces

// CreateUpdatedWorkspaceImage is generated as a reference stub.
// Executable command wiring lives under cmd/workspaces.go.
//
// Creates a new updated WorkSpace image based on the specified source image. The
// new updated WorkSpace image has the latest drivers and other updates required by
// the Amazon WorkSpaces components.
//
// To determine which WorkSpace images need to be updated with the latest Amazon
// WorkSpaces requirements, use [DescribeWorkspaceImages].
//
// - Only Windows 10, Windows Server 2016, and Windows Server 2019 WorkSpace
// images can be programmatically updated at this time.
//
// - Microsoft Windows updates and other application updates are not included in
// the update process.
//
// - The source WorkSpace image is not deleted. You can delete the source image
// after you've verified your new updated image and created a new bundle.
//
// [DescribeWorkspaceImages]: https://docs.aws.amazon.com/workspaces/latest/api/API_DescribeWorkspaceImages.html
