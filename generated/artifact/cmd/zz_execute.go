package cmd

func Execute(args []string) error {
	if p := _artifactCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_artifactCmd.Name()}, args...))
		return p.Execute()
	}
	_artifactCmd.SetArgs(args)
	return _artifactCmd.Execute()
}
