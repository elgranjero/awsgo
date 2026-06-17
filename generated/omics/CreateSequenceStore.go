package omics

// CreateSequenceStore is generated as a reference stub.
// Executable command wiring lives under cmd/omics.go.
//
// Creates a sequence store and returns its metadata. Sequence stores are used to
// store sequence data files called read sets that are saved in FASTQ, BAM, uBAM,
// or CRAM formats. For aligned formats (BAM and CRAM), a sequence store can only
// use one reference genome. For unaligned formats (FASTQ and uBAM), a reference
// genome is not required. You can create multiple sequence stores per region per
// account.
//
// The following are optional parameters you can specify for your sequence store:
//
// - Use s3AccessConfig to configure your sequence store with S3 access logs
// (recommended).
//
// - Use sseConfig to define your own KMS key for encryption.
//
// - Use eTagAlgorithmFamily to define which algorithm to use for the HealthOmics
// eTag on objects.
//
// - Use fallbackLocation to define a backup location for storing files that have
// failed a direct upload.
//
// - Use propagatedSetLevelTags to configure tags that propagate to all objects
// in your store.
//
// For more information, see [Creating a HealthOmics sequence store] in the Amazon Web Services HealthOmics User Guide.
//
// [Creating a HealthOmics sequence store]: https://docs.aws.amazon.com/omics/latest/dev/create-sequence-store.html
