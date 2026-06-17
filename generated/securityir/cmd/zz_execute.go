package cmd

func Execute(args []string) error {
	if p := _securityirCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_securityirCmd.Name()}, args...))
		return p.Execute()
	}
	_securityirCmd.SetArgs(args)
	return _securityirCmd.Execute()
}
