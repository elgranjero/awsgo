package cmd

func Execute(args []string) error {
	if p := _computeoptimizerCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_computeoptimizerCmd.Name()}, args...))
		return p.Execute()
	}
	_computeoptimizerCmd.SetArgs(args)
	return _computeoptimizerCmd.Execute()
}
