package ecr

// StartImageScan is generated as a reference stub.
// Executable command wiring lives under cmd/ecr.go.
//
// Starts a basic image vulnerability scan.
//
// A basic image scan can only be started once per 24 hours on an individual
// image. This limit includes if an image was scanned on initial push. You can
// start up to 100,000 basic scans per 24 hours. This limit includes both scans on
// initial push and scans initiated by the StartImageScan API. For more
// information, see [Basic scanning]in the Amazon Elastic Container Registry User Guide.
//
// [Basic scanning]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning-basic.html
