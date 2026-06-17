package cmd

func Execute(args []string) error {
	if p := _redshiftdataCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_redshiftdataCmd.Name()}, args...))
		return p.Execute()
	}
	_redshiftdataCmd.SetArgs(args)
	return _redshiftdataCmd.Execute()
}
