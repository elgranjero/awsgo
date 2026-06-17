package cmd

func Execute(args []string) error {
	if p := _locationCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_locationCmd.Name()}, args...))
		return p.Execute()
	}
	_locationCmd.SetArgs(args)
	return _locationCmd.Execute()
}
