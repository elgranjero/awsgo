package gameliftstreams

// ExportStreamSessionFiles is generated as a reference stub.
// Executable command wiring lives under cmd/gameliftstreams.go.
//
// Export the files that your application modifies or generates in a stream
//
// session, which can help you debug or verify your application. When your
// application runs, it generates output files such as logs, diagnostic
// information, crash dumps, save files, user data, screenshots, and so on. The
// files can be defined by the engine or frameworks that your application uses, or
// information that you've programmed your application to output.
//
// You can only call this action on a stream session that is in progress,
// specifically in one of the following statuses ACTIVE , CONNECTED ,
// PENDING_CLIENT_RECONNECTION , and RECONNECTING . You must provide an Amazon
// Simple Storage Service (Amazon S3) bucket to store the files in. When the
// session ends, Amazon GameLift Streams produces a compressed folder that contains
// all of the files and directories that were modified or created by the
// application during the stream session. AWS uses your security credentials to
// authenticate and authorize access to your Amazon S3 bucket.
//
// Amazon GameLift Streams collects the following generated and modified files.
// Find them in the corresponding folders in the .zip archive.
//
// - application/ : The folder where your application or game is stored.
//
// - profile/ : The user profile folder.
//
// - temp/ : The system temp folder.
//
// To verify the status of the exported files, use GetStreamSession.
//
// To delete the files, delete the object in the S3 bucket.
