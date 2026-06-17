package dsql

// UpdateCluster is generated as a reference stub.
// Executable command wiring lives under cmd/dsql.go.
//
// The UpdateCluster API allows you to modify both single-Region and multi-Region
// cluster configurations. With the multiRegionProperties parameter, you can add or
// modify witness Region support and manage peer relationships with clusters in
// other Regions.
//
// Note that updating multi-Region clusters requires additional IAM permissions
// beyond those needed for standard cluster updates, as detailed in the Permissions
// section.
//
// # Required permissions
//
// dsql:UpdateCluster Permission to update a DSQL cluster.
//
// Resources: arn:aws:dsql:region:account-id:cluster/cluster-id
//
// dsql:PutMultiRegionProperties Permission to configure multi-Region properties
// for a cluster.
//
// Resources: arn:aws:dsql:region:account-id:cluster/cluster-id
//
// dsql:GetCluster Permission to retrieve cluster information.
//
// Resources: arn:aws:dsql:region:account-id:cluster/cluster-id
//
// dsql:AddPeerCluster Permission to add peer clusters.
//
// Resources:
//
// - Local cluster: arn:aws:dsql:region:account-id:cluster/cluster-id
//
// - Each peer cluster: exact ARN of each specified peer cluster
//
// dsql:RemovePeerCluster Permission to remove peer clusters. The
// dsql:RemovePeerCluster permission uses a wildcard ARN pattern to simplify
// permission management during updates.
//
// Resources: arn:aws:dsql:*:account-id:cluster/*
//
// dsql:PutWitnessRegion Permission to set a witness Region.
//
// Resources: arn:aws:dsql:region:account-id:cluster/cluster-id
//
// Condition Keys: dsql:WitnessRegion (matching the specified witness Region)
//
// This permission is checked both in the cluster Region and in the witness
// Region.
//
// - The witness region specified in multiRegionProperties.witnessRegion cannot
// be the same as the cluster's Region.
//
// - When updating clusters with peer relationships, permissions are checked for
// both adding and removing peers.
//
// - The dsql:RemovePeerCluster permission uses a wildcard ARN pattern to
// simplify permission management during updates.
