package cmd

func Execute(args []string) error {
	if p := _daxCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_daxCmd.Name()}, args...))
		return p.Execute()
	}
	_daxCmd.SetArgs(args)
	return _daxCmd.Execute()
}
