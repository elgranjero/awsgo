package cmd

func Execute(args []string) error {
	if p := _internetmonitorCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_internetmonitorCmd.Name()}, args...))
		return p.Execute()
	}
	_internetmonitorCmd.SetArgs(args)
	return _internetmonitorCmd.Execute()
}
