package iot

// CreateCertificateFromCsr is generated as a reference stub.
// Executable command wiring lives under cmd/iot.go.
//
// Creates an X.509 certificate using the specified certificate signing request.
//
// Requires permission to access the [CreateCertificateFromCsr] action.
//
// The CSR must include a public key that is either an RSA key with a length of at
// least 2048 bits or an ECC key from NIST P-256, NIST P-384, or NIST P-521 curves.
// For supported certificates, consult [Certificate signing algorithms supported by IoT].
//
// Reusing the same certificate signing request (CSR) results in a distinct
// certificate.
//
// You can create multiple certificates in a batch by creating a directory,
// copying multiple .csr files into that directory, and then specifying that
// directory on the command line. The following commands show how to create a batch
// of certificates given a batch of CSRs. In the following commands, we assume that
// a set of CSRs are located inside of the directory my-csr-directory:
//
// On Linux and OS X, the command is:
//
// $ ls my-csr-directory/ | xargs -I {} aws iot create-certificate-from-csr
// --certificate-signing-request file://my-csr-directory/{}
//
// This command lists all of the CSRs in my-csr-directory and pipes each CSR file
// name to the aws iot create-certificate-from-csr Amazon Web Services CLI command
// to create a certificate for the corresponding CSR.
//
// You can also run the aws iot create-certificate-from-csr part of the command in
// parallel to speed up the certificate creation process:
//
// $ ls my-csr-directory/ | xargs -P 10 -I {} aws iot create-certificate-from-csr
// --certificate-signing-request file://my-csr-directory/{}
//
// On Windows PowerShell, the command to create certificates for all CSRs in
// my-csr-directory is:
//
// > ls -Name my-csr-directory | %{aws iot create-certificate-from-csr
// --certificate-signing-request file://my-csr-directory/$_}
//
// On a Windows command prompt, the command to create certificates for all CSRs in
// my-csr-directory is:
//
// > forfiles /p my-csr-directory /c "cmd /c aws iot create-certificate-from-csr
// --certificate-signing-request file://(at)path"
//
// [CreateCertificateFromCsr]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [Certificate signing algorithms supported by IoT]: https://docs.aws.amazon.com/iot/latest/developerguide/x509-client-certs.html#x509-cert-algorithms
