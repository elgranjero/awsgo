package dsql

// CreateCluster is generated as a reference stub.
// Executable command wiring lives under cmd/dsql.go.
//
// The CreateCluster API allows you to create both single-Region clusters and
// multi-Region clusters. With the addition of the multiRegionProperties parameter,
// you can create a cluster with witness Region support and establish peer
// relationships with clusters in other Regions during creation.
//
// Creating multi-Region clusters requires additional IAM permissions beyond those
// needed for single-Region clusters, as detailed in the Required permissions
// section below.
//
// # Required permissions
//
// dsql:CreateCluster Required to create a cluster.
//
// Resources: arn:aws:dsql:region:account-id:cluster/*
//
// dsql:TagResource Permission to add tags to a resource.
//
// Resources: arn:aws:dsql:region:account-id:cluster/*
//
// dsql:PutMultiRegionProperties Permission to configure multi-Region properties
// for a cluster.
//
// Resources: arn:aws:dsql:region:account-id:cluster/*
//
// dsql:AddPeerCluster When specifying multiRegionProperties.clusters , permission
// to add peer clusters.
//
// Resources:
//
// - Local cluster: arn:aws:dsql:region:account-id:cluster/*
//
// - Each peer cluster: exact ARN of each specified peer cluster
//
// dsql:PutWitnessRegion When specifying multiRegionProperties.witnessRegion ,
// permission to set a witness Region. This permission is checked both in the
// cluster Region and in the witness Region.
//
// Resources: arn:aws:dsql:region:account-id:cluster/*
//
// Condition Keys: dsql:WitnessRegion (matching the specified witness region)
//
// - The witness Region specified in multiRegionProperties.witnessRegion cannot
// be the same as the cluster's Region.
