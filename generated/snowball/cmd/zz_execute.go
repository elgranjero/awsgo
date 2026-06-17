package cmd

func Execute(args []string) error {
	if p := _snowballCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_snowballCmd.Name()}, args...))
		return p.Execute()
	}
	_snowballCmd.SetArgs(args)
	return _snowballCmd.Execute()
}
