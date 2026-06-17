package cmd

func Execute(args []string) error {
	if p := _neptunegraphCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_neptunegraphCmd.Name()}, args...))
		return p.Execute()
	}
	_neptunegraphCmd.SetArgs(args)
	return _neptunegraphCmd.Execute()
}
