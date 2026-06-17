package cmd

func Execute(args []string) error {
	if p := _entityresolutionCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_entityresolutionCmd.Name()}, args...))
		return p.Execute()
	}
	_entityresolutionCmd.SetArgs(args)
	return _entityresolutionCmd.Execute()
}
