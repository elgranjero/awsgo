package cmd

func Execute(args []string) error {
	if p := _applicationcostprofilerCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_applicationcostprofilerCmd.Name()}, args...))
		return p.Execute()
	}
	_applicationcostprofilerCmd.SetArgs(args)
	return _applicationcostprofilerCmd.Execute()
}
