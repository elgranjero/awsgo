package mediastore

// ListContainers is generated as a reference stub.
// Executable command wiring lives under cmd/mediastore.go.
//
// Lists the properties of all containers in AWS Elemental MediaStore.
//
// You can query to receive all the containers in one response. Or you can include
// the MaxResults parameter to receive a limited number of containers in each
// response. In this case, the response includes a token. To get the next set of
// containers, send the command again, this time with the NextToken parameter
// (with the returned token as its value). The next set of responses appears, with
// a token if there are still more containers to receive.
//
// See also DescribeContainer, which gets the properties of one container.
