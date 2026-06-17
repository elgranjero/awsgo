package cmd

func Execute(args []string) error {
	if p := _comprehendCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_comprehendCmd.Name()}, args...))
		return p.Execute()
	}
	_comprehendCmd.SetArgs(args)
	return _comprehendCmd.Execute()
}
