package cmd

func Execute(args []string) error {
	if p := _translateCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_translateCmd.Name()}, args...))
		return p.Execute()
	}
	_translateCmd.SetArgs(args)
	return _translateCmd.Execute()
}
