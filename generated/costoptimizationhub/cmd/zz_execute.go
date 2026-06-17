package cmd

func Execute(args []string) error {
	if p := _costoptimizationhubCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_costoptimizationhubCmd.Name()}, args...))
		return p.Execute()
	}
	_costoptimizationhubCmd.SetArgs(args)
	return _costoptimizationhubCmd.Execute()
}
