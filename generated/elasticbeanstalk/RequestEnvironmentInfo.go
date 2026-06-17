package elasticbeanstalk

// RequestEnvironmentInfo is generated as a reference stub.
// Executable command wiring lives under cmd/elasticbeanstalk.go.
//
// Initiates a request to compile the specified type of information of the
// deployed environment.
//
// Setting the InfoType to tail compiles the last lines from the application
// server log files of every Amazon EC2 instance in your environment.
//
// Setting the InfoType to bundle compresses the application server log files for
// every Amazon EC2 instance into a .zip file. Legacy and .NET containers do not
// support bundle logs.
//
// Use RetrieveEnvironmentInfo to obtain the set of logs.
//
// # Related Topics
//
// RetrieveEnvironmentInfo
