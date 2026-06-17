package cmd

func Execute(args []string) error {
	if p := _rdsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_rdsCmd.Name()}, args...))
		return p.Execute()
	}
	_rdsCmd.SetArgs(args)
	return _rdsCmd.Execute()
}
