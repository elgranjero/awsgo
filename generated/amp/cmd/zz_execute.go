package cmd

func Execute(args []string) error {
	if p := _ampCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ampCmd.Name()}, args...))
		return p.Execute()
	}
	_ampCmd.SetArgs(args)
	return _ampCmd.Execute()
}
