package cmd

func Execute(args []string) error {
	if p := _accessanalyzerCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_accessanalyzerCmd.Name()}, args...))
		return p.Execute()
	}
	_accessanalyzerCmd.SetArgs(args)
	return _accessanalyzerCmd.Execute()
}
