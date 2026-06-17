package cmd

func Execute(args []string) error {
	if p := _migrationhubconfigCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_migrationhubconfigCmd.Name()}, args...))
		return p.Execute()
	}
	_migrationhubconfigCmd.SetArgs(args)
	return _migrationhubconfigCmd.Execute()
}
