package devicefarm

// ListUniqueProblems is generated as a reference stub.
// Executable command wiring lives under cmd/devicefarm.go.
//
// Gets information about unique problems, such as exceptions or crashes.
//
// Unique problems are defined as a single instance of an error across a run, job,
// or suite. For example, if a call in your application consistently raises an
// exception ( OutOfBoundsException in MyActivity.java:386 ), ListUniqueProblems
// returns a single entry instead of many individual entries for that exception.
