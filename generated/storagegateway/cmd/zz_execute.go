package cmd

func Execute(args []string) error {
	if p := _storagegatewayCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_storagegatewayCmd.Name()}, args...))
		return p.Execute()
	}
	_storagegatewayCmd.SetArgs(args)
	return _storagegatewayCmd.Execute()
}
