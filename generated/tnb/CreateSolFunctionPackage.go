package tnb

// CreateSolFunctionPackage is generated as a reference stub.
// Executable command wiring lives under cmd/tnb.go.
//
// Creates a function package.
//
// A function package is a .zip file in CSAR (Cloud Service Archive) format that
// contains a network function (an ETSI standard telecommunication application) and
// function package descriptor that uses the TOSCA standard to describe how the
// network functions should run on your network. For more information, see [Function packages]in the
// Amazon Web Services Telco Network Builder User Guide.
//
// Creating a function package is the first step for creating a network in AWS
// TNB. This request creates an empty container with an ID. The next step is to
// upload the actual CSAR zip file into that empty container. To upload function
// package content, see [PutSolFunctionPackageContent].
//
// [Function packages]: https://docs.aws.amazon.com/tnb/latest/ug/function-packages.html
// [PutSolFunctionPackageContent]: https://docs.aws.amazon.com/tnb/latest/APIReference/API_PutSolFunctionPackageContent.html
