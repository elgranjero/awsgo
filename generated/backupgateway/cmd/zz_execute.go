package cmd

func Execute(args []string) error {
	if p := _backupgatewayCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_backupgatewayCmd.Name()}, args...))
		return p.Execute()
	}
	_backupgatewayCmd.SetArgs(args)
	return _backupgatewayCmd.Execute()
}
