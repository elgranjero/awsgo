package greengrass

// StartBulkDeployment is generated as a reference stub.
// Executable command wiring lives under cmd/greengrass.go.
//
// Deploys multiple groups in one operation. This action starts the bulk
// deployment of a specified set of group versions. Each group version deployment
// will be triggered with an adaptive rate that has a fixed upper limit. We
// recommend that you include an ”X-Amzn-Client-Token” token in every
// ”StartBulkDeployment” request. These requests are idempotent with respect to
// the token and the request parameters.
