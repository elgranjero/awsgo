package cmd

func Execute(args []string) error {
	if p := _drsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_drsCmd.Name()}, args...))
		return p.Execute()
	}
	_drsCmd.SetArgs(args)
	return _drsCmd.Execute()
}
