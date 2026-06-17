package ec2

// CreateKeyPair is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates an ED25519 or 2048-bit RSA key pair with the specified name and in the
// specified format. Amazon EC2 stores the public key and displays the private key
// for you to save to a file. The private key is returned as an unencrypted PEM
// encoded PKCS#1 private key or an unencrypted PPK formatted private key for use
// with PuTTY. If a key with the specified name already exists, Amazon EC2 returns
// an error.
//
// The key pair returned to you is available only in the Amazon Web Services
// Region in which you create it. If you prefer, you can create your own key pair
// using a third-party tool and upload it to any Region using ImportKeyPair.
//
// You can have up to 5,000 key pairs per Amazon Web Services Region.
//
// For more information, see [Amazon EC2 key pairs] in the Amazon EC2 User Guide.
//
// [Amazon EC2 key pairs]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html
