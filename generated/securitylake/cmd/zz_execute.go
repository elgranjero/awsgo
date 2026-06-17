package cmd

func Execute(args []string) error {
	if p := _securitylakeCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_securitylakeCmd.Name()}, args...))
		return p.Execute()
	}
	_securitylakeCmd.SetArgs(args)
	return _securitylakeCmd.Execute()
}
