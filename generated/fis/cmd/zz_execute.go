package cmd

func Execute(args []string) error {
	if p := _fisCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_fisCmd.Name()}, args...))
		return p.Execute()
	}
	_fisCmd.SetArgs(args)
	return _fisCmd.Execute()
}
