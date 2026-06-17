package lightsail

// CreateCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/lightsail.go.
//
// Creates an SSL/TLS certificate for an Amazon Lightsail content delivery network
// (CDN) distribution and a container service.
//
// After the certificate is valid, use the AttachCertificateToDistribution action
// to use the certificate and its domains with your distribution. Or use the
// UpdateContainerService action to use the certificate and its domains with your
// container service.
//
// Only certificates created in the us-east-1 Amazon Web Services Region can be
// attached to Lightsail distributions. Lightsail distributions are global
// resources that can reference an origin in any Amazon Web Services Region, and
// distribute its content globally. However, all distributions are located in the
// us-east-1 Region.
