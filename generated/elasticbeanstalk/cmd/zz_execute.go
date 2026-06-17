package cmd

func Execute(args []string) error {
	if p := _elasticbeanstalkCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_elasticbeanstalkCmd.Name()}, args...))
		return p.Execute()
	}
	_elasticbeanstalkCmd.SetArgs(args)
	return _elasticbeanstalkCmd.Execute()
}
