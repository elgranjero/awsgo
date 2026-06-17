package ec2

// DeleteVpnConnection is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Deletes the specified VPN connection.
//
// If you're deleting the VPC and its associated components, we recommend that you
// detach the virtual private gateway from the VPC and delete the VPC before
// deleting the VPN connection. If you believe that the tunnel credentials for your
// VPN connection have been compromised, you can delete the VPN connection and
// create a new one that has new keys, without needing to delete the VPC or virtual
// private gateway. If you create a new VPN connection, you must reconfigure the
// customer gateway device using the new configuration information returned with
// the new VPN connection ID.
//
// For certificate-based authentication, delete all Certificate Manager (ACM)
// private certificates used for the Amazon Web Services-side tunnel endpoints for
// the VPN connection before deleting the VPN connection.
