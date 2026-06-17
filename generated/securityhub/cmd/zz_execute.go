package cmd

func Execute(args []string) error {
	if p := _securityhubCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_securityhubCmd.Name()}, args...))
		return p.Execute()
	}
	_securityhubCmd.SetArgs(args)
	return _securityhubCmd.Execute()
}
