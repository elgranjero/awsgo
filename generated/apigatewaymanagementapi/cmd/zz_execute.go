package cmd

func Execute(args []string) error {
	if p := _apigatewaymanagementapiCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_apigatewaymanagementapiCmd.Name()}, args...))
		return p.Execute()
	}
	_apigatewaymanagementapiCmd.SetArgs(args)
	return _apigatewaymanagementapiCmd.Execute()
}
