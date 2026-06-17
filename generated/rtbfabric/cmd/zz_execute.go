package cmd

func Execute(args []string) error {
	if p := _rtbfabricCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_rtbfabricCmd.Name()}, args...))
		return p.Execute()
	}
	_rtbfabricCmd.SetArgs(args)
	return _rtbfabricCmd.Execute()
}
