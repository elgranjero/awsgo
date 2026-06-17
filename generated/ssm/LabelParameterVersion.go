package ssm

// LabelParameterVersion is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// A parameter label is a user-defined alias to help you manage different versions
// of a parameter. When you modify a parameter, Amazon Web Services Systems Manager
// automatically saves a new version and increments the version number by one. A
// label can help you remember the purpose of a parameter when there are multiple
// versions.
//
// Parameter labels have the following requirements and restrictions.
//
// - A version of a parameter can have a maximum of 10 labels.
//
// - You can't attach the same label to different versions of the same
// parameter. For example, if version 1 has the label Production, then you can't
// attach Production to version 2.
//
// - You can move a label from one version of a parameter to another.
//
// - You can't create a label when you create a new parameter. You must attach a
// label to a specific version of a parameter.
//
// - If you no longer want to use a parameter label, then you can either delete
// it or move it to a different version of a parameter.
//
// - A label can have a maximum of 100 characters.
//
// - Labels can contain letters (case sensitive), numbers, periods (.), hyphens
// (-), or underscores (_).
//
// - Labels can't begin with a number, " aws " or " ssm " (not case sensitive).
// If a label fails to meet these requirements, then the label isn't associated
// with a parameter and the system displays it in the list of InvalidLabels.
//
// - Parameter names can't contain spaces. The service removes any spaces
// specified for the beginning or end of a parameter name. If the specified name
// for a parameter contains spaces between characters, the request fails with a
// ValidationException error.
