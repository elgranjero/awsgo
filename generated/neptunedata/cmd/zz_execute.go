package cmd

func Execute(args []string) error {
	if p := _neptunedataCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_neptunedataCmd.Name()}, args...))
		return p.Execute()
	}
	_neptunedataCmd.SetArgs(args)
	return _neptunedataCmd.Execute()
}
