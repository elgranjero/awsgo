package cmd

func Execute(args []string) error {
	if p := _sqsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_sqsCmd.Name()}, args...))
		return p.Execute()
	}
	_sqsCmd.SetArgs(args)
	return _sqsCmd.Execute()
}
