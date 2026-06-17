package gameliftstreams

// DisassociateApplications is generated as a reference stub.
// Executable command wiring lives under cmd/gameliftstreams.go.
//
// When you disassociate, or unlink, an application from a stream group, you can
//
// no longer stream this application by using that stream group's allocated compute
// resources. Any streams in process will continue until they terminate, which
// helps avoid interrupting an end-user's stream. Amazon GameLift Streams will not
// initiate new streams in the stream group using the disassociated application.
// The disassociate action does not affect the stream capacity of a stream group.
// To disassociate an application, the stream group must be in ACTIVE status.
//
// If you disassociate the default application, Amazon GameLift Streams will
// automatically choose a new default application from the remaining associated
// applications. To change which application is the default application, call [UpdateStreamGroup]and
// specify a new DefaultApplicationIdentifier .
//
// [UpdateStreamGroup]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_UpdateStreamGroup.html
