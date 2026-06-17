package cmd

func Execute(args []string) error {
	if p := _codecommitCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_codecommitCmd.Name()}, args...))
		return p.Execute()
	}
	_codecommitCmd.SetArgs(args)
	return _codecommitCmd.Execute()
}
