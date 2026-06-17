package cmd

func Execute(args []string) error {
	if p := _backupsearchCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_backupsearchCmd.Name()}, args...))
		return p.Execute()
	}
	_backupsearchCmd.SetArgs(args)
	return _backupsearchCmd.Execute()
}
