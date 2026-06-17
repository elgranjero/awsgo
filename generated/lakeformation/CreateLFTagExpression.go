package lakeformation

// CreateLFTagExpression is generated as a reference stub.
// Executable command wiring lives under cmd/lakeformation.go.
//
// Creates a new LF-Tag expression with the provided name, description, catalog
// ID, and expression body. This call fails if a LF-Tag expression with the same
// name already exists in the caller’s account or if the underlying LF-Tags don't
// exist. To call this API operation, caller needs the following Lake Formation
// permissions:
//
// CREATE_LF_TAG_EXPRESSION on the root catalog resource.
//
// GRANT_WITH_LF_TAG_EXPRESSION on all underlying LF-Tag key:value pairs included
// in the expression.
