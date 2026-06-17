package directoryservice

// UpdateNumberOfDomainControllers is generated as a reference stub.
// Executable command wiring lives under cmd/directoryservice.go.
//
// Adds or removes domain controllers to or from the directory. Based on the
// difference between current value and new value (provided through this API call),
// domain controllers will be added or removed. It may take up to 45 minutes for
// any new domain controllers to become fully active once the requested number of
// domain controllers is updated. During this time, you cannot make another update
// request.
