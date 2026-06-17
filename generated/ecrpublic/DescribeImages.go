package ecrpublic

// DescribeImages is generated as a reference stub.
// Executable command wiring lives under cmd/ecrpublic.go.
//
// Returns metadata that's related to the images in a repository in a public
// registry.
//
// Beginning with Docker version 1.9, the Docker client compresses image layers
// before pushing them to a V2 Docker registry. The output of the docker images
// command shows the uncompressed image size. Therefore, it might return a larger
// image size than the image sizes that are returned by DescribeImages.
