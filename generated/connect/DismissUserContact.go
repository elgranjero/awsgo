package connect

// DismissUserContact is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Dismisses contacts from an agent’s CCP and returns the agent to an available
// state, which allows the agent to receive a new routed contact. Contacts can only
// be dismissed if they are in a MISSED , ERROR , ENDED , or REJECTED state in the [Agent Event Stream]
// .
//
// [Agent Event Stream]: https://docs.aws.amazon.com/connect/latest/adminguide/about-contact-states.html
