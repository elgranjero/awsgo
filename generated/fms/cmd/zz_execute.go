package cmd

func Execute(args []string) error {
	if p := _fmsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_fmsCmd.Name()}, args...))
		return p.Execute()
	}
	_fmsCmd.SetArgs(args)
	return _fmsCmd.Execute()
}
