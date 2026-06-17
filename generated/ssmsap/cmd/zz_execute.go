package cmd

func Execute(args []string) error {
	if p := _ssmsapCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ssmsapCmd.Name()}, args...))
		return p.Execute()
	}
	_ssmsapCmd.SetArgs(args)
	return _ssmsapCmd.Execute()
}
