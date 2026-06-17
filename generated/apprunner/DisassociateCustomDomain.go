package apprunner

// DisassociateCustomDomain is generated as a reference stub.
// Executable command wiring lives under cmd/apprunner.go.
//
// Disassociate a custom domain name from an App Runner service.
//
// Certificates tracking domain validity are associated with a custom domain and
// are stored in [AWS Certificate Manager (ACM)]. These certificates aren't deleted as part of this action. App
// Runner delays certificate deletion for 30 days after a domain is disassociated
// from your service.
//
// [AWS Certificate Manager (ACM)]: https://docs.aws.amazon.com/acm/latest/userguide
