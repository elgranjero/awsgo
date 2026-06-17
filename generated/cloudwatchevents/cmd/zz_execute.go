package cmd

func Execute(args []string) error {
	if p := _cloudwatcheventsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cloudwatcheventsCmd.Name()}, args...))
		return p.Execute()
	}
	_cloudwatcheventsCmd.SetArgs(args)
	return _cloudwatcheventsCmd.Execute()
}
