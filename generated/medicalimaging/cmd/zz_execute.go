package cmd

func Execute(args []string) error {
	if p := _medicalimagingCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_medicalimagingCmd.Name()}, args...))
		return p.Execute()
	}
	_medicalimagingCmd.SetArgs(args)
	return _medicalimagingCmd.Execute()
}
