package imagebuilder

// ListComponents is generated as a reference stub.
// Executable command wiring lives under cmd/imagebuilder.go.
//
// Returns the list of components that can be filtered by name, or by using the
// listed filters to streamline results. Newly created components can take up to
// two minutes to appear in the ListComponents API Results.
//
// The semantic version has four nodes: ../. You can assign values for the first
// three, and can filter on all of them.
//
// Filtering: With semantic versioning, you have the flexibility to use wildcards
// (x) to specify the most recent versions or nodes when selecting the base image
// or components for your recipe. When you use a wildcard in any node, all nodes to
// the right of the first wildcard must also be wildcards.
