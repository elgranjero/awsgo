package cmd

func Execute(args []string) error {
	if p := _networkmonitorCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_networkmonitorCmd.Name()}, args...))
		return p.Execute()
	}
	_networkmonitorCmd.SetArgs(args)
	return _networkmonitorCmd.Execute()
}
