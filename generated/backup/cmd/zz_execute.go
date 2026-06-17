package cmd

func Execute(args []string) error {
	if p := _backupCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_backupCmd.Name()}, args...))
		return p.Execute()
	}
	_backupCmd.SetArgs(args)
	return _backupCmd.Execute()
}
