package cmd

func Execute(args []string) error {
	if p := _managedblockchainCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_managedblockchainCmd.Name()}, args...))
		return p.Execute()
	}
	_managedblockchainCmd.SetArgs(args)
	return _managedblockchainCmd.Execute()
}
