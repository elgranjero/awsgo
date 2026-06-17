package mediastore

// PutCorsPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/mediastore.go.
//
// Sets the cross-origin resource sharing (CORS) configuration on a container so
// that the container can service cross-origin requests. For example, you might
// want to enable a request whose origin is http://www.example.com to access your
// AWS Elemental MediaStore container at my.example.container.com by using the
// browser's XMLHttpRequest capability.
//
// To enable CORS on a container, you attach a CORS policy to the container. In
// the CORS policy, you configure rules that identify origins and the HTTP methods
// that can be executed on your container. The policy can contain up to 398,000
// characters. You can add up to 100 rules to a CORS policy. If more than one rule
// applies, the service uses the first applicable rule listed.
//
// To learn more about CORS, see [Cross-Origin Resource Sharing (CORS) in AWS Elemental MediaStore].
//
// [Cross-Origin Resource Sharing (CORS) in AWS Elemental MediaStore]: https://docs.aws.amazon.com/mediastore/latest/ug/cors-policy.html
