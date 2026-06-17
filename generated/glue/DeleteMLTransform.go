package glue

// DeleteMLTransform is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Deletes an Glue machine learning transform. Machine learning transforms are a
// special type of transform that use machine learning to learn the details of the
// transformation to be performed by learning from examples provided by humans.
// These transformations are then saved by Glue. If you no longer need a transform,
// you can delete it by calling DeleteMLTransforms . However, any Glue jobs that
// still reference the deleted transform will no longer succeed.
