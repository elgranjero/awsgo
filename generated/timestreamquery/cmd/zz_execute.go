package cmd

func Execute(args []string) error {
	if p := _timestreamqueryCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_timestreamqueryCmd.Name()}, args...))
		return p.Execute()
	}
	_timestreamqueryCmd.SetArgs(args)
	return _timestreamqueryCmd.Execute()
}
