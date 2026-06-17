package cmd

func Execute(args []string) error {
	if p := _mwaaCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_mwaaCmd.Name()}, args...))
		return p.Execute()
	}
	_mwaaCmd.SetArgs(args)
	return _mwaaCmd.Execute()
}
