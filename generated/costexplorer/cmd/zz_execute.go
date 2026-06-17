package cmd

func Execute(args []string) error {
	if p := _costexplorerCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_costexplorerCmd.Name()}, args...))
		return p.Execute()
	}
	_costexplorerCmd.SetArgs(args)
	return _costexplorerCmd.Execute()
}
