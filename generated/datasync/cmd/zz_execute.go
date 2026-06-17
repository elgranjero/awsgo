package cmd

func Execute(args []string) error {
	if p := _datasyncCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_datasyncCmd.Name()}, args...))
		return p.Execute()
	}
	_datasyncCmd.SetArgs(args)
	return _datasyncCmd.Execute()
}
