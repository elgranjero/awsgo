package cmd

func Execute(args []string) error {
	if p := _cognitoidentityproviderCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cognitoidentityproviderCmd.Name()}, args...))
		return p.Execute()
	}
	_cognitoidentityproviderCmd.SetArgs(args)
	return _cognitoidentityproviderCmd.Execute()
}
