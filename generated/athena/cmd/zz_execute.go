package cmd

func Execute(args []string) error {
	if p := _athenaCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_athenaCmd.Name()}, args...))
		return p.Execute()
	}
	_athenaCmd.SetArgs(args)
	return _athenaCmd.Execute()
}
