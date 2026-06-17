package b2bi

// CreateStarterMappingTemplate is generated as a reference stub.
// Executable command wiring lives under cmd/b2bi.go.
//
// Amazon Web Services B2B Data Interchange uses a mapping template in JSONata or
// XSLT format to transform a customer input file into a JSON or XML file that can
// be converted to EDI.
//
// If you provide a sample EDI file with the same structure as the EDI files that
// you wish to generate, then the service can generate a mapping template. The
// starter template contains placeholder values which you can replace with JSONata
// or XSLT expressions to take data from your input file and insert it into the
// JSON or XML file that is used to generate the EDI.
//
// If you do not provide a sample EDI file, then the service can generate a
// mapping template based on the EDI settings in the templateDetails parameter.
//
// Currently, we only support generating a template that can generate the input to
// produce an Outbound X12 EDI file.
