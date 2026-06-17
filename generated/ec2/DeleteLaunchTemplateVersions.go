package ec2

// DeleteLaunchTemplateVersions is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Deletes one or more versions of a launch template.
//
// You can't delete the default version of a launch template; you must first
// assign a different version as the default. If the default version is the only
// version for the launch template, you must delete the entire launch template
// using DeleteLaunchTemplate.
//
// You can delete up to 200 launch template versions in a single request. To
// delete more than 200 versions in a single request, use DeleteLaunchTemplate, which deletes the
// launch template and all of its versions.
//
// For more information, see [Delete a launch template version] in the Amazon EC2 User Guide.
//
// [Delete a launch template version]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/delete-launch-template.html#delete-launch-template-version
