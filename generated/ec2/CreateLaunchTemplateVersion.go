package ec2

// CreateLaunchTemplateVersion is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a new version of a launch template. You must specify an existing launch
// template, either by name or ID. You can determine whether the new version
// inherits parameters from a source version, and add or overwrite parameters as
// needed.
//
// Launch template versions are numbered in the order in which they are created.
// You can't specify, change, or replace the numbering of launch template versions.
//
// Launch templates are immutable; after you create a launch template, you can't
// modify it. Instead, you can create a new version of the launch template that
// includes the changes that you require.
//
// For more information, see [Modify a launch template (manage launch template versions)] in the Amazon EC2 User Guide.
//
// [Modify a launch template (manage launch template versions)]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/manage-launch-template-versions.html
