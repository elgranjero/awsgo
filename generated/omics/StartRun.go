package omics

// StartRun is generated as a reference stub.
// Executable command wiring lives under cmd/omics.go.
//
// Starts a new run and returns details about the run, or duplicates an existing
// run. A run is a single invocation of a workflow. If you provide request IDs,
// Amazon Web Services HealthOmics identifies duplicate requests and starts the run
// only once. Monitor the progress of the run by calling the GetRun API operation.
//
// To start a new run, the following inputs are required:
//
// - A service role ARN ( roleArn ).
//
// - The run's workflow ID ( workflowId , not the uuid or runId ).
//
// - An Amazon S3 location ( outputUri ) where the run outputs will be saved.
//
// - All required workflow parameters ( parameter ), which can include optional
// parameters from the parameter template. The run cannot include any parameters
// that are not defined in the parameter template. To see all possible parameters,
// use the GetRun API operation.
//
// - For runs with a STATIC (default) storage type, specify the required storage
// capacity (in gibibytes). A storage capacity value is not required for runs that
// use DYNAMIC storage.
//
// StartRun can also duplicate an existing run using the run's default values. You
// can modify these default values and/or add other optional inputs. To duplicate a
// run, the following inputs are required:
//
// - A service role ARN ( roleArn ).
//
// - The ID of the run to duplicate ( runId ).
//
// - An Amazon S3 location where the run outputs will be saved ( outputUri ).
//
// To learn more about the optional parameters for StartRun , see [Starting a run] in the Amazon
// Web Services HealthOmics User Guide.
//
// Use the retentionMode input to control how long the metadata for each run is
// stored in CloudWatch. There are two retention modes:
//
// - Specify REMOVE to automatically remove the oldest runs when you reach the
// maximum service retention limit for runs. It is recommended that you use the
// REMOVE mode to initiate major run requests so that your runs do not fail when
// you reach the limit.
//
// - The retentionMode is set to the RETAIN mode by default, which allows you to
// manually remove runs after reaching the maximum service retention limit. Under
// this setting, you cannot create additional runs until you remove the excess
// runs.
//
// To learn more about the retention modes, see [Run retention mode] in the Amazon Web Services
// HealthOmics User Guide.
//
// You can use Amazon Q CLI to analyze run logs and make performance optimization
// recommendations. To get started, see the [Amazon Web Services HealthOmics MCP server]on GitHub.
//
// [Starting a run]: https://docs.aws.amazon.com/omics/latest/dev/starting-a-run.html
// [Amazon Web Services HealthOmics MCP server]: https://github.com/awslabs/mcp/tree/main/src/aws-healthomics-mcp-server
// [Run retention mode]: https://docs.aws.amazon.com/omics/latest/dev/run-retention.html
