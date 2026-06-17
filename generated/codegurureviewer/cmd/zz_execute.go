package cmd

func Execute(args []string) error {
	if p := _codegurureviewerCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_codegurureviewerCmd.Name()}, args...))
		return p.Execute()
	}
	_codegurureviewerCmd.SetArgs(args)
	return _codegurureviewerCmd.Execute()
}
