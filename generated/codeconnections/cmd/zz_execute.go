package cmd

func Execute(args []string) error {
	if p := _codeconnectionsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_codeconnectionsCmd.Name()}, args...))
		return p.Execute()
	}
	_codeconnectionsCmd.SetArgs(args)
	return _codeconnectionsCmd.Execute()
}
