package cmd

func Execute(args []string) error {
	if p := _cognitoidentityCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cognitoidentityCmd.Name()}, args...))
		return p.Execute()
	}
	_cognitoidentityCmd.SetArgs(args)
	return _cognitoidentityCmd.Execute()
}
