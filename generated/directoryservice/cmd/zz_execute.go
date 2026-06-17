package cmd

func Execute(args []string) error {
	if p := _directoryserviceCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_directoryserviceCmd.Name()}, args...))
		return p.Execute()
	}
	_directoryserviceCmd.SetArgs(args)
	return _directoryserviceCmd.Execute()
}
