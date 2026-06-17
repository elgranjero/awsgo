package cmd

func Execute(args []string) error {
	if p := _migrationhubCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_migrationhubCmd.Name()}, args...))
		return p.Execute()
	}
	_migrationhubCmd.SetArgs(args)
	return _migrationhubCmd.Execute()
}
