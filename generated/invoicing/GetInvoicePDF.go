package invoicing

// GetInvoicePDF is generated as a reference stub.
// Executable command wiring lives under cmd/invoicing.go.
//
// Returns a URL to download the invoice document and supplemental documents
// associated with an invoice. The URLs are pre-signed and have expiration time.
// For special cases like Brazil, where Amazon Web Services generated invoice
// identifiers and government provided identifiers do not match, use the Amazon Web
// Services generated invoice identifier when making API requests. To grant IAM
// permission to use this operation, the caller needs the invoicing:GetInvoicePDF
// policy action.
