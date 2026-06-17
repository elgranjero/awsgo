package lightsail

// DeleteKeyPair is generated as a reference stub.
// Executable command wiring lives under cmd/lightsail.go.
//
// Deletes the specified key pair by removing the public key from Amazon Lightsail.
//
// You can delete key pairs that were created using the [ImportKeyPair] and [CreateKeyPair] actions, as well as
// the Lightsail default key pair. A new default key pair will not be created
// unless you launch an instance without specifying a custom key pair, or you call
// the [DownloadDefaultKeyPair]API.
//
// The delete key pair operation supports tag-based access control via resource
// tags applied to the resource identified by key pair name . For more information,
// see the [Amazon Lightsail Developer Guide].
//
// [ImportKeyPair]: https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_ImportKeyPair.html
// [CreateKeyPair]: https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_CreateKeyPair.html
// [DownloadDefaultKeyPair]: https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_DownloadDefaultKeyPair.html
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
