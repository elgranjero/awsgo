package codebuild

// CreateWebhook is generated as a reference stub.
// Executable command wiring lives under cmd/codebuild.go.
//
// For an existing CodeBuild build project that has its source code stored in a
// GitHub or Bitbucket repository, enables CodeBuild to start rebuilding the source
// code every time a code change is pushed to the repository.
//
// If you enable webhooks for an CodeBuild project, and the project is used as a
// build step in CodePipeline, then two identical builds are created for each
// commit. One build is triggered through webhooks, and one through CodePipeline.
// Because billing is on a per-build basis, you are billed for both builds.
// Therefore, if you are using CodePipeline, we recommend that you disable webhooks
// in CodeBuild. In the CodeBuild console, clear the Webhook box. For more
// information, see step 5 in [Change a Build Project's Settings].
//
// [Change a Build Project's Settings]: https://docs.aws.amazon.com/codebuild/latest/userguide/change-project.html#change-project-console
