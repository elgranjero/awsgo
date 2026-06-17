package mgn

// DisconnectFromService is generated as a reference stub.
// Executable command wiring lives under cmd/mgn.go.
//
// Disconnects specific Source Servers from Application Migration Service. Data
// replication is stopped immediately. All AWS resources created by Application
// Migration Service for enabling the replication of these source servers will be
// terminated / deleted within 90 minutes. Launched Test or Cutover instances will
// NOT be terminated. If the agent on the source server has not been prevented from
// communicating with the Application Migration Service service, then it will
// receive a command to uninstall itself (within approximately 10 minutes). The
// following properties of the SourceServer will be changed immediately:
// dataReplicationInfo.dataReplicationState will be set to DISCONNECTED; The
// totalStorageBytes property for each of dataReplicationInfo.replicatedDisks will
// be set to zero; dataReplicationInfo.lagDuration and
// dataReplicationInfo.lagDuration will be nullified.
