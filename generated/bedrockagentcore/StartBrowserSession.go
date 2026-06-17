package bedrockagentcore

// StartBrowserSession is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagentcore.go.
//
// Creates and initializes a browser session in Amazon Bedrock AgentCore. The
// session enables agents to navigate and interact with web content, extract
// information from websites, and perform web-based tasks as part of their response
// generation.
//
// To create a session, you must specify a browser identifier and a name. You can
// also configure the viewport dimensions to control the visible area of web
// content. The session remains active until it times out or you explicitly stop it
// using the StopBrowserSession operation.
//
// The following operations are related to StartBrowserSession :
//
// [GetBrowserSession]
//
// [UpdateBrowserStream]
//
// [SaveBrowserSessionProfile]
//
// [StopBrowserSession]
//
// [SaveBrowserSessionProfile]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_SaveBrowserSessionProfile.html
// [UpdateBrowserStream]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_UpdateBrowserStream.html
// [GetBrowserSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_GetBrowserSession.html
// [StopBrowserSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_StopBrowserSession.html
