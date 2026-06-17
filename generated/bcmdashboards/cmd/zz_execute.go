package cmd

func Execute(args []string) error {
	if p := _bcmdashboardsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_bcmdashboardsCmd.Name()}, args...))
		return p.Execute()
	}
	_bcmdashboardsCmd.SetArgs(args)
	return _bcmdashboardsCmd.Execute()
}
