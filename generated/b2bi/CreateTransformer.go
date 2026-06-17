package b2bi

// CreateTransformer is generated as a reference stub.
// Executable command wiring lives under cmd/b2bi.go.
//
// Creates a transformer. Amazon Web Services B2B Data Interchange currently
// supports two scenarios:
//
// - Inbound EDI: the Amazon Web Services customer receives an EDI file from
// their trading partner. Amazon Web Services B2B Data Interchange converts this
// EDI file into a JSON or XML file with a service-defined structure. A mapping
// template provided by the customer, in JSONata or XSLT format, is optionally
// applied to this file to produce a JSON or XML file with the structure the
// customer requires.
//
// - Outbound EDI: the Amazon Web Services customer has a JSON or XML file
// containing data that they wish to use in an EDI file. A mapping template,
// provided by the customer (in either JSONata or XSLT format) is applied to this
// file to generate a JSON or XML file in the service-defined structure. This file
// is then converted to an EDI file.
//
// The following fields are provided for backwards compatibility only: fileFormat ,
// mappingTemplate , ediType , and sampleDocument .
//
// - Use the mapping data type in place of mappingTemplate and fileFormat
//
// - Use the sampleDocuments data type in place of sampleDocument
//
// - Use either the inputConversion or outputConversion in place of ediType
