package cmd

func Execute(args []string) error {
	if p := _accountCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_accountCmd.Name()}, args...))
		return p.Execute()
	}
	_accountCmd.SetArgs(args)
	return _accountCmd.Execute()
}
