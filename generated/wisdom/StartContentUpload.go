package wisdom

// StartContentUpload is generated as a reference stub.
// Executable command wiring lives under cmd/wisdom.go.
//
// Get a URL to upload content to a knowledge base. To upload content, first make
// a PUT request to the returned URL with your file, making sure to include the
// required headers. Then use [CreateContent]to finalize the content creation process or [UpdateContent] to
// modify an existing resource. You can only upload content to a knowledge base of
// type CUSTOM.
//
// [CreateContent]: https://docs.aws.amazon.com/wisdom/latest/APIReference/API_CreateContent.html
// [UpdateContent]: https://docs.aws.amazon.com/wisdom/latest/APIReference/API_UpdateContent.html
