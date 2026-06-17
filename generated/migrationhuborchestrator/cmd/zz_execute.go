package cmd

func Execute(args []string) error {
	if p := _migrationhuborchestratorCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_migrationhuborchestratorCmd.Name()}, args...))
		return p.Execute()
	}
	_migrationhuborchestratorCmd.SetArgs(args)
	return _migrationhuborchestratorCmd.Execute()
}
