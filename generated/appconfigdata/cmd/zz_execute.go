package cmd

func Execute(args []string) error {
	if p := _appconfigdataCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_appconfigdataCmd.Name()}, args...))
		return p.Execute()
	}
	_appconfigdataCmd.SetArgs(args)
	return _appconfigdataCmd.Execute()
}
