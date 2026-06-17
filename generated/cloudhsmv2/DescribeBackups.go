package cloudhsmv2

// DescribeBackups is generated as a reference stub.
// Executable command wiring lives under cmd/cloudhsmv2.go.
//
// Gets information about backups of CloudHSM clusters. Lists either the backups
// you own or the backups shared with you when the Shared parameter is true.
//
// This is a paginated operation, which means that each response might contain
// only a subset of all the backups. When the response contains only a subset of
// backups, it includes a NextToken value. Use this value in a subsequent
// DescribeBackups request to get more backups. When you receive a response with no
// NextToken (or an empty or null value), that means there are no more backups to
// get.
//
// Cross-account use: Yes. Customers can describe backups in other Amazon Web
// Services accounts that are shared with them.
