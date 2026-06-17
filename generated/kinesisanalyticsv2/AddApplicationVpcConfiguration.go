package kinesisanalyticsv2

// AddApplicationVpcConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisanalyticsv2.go.
//
// Adds a Virtual Private Cloud (VPC) configuration to the application.
// Applications can use VPCs to store and access resources securely.
//
// Note the following about VPC configurations for Managed Service for Apache
// Flink applications:
//
// - VPC configurations are not supported for SQL applications.
//
// - When a VPC is added to a Managed Service for Apache Flink application, the
// application can no longer be accessed from the Internet directly. To enable
// Internet access to the application, add an Internet gateway to your VPC.
