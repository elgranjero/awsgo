package cmd

func Execute(args []string) error {
	if p := _appfabricCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_appfabricCmd.Name()}, args...))
		return p.Execute()
	}
	_appfabricCmd.SetArgs(args)
	return _appfabricCmd.Execute()
}
