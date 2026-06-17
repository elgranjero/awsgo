package cmd

func Execute(args []string) error {
	if p := _sagemakerruntimeCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_sagemakerruntimeCmd.Name()}, args...))
		return p.Execute()
	}
	_sagemakerruntimeCmd.SetArgs(args)
	return _sagemakerruntimeCmd.Execute()
}
