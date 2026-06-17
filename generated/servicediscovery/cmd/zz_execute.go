package cmd

func Execute(args []string) error {
	if p := _servicediscoveryCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_servicediscoveryCmd.Name()}, args...))
		return p.Execute()
	}
	_servicediscoveryCmd.SetArgs(args)
	return _servicediscoveryCmd.Execute()
}
