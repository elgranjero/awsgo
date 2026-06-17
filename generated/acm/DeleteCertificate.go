package acm

// DeleteCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/acm.go.
//
// Deletes a certificate and its associated private key. If this action succeeds,
// the certificate no longer appears in the list that can be displayed by calling
// the ListCertificatesaction or be retrieved by calling the GetCertificate action. The certificate will not be
// available for use by Amazon Web Services services integrated with ACM.
//
// You cannot delete an ACM certificate that is being used by another Amazon Web
// Services service. To delete a certificate that is in use, the certificate
// association must first be removed.
