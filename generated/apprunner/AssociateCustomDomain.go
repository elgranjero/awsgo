package apprunner

// AssociateCustomDomain is generated as a reference stub.
// Executable command wiring lives under cmd/apprunner.go.
//
// Associate your own domain name with the App Runner subdomain URL of your App
// Runner service.
//
// After you call AssociateCustomDomain and receive a successful response, use the
// information in the CustomDomainrecord that's returned to add CNAME records to your Domain
// Name System (DNS). For each mapped domain name, add a mapping to the target App
// Runner subdomain and one or more certificate validation records. App Runner then
// performs DNS validation to verify that you own or control the domain name that
// you associated. App Runner tracks domain validity in a certificate stored in [AWS Certificate Manager (ACM)].
//
// [AWS Certificate Manager (ACM)]: https://docs.aws.amazon.com/acm/latest/userguide
