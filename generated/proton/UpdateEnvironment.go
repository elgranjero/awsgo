package proton

// UpdateEnvironment is generated as a reference stub.
// Executable command wiring lives under cmd/proton.go.
//
// Update an environment.
//
// If the environment is associated with an environment account connection, don't
// update or include the protonServiceRoleArn and provisioningRepository parameter
// to update or connect to an environment account connection.
//
// You can only update to a new environment account connection if that connection
// was created in the same environment account that the current environment account
// connection was created in. The account connection must also be associated with
// the current environment.
//
// If the environment isn't associated with an environment account connection,
// don't update or include the environmentAccountConnectionId parameter. You can't
// update or connect the environment to an environment account connection if it
// isn't already associated with an environment connection.
//
// You can update either the environmentAccountConnectionId or protonServiceRoleArn
// parameter and value. You can’t update both.
//
// If the environment was configured for Amazon Web Services-managed provisioning,
// omit the provisioningRepository parameter.
//
// If the environment was configured for self-managed provisioning, specify the
// provisioningRepository parameter and omit the protonServiceRoleArn and
// environmentAccountConnectionId parameters.
//
// For more information, see [Environments] and [Provisioning methods] in the Proton User Guide.
//
// There are four modes for updating an environment. The deploymentType field
// defines the mode.
//
// NONE
//
// In this mode, a deployment doesn't occur. Only the requested metadata
// parameters are updated.
//
// CURRENT_VERSION
//
// In this mode, the environment is deployed and updated with the new spec that
// you provide. Only requested parameters are updated. Don’t include minor or major
// version parameters when you use this deployment-type .
//
// MINOR_VERSION
//
// In this mode, the environment is deployed and updated with the published,
// recommended (latest) minor version of the current major version in use, by
// default. You can also specify a different minor version of the current major
// version in use.
//
// MAJOR_VERSION
//
// In this mode, the environment is deployed and updated with the published,
// recommended (latest) major and minor version of the current template, by
// default. You can also specify a different major version that's higher than the
// major version in use and a minor version.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Environments]: https://docs.aws.amazon.com/proton/latest/userguide/ag-environments.html
// [Provisioning methods]: https://docs.aws.amazon.com/proton/latest/userguide/ag-works-prov-methods.html
