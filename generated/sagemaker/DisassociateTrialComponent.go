package sagemaker

// DisassociateTrialComponent is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Disassociates a trial component from a trial. This doesn't effect other trials
// the component is associated with. Before you can delete a component, you must
// disassociate the component from all trials it is associated with. To associate a
// trial component with a trial, call the [AssociateTrialComponent]API.
//
// To get a list of the trials a component is associated with, use the [Search] API.
// Specify ExperimentTrialComponent for the Resource parameter. The list appears
// in the response under Results.TrialComponent.Parents .
//
// [Search]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_Search.html
// [AssociateTrialComponent]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AssociateTrialComponent.html
