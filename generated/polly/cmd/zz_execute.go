package cmd

func Execute(args []string) error {
	if p := _pollyCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_pollyCmd.Name()}, args...))
		return p.Execute()
	}
	_pollyCmd.SetArgs(args)
	return _pollyCmd.Execute()
}
