package cmd

func Execute(args []string) error {
	if p := _codebuildCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_codebuildCmd.Name()}, args...))
		return p.Execute()
	}
	_codebuildCmd.SetArgs(args)
	return _codebuildCmd.Execute()
}
