package cmd

func Execute(args []string) error {
	if p := _repostspaceCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_repostspaceCmd.Name()}, args...))
		return p.Execute()
	}
	_repostspaceCmd.SetArgs(args)
	return _repostspaceCmd.Execute()
}
