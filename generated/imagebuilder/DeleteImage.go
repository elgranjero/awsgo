package imagebuilder

// DeleteImage is generated as a reference stub.
// Executable command wiring lives under cmd/imagebuilder.go.
//
// Deletes an Image Builder image resource. This does not delete any EC2 AMIs or
// ECR container images that are created during the image build process. You must
// clean those up separately, using the appropriate Amazon EC2 or Amazon ECR
// console actions, or API or CLI commands.
//
// - To deregister an EC2 Linux AMI, see [Deregister your Linux AMI]in the Amazon EC2 User Guide .
//
// - To deregister an EC2 Windows AMI, see [Deregister your Windows AMI]in the Amazon EC2 Windows Guide .
//
// - To delete a container image from Amazon ECR, see [Deleting an image]in the Amazon ECR User
// Guide.
//
// [Deregister your Linux AMI]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/deregister-ami.html
// [Deregister your Windows AMI]: https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/deregister-ami.html
// [Deleting an image]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/delete_image.html
