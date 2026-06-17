package cmd

func Execute(args []string) error {
	if p := _connectcasesCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_connectcasesCmd.Name()}, args...))
		return p.Execute()
	}
	_connectcasesCmd.SetArgs(args)
	return _connectcasesCmd.Execute()
}
