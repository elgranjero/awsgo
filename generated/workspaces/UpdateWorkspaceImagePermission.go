package workspaces

// UpdateWorkspaceImagePermission is generated as a reference stub.
// Executable command wiring lives under cmd/workspaces.go.
//
// Shares or unshares an image with one account in the same Amazon Web Services
// Region by specifying whether that account has permission to copy the image. If
// the copy image permission is granted, the image is shared with that account. If
// the copy image permission is revoked, the image is unshared with the account.
//
// After an image has been shared, the recipient account can copy the image to
// other Regions as needed.
//
// In the China (Ningxia) Region, you can copy images only within the same Region.
//
// In Amazon Web Services GovCloud (US), to copy images to and from other Regions,
// contact Amazon Web Services Support.
//
// For more information about sharing images, see [Share or Unshare a Custom WorkSpaces Image].
//
// - To delete an image that has been shared, you must unshare the image before
// you delete it.
//
// - Sharing Bring Your Own License (BYOL) images across Amazon Web Services
// accounts isn't supported at this time in Amazon Web Services GovCloud (US). To
// share BYOL images across accounts in Amazon Web Services GovCloud (US), contact
// Amazon Web Services Support.
//
// [Share or Unshare a Custom WorkSpaces Image]: https://docs.aws.amazon.com/workspaces/latest/adminguide/share-custom-image.html
