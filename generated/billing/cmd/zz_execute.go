package cmd

func Execute(args []string) error {
	if p := _billingCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_billingCmd.Name()}, args...))
		return p.Execute()
	}
	_billingCmd.SetArgs(args)
	return _billingCmd.Execute()
}
