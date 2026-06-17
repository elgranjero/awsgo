package gameliftstreams

// DeleteApplication is generated as a reference stub.
// Executable command wiring lives under cmd/gameliftstreams.go.
//
// Permanently deletes an Amazon GameLift Streams application resource. This also
// deletes the application content files stored with Amazon GameLift Streams.
// However, this does not delete the original files that you uploaded to your
// Amazon S3 bucket; you can delete these any time after Amazon GameLift Streams
// creates an application, which is the only time Amazon GameLift Streams accesses
// your Amazon S3 bucket.
//
// You can only delete an application that meets the following conditions:
//
// - The application is in READY or ERROR status. You cannot delete an
// application that's in PROCESSING or INITIALIZED status.
//
// - The application is not the default application of any stream groups. You
// must first delete the stream group by using [DeleteStreamGroup].
//
// - The application is not linked to any stream groups. You must first unlink
// the stream group by using [DisassociateApplications].
//
// - An application is not streaming in any ongoing stream session. You must
// wait until the client ends the stream session or call [TerminateStreamSession]to end the stream.
//
// If any active stream groups exist for this application, this request returns a
// ValidationException .
//
// [DisassociateApplications]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_DisassociateApplications.html
// [TerminateStreamSession]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_TerminateStreamSession.html
// [DeleteStreamGroup]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_DeleteStreamGroup.html
