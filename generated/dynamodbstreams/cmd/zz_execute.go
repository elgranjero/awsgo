package cmd

func Execute(args []string) error {
	if p := _dynamodbstreamsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_dynamodbstreamsCmd.Name()}, args...))
		return p.Execute()
	}
	_dynamodbstreamsCmd.SetArgs(args)
	return _dynamodbstreamsCmd.Execute()
}
