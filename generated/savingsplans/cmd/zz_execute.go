package cmd

func Execute(args []string) error {
	if p := _savingsplansCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_savingsplansCmd.Name()}, args...))
		return p.Execute()
	}
	_savingsplansCmd.SetArgs(args)
	return _savingsplansCmd.Execute()
}
