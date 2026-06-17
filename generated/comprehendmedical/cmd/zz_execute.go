package cmd

func Execute(args []string) error {
	if p := _comprehendmedicalCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_comprehendmedicalCmd.Name()}, args...))
		return p.Execute()
	}
	_comprehendmedicalCmd.SetArgs(args)
	return _comprehendmedicalCmd.Execute()
}
