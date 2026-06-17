package cmd

func Execute(args []string) error {
	if p := _groundstationCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_groundstationCmd.Name()}, args...))
		return p.Execute()
	}
	_groundstationCmd.SetArgs(args)
	return _groundstationCmd.Execute()
}
