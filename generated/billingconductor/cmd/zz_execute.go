package cmd

func Execute(args []string) error {
	if p := _billingconductorCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_billingconductorCmd.Name()}, args...))
		return p.Execute()
	}
	_billingconductorCmd.SetArgs(args)
	return _billingconductorCmd.Execute()
}
