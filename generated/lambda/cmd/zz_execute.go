package cmd

func Execute(args []string) error {
	if p := _lambdaCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_lambdaCmd.Name()}, args...))
		return p.Execute()
	}
	_lambdaCmd.SetArgs(args)
	return _lambdaCmd.Execute()
}
