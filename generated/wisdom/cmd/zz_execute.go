package cmd

func Execute(args []string) error {
	if p := _wisdomCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_wisdomCmd.Name()}, args...))
		return p.Execute()
	}
	_wisdomCmd.SetArgs(args)
	return _wisdomCmd.Execute()
}
