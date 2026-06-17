package cmd

func Execute(args []string) error {
	if p := _eksauthCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_eksauthCmd.Name()}, args...))
		return p.Execute()
	}
	_eksauthCmd.SetArgs(args)
	return _eksauthCmd.Execute()
}
