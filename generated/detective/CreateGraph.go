package detective

// CreateGraph is generated as a reference stub.
// Executable command wiring lives under cmd/detective.go.
//
// Creates a new behavior graph for the calling account, and sets that account as
// the administrator account. This operation is called by the account that is
// enabling Detective.
//
// The operation also enables Detective for the calling account in the currently
// selected Region. It returns the ARN of the new behavior graph.
//
// CreateGraph triggers a process to create the corresponding data tables for the
// new behavior graph.
//
// An account can only be the administrator account for one behavior graph within
// a Region. If the same account calls CreateGraph with the same administrator
// account, it always returns the same behavior graph ARN. It does not create a new
// behavior graph.
