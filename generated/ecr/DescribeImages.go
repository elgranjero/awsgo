package ecr

// DescribeImages is generated as a reference stub.
// Executable command wiring lives under cmd/ecr.go.
//
// Returns metadata about the images in a repository.
//
// Starting with Docker version 1.9, the Docker client compresses image layers
// before pushing them to a V2 Docker registry. The output of the docker images
// command shows the uncompressed image size. Therefore, Docker might return a
// larger image than the image shown in the Amazon Web Services Management Console.
//
// The new version of Amazon ECR Basic Scanning doesn't use the ImageDetail$imageScanFindingsSummary and ImageDetail$imageScanStatus attributes
// from the API response to return scan results. Use the DescribeImageScanFindingsAPI instead. For more
// information about Amazon Web Services native basic scanning, see [Scan images for software vulnerabilities in Amazon ECR].
//
// [Scan images for software vulnerabilities in Amazon ECR]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html
