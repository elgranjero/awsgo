package cmd

func Execute(args []string) error {
	if p := _timestreamwriteCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_timestreamwriteCmd.Name()}, args...))
		return p.Execute()
	}
	_timestreamwriteCmd.SetArgs(args)
	return _timestreamwriteCmd.Execute()
}
