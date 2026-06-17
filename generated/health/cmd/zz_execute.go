package cmd

func Execute(args []string) error {
	if p := _healthCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_healthCmd.Name()}, args...))
		return p.Execute()
	}
	_healthCmd.SetArgs(args)
	return _healthCmd.Execute()
}
