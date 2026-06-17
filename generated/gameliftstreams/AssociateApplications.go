package gameliftstreams

// AssociateApplications is generated as a reference stub.
// Executable command wiring lives under cmd/gameliftstreams.go.
//
// When you associate, or link, an application with a stream group, then Amazon
// GameLift Streams can launch the application using the stream group's allocated
// compute resources. The stream group must be in ACTIVE status. You can reverse
// this action by using [DisassociateApplications].
//
// If a stream group does not already have a linked application, Amazon GameLift
// Streams will automatically assign the first application provided in
// ApplicationIdentifiers as the default.
//
// [DisassociateApplications]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_DisassociateApplications.html
