package mwaaserverless

// CreateWorkflow is generated as a reference stub.
// Executable command wiring lives under cmd/mwaaserverless.go.
//
// Creates a new workflow in Amazon Managed Workflows for Apache Airflow
// Serverless. This operation initializes a workflow with the specified
// configuration including the workflow definition, execution role, and optional
// settings for encryption, logging, and networking. You must provide the workflow
// definition as a YAML file stored in Amazon S3 that defines the DAG structure
// using supported Amazon Web Services operators. Amazon Managed Workflows for
// Apache Airflow Serverless automatically creates the first version of the
// workflow and sets up the necessary execution environment with multi-tenant
// isolation and security controls.
