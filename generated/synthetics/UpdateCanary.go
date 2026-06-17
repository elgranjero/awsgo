package synthetics

// UpdateCanary is generated as a reference stub.
// Executable command wiring lives under cmd/synthetics.go.
//
// Updates the configuration of a canary that has already been created.
//
// For multibrowser canaries, you can add or remove browsers by updating the
// browserConfig list in the update call. For example:
//
// - To add Firefox to a canary that currently uses Chrome, specify
// browserConfigs as [CHROME, FIREFOX]
//
// - To remove Firefox and keep only Chrome, specify browserConfigs as [CHROME]
//
// You can't use this operation to update the tags of an existing canary. To
// change the tags of an existing canary, use [TagResource].
//
// When you use the dryRunId field when updating a canary, the only other field
// you can provide is the Schedule . Adding any other field will thrown an
// exception.
//
// [TagResource]: https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_TagResource.html
