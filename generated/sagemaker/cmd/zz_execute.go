package cmd

func Execute(args []string) error {
	if p := _sagemakerCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_sagemakerCmd.Name()}, args...))
		return p.Execute()
	}
	_sagemakerCmd.SetArgs(args)
	return _sagemakerCmd.Execute()
}
