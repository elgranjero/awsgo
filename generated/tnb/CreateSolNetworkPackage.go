package tnb

// CreateSolNetworkPackage is generated as a reference stub.
// Executable command wiring lives under cmd/tnb.go.
//
// Creates a network package.
//
// A network package is a .zip file in CSAR (Cloud Service Archive) format defines
// the function packages you want to deploy and the Amazon Web Services
// infrastructure you want to deploy them on. For more information, see [Network instances]in the
// Amazon Web Services Telco Network Builder User Guide.
//
// A network package consists of a network service descriptor (NSD) file
// (required) and any additional files (optional), such as scripts specific to your
// needs. For example, if you have multiple function packages in your network
// package, you can use the NSD to define which network functions should run in
// certain VPCs, subnets, or EKS clusters.
//
// This request creates an empty network package container with an ID. Once you
// create a network package, you can upload the network package content using [PutSolNetworkPackageContent].
//
// [Network instances]: https://docs.aws.amazon.com/tnb/latest/ug/network-instances.html
// [PutSolNetworkPackageContent]: https://docs.aws.amazon.com/tnb/latest/APIReference/API_PutSolNetworkPackageContent.html
