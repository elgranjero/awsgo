package cmd

func Execute(args []string) error {
	if p := _computeoptimizerautomationCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_computeoptimizerautomationCmd.Name()}, args...))
		return p.Execute()
	}
	_computeoptimizerautomationCmd.SetArgs(args)
	return _computeoptimizerautomationCmd.Execute()
}
