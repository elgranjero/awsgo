package cmd

func Execute(args []string) error {
	if p := _appsyncCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_appsyncCmd.Name()}, args...))
		return p.Execute()
	}
	_appsyncCmd.SetArgs(args)
	return _appsyncCmd.Execute()
}
