package ec2

// EnableImageBlockPublicAccess is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Enables block public access for AMIs at the account level in the specified
// Amazon Web Services Region. This prevents the public sharing of your AMIs.
// However, if you already have public AMIs, they will remain publicly available.
//
// The API can take up to 10 minutes to configure this setting. During this time,
// if you run [GetImageBlockPublicAccessState], the response will be unblocked . When the API has completed the
// configuration, the response will be block-new-sharing .
//
// For more information, see [Block public access to your AMIs] in the Amazon EC2 User Guide.
//
// [Block public access to your AMIs]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-public-access-to-amis.html
// [GetImageBlockPublicAccessState]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetImageBlockPublicAccessState.html
