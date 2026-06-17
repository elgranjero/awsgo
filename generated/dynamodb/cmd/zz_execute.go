package cmd

func Execute(args []string) error {
	if p := _dynamodbCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_dynamodbCmd.Name()}, args...))
		return p.Execute()
	}
	_dynamodbCmd.SetArgs(args)
	return _dynamodbCmd.Execute()
}
