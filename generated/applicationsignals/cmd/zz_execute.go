package cmd

func Execute(args []string) error {
	if p := _applicationsignalsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_applicationsignalsCmd.Name()}, args...))
		return p.Execute()
	}
	_applicationsignalsCmd.SetArgs(args)
	return _applicationsignalsCmd.Execute()
}
