package cmd

func Execute(args []string) error {
	if p := _rumCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_rumCmd.Name()}, args...))
		return p.Execute()
	}
	_rumCmd.SetArgs(args)
	return _rumCmd.Execute()
}
