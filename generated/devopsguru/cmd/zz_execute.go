package cmd

func Execute(args []string) error {
	if p := _devopsguruCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_devopsguruCmd.Name()}, args...))
		return p.Execute()
	}
	_devopsguruCmd.SetArgs(args)
	return _devopsguruCmd.Execute()
}
