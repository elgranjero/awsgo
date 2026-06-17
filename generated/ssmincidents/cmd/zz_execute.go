package cmd

func Execute(args []string) error {
	if p := _ssmincidentsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ssmincidentsCmd.Name()}, args...))
		return p.Execute()
	}
	_ssmincidentsCmd.SetArgs(args)
	return _ssmincidentsCmd.Execute()
}
