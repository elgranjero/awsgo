package cmd

func Execute(args []string) error {
	if p := _apigatewayv2Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_apigatewayv2Cmd.Name()}, args...))
		return p.Execute()
	}
	_apigatewayv2Cmd.SetArgs(args)
	return _apigatewayv2Cmd.Execute()
}
