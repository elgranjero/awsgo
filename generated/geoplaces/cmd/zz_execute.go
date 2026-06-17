package cmd

func Execute(args []string) error {
	if p := _geoplacesCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_geoplacesCmd.Name()}, args...))
		return p.Execute()
	}
	_geoplacesCmd.SetArgs(args)
	return _geoplacesCmd.Execute()
}
