package cmd

func Execute(args []string) error {
	if p := _geomapsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_geomapsCmd.Name()}, args...))
		return p.Execute()
	}
	_geomapsCmd.SetArgs(args)
	return _geomapsCmd.Execute()
}
