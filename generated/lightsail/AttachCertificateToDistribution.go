package lightsail

// AttachCertificateToDistribution is generated as a reference stub.
// Executable command wiring lives under cmd/lightsail.go.
//
// Attaches an SSL/TLS certificate to your Amazon Lightsail content delivery
// network (CDN) distribution.
//
// After the certificate is attached, your distribution accepts HTTPS traffic for
// all of the domains that are associated with the certificate.
//
// Use the CreateCertificate action to create a certificate that you can attach to
// your distribution.
//
// Only certificates created in the us-east-1 Amazon Web Services Region can be
// attached to Lightsail distributions. Lightsail distributions are global
// resources that can reference an origin in any Amazon Web Services Region, and
// distribute its content globally. However, all distributions are located in the
// us-east-1 Region.
