package cmd

func Execute(args []string) error {
	if p := _iottwinmakerCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_iottwinmakerCmd.Name()}, args...))
		return p.Execute()
	}
	_iottwinmakerCmd.SetArgs(args)
	return _iottwinmakerCmd.Execute()
}
