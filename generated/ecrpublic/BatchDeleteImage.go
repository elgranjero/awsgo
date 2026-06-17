package ecrpublic

// BatchDeleteImage is generated as a reference stub.
// Executable command wiring lives under cmd/ecrpublic.go.
//
// Deletes a list of specified images that are within a repository in a public
// registry. Images are specified with either an imageTag or imageDigest .
//
// You can remove a tag from an image by specifying the image's tag in your
// request. When you remove the last tag from an image, the image is deleted from
// your repository.
//
// You can completely delete an image (and all of its tags) by specifying the
// digest of the image in your request.
