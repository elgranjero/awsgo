package b2bi

// GenerateMapping is generated as a reference stub.
// Executable command wiring lives under cmd/b2bi.go.
//
// Takes sample input and output documents and uses Amazon Bedrock to generate a
// mapping automatically. Depending on the accuracy and other factors, you can then
// edit the mapping for your needs.
//
// Before you can use the AI-assisted feature for Amazon Web Services B2B Data
// Interchange you must enable models in Amazon Bedrock. For details, see [AI-assisted template mapping prerequisites]in the
// Amazon Web Services B2B Data Interchange User guide.
//
// To generate a mapping, perform the following steps:
//
// - Start with an X12 EDI document to use as the input.
//
// - Call TestMapping using your EDI document.
//
// - Use the output from the TestMapping operation as either input or output for
// your GenerateMapping call, along with your sample file.
//
// [AI-assisted template mapping prerequisites]: https://docs.aws.amazon.com/b2bi/latest/userguide/ai-assisted-mapping.html#ai-assist-prereq
