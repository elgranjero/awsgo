package cmd

func Execute(args []string) error {
	if p := _lakeformationCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_lakeformationCmd.Name()}, args...))
		return p.Execute()
	}
	_lakeformationCmd.SetArgs(args)
	return _lakeformationCmd.Execute()
}
