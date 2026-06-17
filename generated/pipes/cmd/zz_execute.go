package cmd

func Execute(args []string) error {
	if p := _pipesCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_pipesCmd.Name()}, args...))
		return p.Execute()
	}
	_pipesCmd.SetArgs(args)
	return _pipesCmd.Execute()
}
