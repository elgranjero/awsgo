package cmd

func Execute(args []string) error {
	if p := _codepipelineCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_codepipelineCmd.Name()}, args...))
		return p.Execute()
	}
	_codepipelineCmd.SetArgs(args)
	return _codepipelineCmd.Execute()
}
