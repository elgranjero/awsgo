package cmd

func Execute(args []string) error {
	if p := _finspacedataCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_finspacedataCmd.Name()}, args...))
		return p.Execute()
	}
	_finspacedataCmd.SetArgs(args)
	return _finspacedataCmd.Execute()
}
