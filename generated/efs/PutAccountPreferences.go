package efs

// PutAccountPreferences is generated as a reference stub.
// Executable command wiring lives under cmd/efs.go.
//
// Use this operation to set the account preference in the current Amazon Web
// Services Region to use long 17 character (63 bit) or short 8 character (32 bit)
// resource IDs for new EFS file system and mount target resources. All existing
// resource IDs are not affected by any changes you make. You can set the ID
// preference during the opt-in period as EFS transitions to long resource IDs. For
// more information, see [Managing Amazon EFS resource IDs].
//
// Starting in October, 2021, you will receive an error if you try to set the
// account preference to use the short 8 character format resource ID. Contact
// Amazon Web Services support if you receive an error and must use short IDs for
// file system and mount target resources.
//
// [Managing Amazon EFS resource IDs]: https://docs.aws.amazon.com/efs/latest/ug/manage-efs-resource-ids.html
