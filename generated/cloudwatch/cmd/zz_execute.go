package cmd

func Execute(args []string) error {
	if p := _cloudwatchCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cloudwatchCmd.Name()}, args...))
		return p.Execute()
	}
	_cloudwatchCmd.SetArgs(args)
	return _cloudwatchCmd.Execute()
}
