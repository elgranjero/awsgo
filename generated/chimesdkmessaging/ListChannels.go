package chimesdkmessaging

// ListChannels is generated as a reference stub.
// Executable command wiring lives under cmd/chimesdkmessaging.go.
//
// Lists all Channels created under a single Chime App as a paginated list. You
// can specify filters to narrow results.
//
// Functionality & restrictions
//
// - Use privacy = PUBLIC to retrieve all public channels in the account.
//
// - Only an AppInstanceAdmin can set privacy = PRIVATE to list the private
// channels in an account.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
