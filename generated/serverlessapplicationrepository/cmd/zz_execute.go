package cmd

func Execute(args []string) error {
	if p := _serverlessapplicationrepositoryCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_serverlessapplicationrepositoryCmd.Name()}, args...))
		return p.Execute()
	}
	_serverlessapplicationrepositoryCmd.SetArgs(args)
	return _serverlessapplicationrepositoryCmd.Execute()
}
