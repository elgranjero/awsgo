package cmd

func Execute(args []string) error {
	if p := _directoryservicedataCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_directoryservicedataCmd.Name()}, args...))
		return p.Execute()
	}
	_directoryservicedataCmd.SetArgs(args)
	return _directoryservicedataCmd.Execute()
}
