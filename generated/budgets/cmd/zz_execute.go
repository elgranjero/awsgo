package cmd

func Execute(args []string) error {
	if p := _budgetsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_budgetsCmd.Name()}, args...))
		return p.Execute()
	}
	_budgetsCmd.SetArgs(args)
	return _budgetsCmd.Execute()
}
