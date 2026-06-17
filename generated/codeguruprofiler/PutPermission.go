package codeguruprofiler

// PutPermission is generated as a reference stub.
// Executable command wiring lives under cmd/codeguruprofiler.go.
//
// Adds permissions to a profiling group's resource-based policy that are
//
// provided using an action group. If a profiling group doesn't have a
// resource-based policy, one is created for it using the permissions in the action
// group and the roles and users in the principals parameter.
//
// The one supported action group that can be added is agentPermission which
// grants ConfigureAgent and PostAgent permissions. For more information, see [Resource-based policies in CodeGuru Profiler] in
// the Amazon CodeGuru Profiler User Guide, [ConfigureAgent]ConfigureAgent , and [PostAgentProfile]PostAgentProfile .
//
// The first time you call PutPermission on a profiling group, do not specify a
// revisionId because it doesn't have a resource-based policy. Subsequent calls
// must provide a revisionId to specify which revision of the resource-based
// policy to add the permissions to.
//
// The response contains the profiling group's JSON-formatted resource policy.
//
// [ConfigureAgent]: https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ConfigureAgent.html
// [Resource-based policies in CodeGuru Profiler]: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/resource-based-policies.html
// [PostAgentProfile]: https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_PostAgentProfile.html
