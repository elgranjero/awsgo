package ec2

// DescribeImages is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Describes the specified images (AMIs, AKIs, and ARIs) available to you or all
// of the images available to you.
//
// The images available to you include public images, private images that you own,
// and private images owned by other Amazon Web Services accounts for which you
// have explicit launch permissions.
//
// Recently deregistered images appear in the returned results for a short
// interval and then return empty results. After all instances that reference a
// deregistered AMI are terminated, specifying the ID of the image will eventually
// return an error indicating that the AMI ID cannot be found.
//
// When Allowed AMIs is set to enabled , only allowed images are returned in the
// results, with the imageAllowed field set to true for each image. In audit-mode ,
// the imageAllowed field is set to true for images that meet the account's
// Allowed AMIs criteria, and false for images that don't meet the criteria. For
// more information, see [Allowed AMIs].
//
// The Amazon EC2 API follows an eventual consistency model. This means that the
// result of an API command you run that creates or modifies resources might not be
// immediately available to all subsequent commands you run. For guidance on how to
// manage eventual consistency, see [Eventual consistency in the Amazon EC2 API]in the Amazon EC2 Developer Guide.
//
// We strongly recommend using only paginated requests. Unpaginated requests are
// susceptible to throttling and timeouts.
//
// The order of the elements in the response, including those within nested
// structures, might vary. Applications should not assume the elements appear in a
// particular order.
//
// [Allowed AMIs]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-allowed-amis.html
// [Eventual consistency in the Amazon EC2 API]: https://docs.aws.amazon.com/ec2/latest/devguide/eventual-consistency.html
