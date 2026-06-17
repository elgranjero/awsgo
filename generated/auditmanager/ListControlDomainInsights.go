package auditmanager

// ListControlDomainInsights is generated as a reference stub.
// Executable command wiring lives under cmd/auditmanager.go.
//
// Lists the latest analytics data for control domains across all of your active
// assessments.
//
// Audit Manager supports the control domains that are provided by Amazon Web
// Services Control Catalog. For information about how to find a list of available
// control domains, see [ListDomains]ListDomains in the Amazon Web Services Control Catalog API
// Reference.
//
// A control domain is listed only if at least one of the controls within that
// domain collected evidence on the lastUpdated date of controlDomainInsights . If
// this condition isn’t met, no data is listed for that control domain.
//
// [ListDomains]: https://docs.aws.amazon.com/controlcatalog/latest/APIReference/API_ListDomains.html
