package workspaces

// CreateWorkspaces is generated as a reference stub.
// Executable command wiring lives under cmd/workspaces.go.
//
// Creates one or more WorkSpaces.
//
// This operation is asynchronous and returns before the WorkSpaces are created.
//
// - The MANUAL running mode value is only supported by Amazon WorkSpaces Core.
// Contact your account team to be allow-listed to use this value. For more
// information, see [Amazon WorkSpaces Core].
//
// - You don't need to specify the PCOIP protocol for Linux bundles because DCV
// (formerly WSP) is the default protocol for those bundles.
//
// - User-decoupled WorkSpaces are only supported by Amazon WorkSpaces Core.
//
// - Review your running mode to ensure you are using one that is optimal for
// your needs and budget. For more information on switching running modes, see [Can I switch between hourly and monthly billing?]
//
// [Can I switch between hourly and monthly billing?]: http://aws.amazon.com/workspaces-family/workspaces/faqs/#:~:text=Can%20I%20switch%20between%20hourly%20and%20monthly%20billing%20on%20WorkSpaces%20Personal%3F
// [Amazon WorkSpaces Core]: http://aws.amazon.com/workspaces/core/
