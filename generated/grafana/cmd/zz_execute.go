package cmd

func Execute(args []string) error {
	if p := _grafanaCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_grafanaCmd.Name()}, args...))
		return p.Execute()
	}
	_grafanaCmd.SetArgs(args)
	return _grafanaCmd.Execute()
}
