package codebuild

// UpdateProjectVisibility is generated as a reference stub.
// Executable command wiring lives under cmd/codebuild.go.
//
// Changes the public visibility for a project. The project's build results, logs,
// and artifacts are available to the general public. For more information, see [Public build projects]in
// the CodeBuild User Guide.
//
// The following should be kept in mind when making your projects public:
//
// - All of a project's build results, logs, and artifacts, including builds
// that were run when the project was private, are available to the general public.
//
// - All build logs and artifacts are available to the public. Environment
// variables, source code, and other sensitive information may have been output to
// the build logs and artifacts. You must be careful about what information is
// output to the build logs. Some best practice are:
//
// - Do not store sensitive values in environment variables. We recommend that
// you use an Amazon EC2 Systems Manager Parameter Store or Secrets Manager to
// store sensitive values.
//
// - Follow [Best practices for using webhooks]in the CodeBuild User Guide to limit which entities can trigger a
// build, and do not store the buildspec in the project itself, to ensure that your
// webhooks are as secure as possible.
//
// - A malicious user can use public builds to distribute malicious artifacts.
// We recommend that you review all pull requests to verify that the pull request
// is a legitimate change. We also recommend that you validate any artifacts with
// their checksums to make sure that the correct artifacts are being downloaded.
//
// [Public build projects]: https://docs.aws.amazon.com/codebuild/latest/userguide/public-builds.html
// [Best practices for using webhooks]: https://docs.aws.amazon.com/codebuild/latest/userguide/webhooks.html#webhook-best-practices
