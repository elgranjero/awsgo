package acm

// AddTagsToCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/acm.go.
//
// Adds one or more tags to an ACM certificate. Tags are labels that you can use
// to identify and organize your Amazon Web Services resources. Each tag consists
// of a key and an optional value . You specify the certificate on input by its
// Amazon Resource Name (ARN). You specify the tag by using a key-value pair.
//
// You can apply a tag to just one certificate if you want to identify a specific
// characteristic of that certificate, or you can apply the same tag to multiple
// certificates if you want to filter for a common relationship among those
// certificates. Similarly, you can apply the same tag to multiple resources if you
// want to specify a relationship among those resources. For example, you can add
// the same tag to an ACM certificate and an Elastic Load Balancing load balancer
// to indicate that they are both used by the same website. For more information,
// see [Tagging ACM certificates].
//
// To remove one or more tags, use the RemoveTagsFromCertificate action. To view all of the tags that have
// been applied to the certificate, use the ListTagsForCertificateaction.
//
// [Tagging ACM certificates]: https://docs.aws.amazon.com/acm/latest/userguide/tags.html
