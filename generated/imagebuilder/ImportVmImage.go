package imagebuilder

// ImportVmImage is generated as a reference stub.
// Executable command wiring lives under cmd/imagebuilder.go.
//
// When you export your virtual machine (VM) from its virtualization environment,
// that process creates a set of one or more disk container files that act as
// snapshots of your VM’s environment, settings, and data. The Amazon EC2 API [ImportImage]
// action uses those files to import your VM and create an AMI. To import using the
// CLI command, see [import-image]
//
// You can reference the task ID from the VM import to pull in the AMI that the
// import created as the base image for your Image Builder recipe.
//
// [ImportImage]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportImage.html
// [import-image]: https://docs.aws.amazon.com/cli/latest/reference/ec2/import-image.html
