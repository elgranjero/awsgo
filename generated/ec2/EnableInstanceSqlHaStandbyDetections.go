package ec2

// EnableInstanceSqlHaStandbyDetections is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Enable Amazon EC2 instances running in an SQL Server High Availability cluster
// for SQL Server High Availability instance standby detection monitoring. Once
// enabled, Amazon Web Services monitors the metadata for the instances to
// determine whether they are active or standby nodes in the SQL Server High
// Availability cluster. If the instances are determined to be standby failover
// nodes, Amazon Web Services automatically applies SQL Server licensing fee waiver
// for those instances.
//
// To register an instance, it must be running a Windows SQL Server
// license-included AMI and have the Amazon Web Services Systems Manager agent
// installed and running. Only Windows Server 2019 and later and SQL Server
// (Standard and Enterprise editions) 2017 and later are supported. For more
// information, see [Prerequisites for using SQL Server High Availability instance standby detection].
//
// [Prerequisites for using SQL Server High Availability instance standby detection]: https://docs.aws.amazon.com/sql-server-ec2/latest/userguide/prerequisites-and-requirements.html
