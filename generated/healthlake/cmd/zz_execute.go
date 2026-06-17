package cmd

func Execute(args []string) error {
	if p := _healthlakeCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_healthlakeCmd.Name()}, args...))
		return p.Execute()
	}
	_healthlakeCmd.SetArgs(args)
	return _healthlakeCmd.Execute()
}
