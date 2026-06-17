package cmd

func Execute(args []string) error {
	if p := _apigatewayCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_apigatewayCmd.Name()}, args...))
		return p.Execute()
	}
	_apigatewayCmd.SetArgs(args)
	return _apigatewayCmd.Execute()
}
