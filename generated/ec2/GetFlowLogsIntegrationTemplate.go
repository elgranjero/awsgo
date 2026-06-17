package ec2

// GetFlowLogsIntegrationTemplate is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Generates a CloudFormation template that streamlines and automates the
// integration of VPC flow logs with Amazon Athena. This make it easier for you to
// query and gain insights from VPC flow logs data. Based on the information that
// you provide, we configure resources in the template to do the following:
//
// - Create a table in Athena that maps fields to a custom log format
//
// - Create a Lambda function that updates the table with new partitions on a
// daily, weekly, or monthly basis
//
// - Create a table partitioned between two timestamps in the past
//
// - Create a set of named queries in Athena that you can use to get started
// quickly
//
// GetFlowLogsIntegrationTemplate does not support integration between Amazon Web
// Services Transit Gateway Flow Logs and Amazon Athena.
