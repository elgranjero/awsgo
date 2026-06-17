package cmd

func Execute(args []string) error {
	if p := _osisCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_osisCmd.Name()}, args...))
		return p.Execute()
	}
	_osisCmd.SetArgs(args)
	return _osisCmd.Execute()
}
