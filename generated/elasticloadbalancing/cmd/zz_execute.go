package cmd

func Execute(args []string) error {
	if p := _elasticloadbalancingCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_elasticloadbalancingCmd.Name()}, args...))
		return p.Execute()
	}
	_elasticloadbalancingCmd.SetArgs(args)
	return _elasticloadbalancingCmd.Execute()
}
