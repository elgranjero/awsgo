package elasticbeanstalk

// CreateApplicationVersion is generated as a reference stub.
// Executable command wiring lives under cmd/elasticbeanstalk.go.
//
// Creates an application version for the specified application. You can create an
// application version from a source bundle in Amazon S3, a commit in AWS
// CodeCommit, or the output of an AWS CodeBuild build as follows:
//
// Specify a commit in an AWS CodeCommit repository with SourceBuildInformation .
//
// Specify a build in an AWS CodeBuild with SourceBuildInformation and
// BuildConfiguration .
//
// # Specify a source bundle in S3 with SourceBundle
//
// Omit both SourceBuildInformation and SourceBundle to use the default sample
// application.
//
// After you create an application version with a specified Amazon S3 bucket and
// key location, you can't change that Amazon S3 location. If you change the Amazon
// S3 location, you receive an exception when you attempt to launch an environment
// from the application version.
