package cmd

func Execute(args []string) error {
	if p := _batchCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_batchCmd.Name()}, args...))
		return p.Execute()
	}
	_batchCmd.SetArgs(args)
	return _batchCmd.Execute()
}
