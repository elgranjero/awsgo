package cmd

func Execute(args []string) error {
	if p := _panoramaCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_panoramaCmd.Name()}, args...))
		return p.Execute()
	}
	_panoramaCmd.SetArgs(args)
	return _panoramaCmd.Execute()
}
