package gameliftstreams

// CreateApplication is generated as a reference stub.
// Executable command wiring lives under cmd/gameliftstreams.go.
//
// Creates an application resource in Amazon GameLift Streams, which specifies the
// application content you want to stream, such as a game build or other software,
// and configures the settings to run it.
//
// Before you create an application, upload your application content files to an
// Amazon Simple Storage Service (Amazon S3) bucket. For more information, see
// Getting Started in the Amazon GameLift Streams Developer Guide.
//
// Make sure that your files in the Amazon S3 bucket are the correct version you
// want to use. If you change the files at a later time, you will need to create a
// new Amazon GameLift Streams application.
//
// If the request is successful, Amazon GameLift Streams begins to create an
// application and sets the status to INITIALIZED . When an application reaches
// READY status, you can use the application to set up stream groups and start
// streams. To track application status, call [GetApplication].
//
// [GetApplication]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_GetApplication.html
