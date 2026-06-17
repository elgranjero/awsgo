package cmd

func Execute(args []string) error {
	if p := _migrationhubrefactorspacesCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_migrationhubrefactorspacesCmd.Name()}, args...))
		return p.Execute()
	}
	_migrationhubrefactorspacesCmd.SetArgs(args)
	return _migrationhubrefactorspacesCmd.Execute()
}
