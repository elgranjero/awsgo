package cmd

func Execute(args []string) error {
	if p := _greengrassCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_greengrassCmd.Name()}, args...))
		return p.Execute()
	}
	_greengrassCmd.SetArgs(args)
	return _greengrassCmd.Execute()
}
