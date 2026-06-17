package cmd

func Execute(args []string) error {
	if p := _managedblockchainqueryCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_managedblockchainqueryCmd.Name()}, args...))
		return p.Execute()
	}
	_managedblockchainqueryCmd.SetArgs(args)
	return _managedblockchainqueryCmd.Execute()
}
