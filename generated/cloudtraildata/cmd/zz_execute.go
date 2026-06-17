package cmd

func Execute(args []string) error {
	if p := _cloudtraildataCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cloudtraildataCmd.Name()}, args...))
		return p.Execute()
	}
	_cloudtraildataCmd.SetArgs(args)
	return _cloudtraildataCmd.Execute()
}
