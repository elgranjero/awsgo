package cmd

func Execute(args []string) error {
	if p := _cloudwatchlogsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cloudwatchlogsCmd.Name()}, args...))
		return p.Execute()
	}
	_cloudwatchlogsCmd.SetArgs(args)
	return _cloudwatchlogsCmd.Execute()
}
