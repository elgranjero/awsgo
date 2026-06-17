package cmd

func Execute(args []string) error {
	if p := _redshiftCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_redshiftCmd.Name()}, args...))
		return p.Execute()
	}
	_redshiftCmd.SetArgs(args)
	return _redshiftCmd.Execute()
}
