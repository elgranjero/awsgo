package cmd

func Execute(args []string) error {
	if p := _supportCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_supportCmd.Name()}, args...))
		return p.Execute()
	}
	_supportCmd.SetArgs(args)
	return _supportCmd.Execute()
}
