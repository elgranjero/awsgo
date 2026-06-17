package cmd

func Execute(args []string) error {
	if p := _migrationhubstrategyCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_migrationhubstrategyCmd.Name()}, args...))
		return p.Execute()
	}
	_migrationhubstrategyCmd.SetArgs(args)
	return _migrationhubstrategyCmd.Execute()
}
