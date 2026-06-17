package cmd

func Execute(args []string) error {
	if p := _sagemakermetricsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_sagemakermetricsCmd.Name()}, args...))
		return p.Execute()
	}
	_sagemakermetricsCmd.SetArgs(args)
	return _sagemakermetricsCmd.Execute()
}
