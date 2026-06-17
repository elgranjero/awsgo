package cmd

func Execute(args []string) error {
	if p := _databasemigrationserviceCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_databasemigrationserviceCmd.Name()}, args...))
		return p.Execute()
	}
	_databasemigrationserviceCmd.SetArgs(args)
	return _databasemigrationserviceCmd.Execute()
}
