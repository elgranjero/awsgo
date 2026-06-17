package servicecatalog

// DisableAWSOrganizationsAccess is generated as a reference stub.
// Executable command wiring lives under cmd/servicecatalog.go.
//
// Disable portfolio sharing through the Organizations service. This command will
// not delete your current shares, but prevents you from creating new shares
// throughout your organization. Current shares are not kept in sync with your
// organization structure if the structure changes after calling this API. Only the
// management account in the organization can call this API.
//
// You cannot call this API if there are active delegated administrators in the
// organization.
//
// Note that a delegated administrator is not authorized to invoke
// DisableAWSOrganizationsAccess .
//
// If you share an Service Catalog portfolio in an organization within
// Organizations, and then disable Organizations access for Service Catalog, the
// portfolio access permissions will not sync with the latest changes to the
// organization structure. Specifically, accounts that you removed from the
// organization after disabling Service Catalog access will retain access to the
// previously shared portfolio.
