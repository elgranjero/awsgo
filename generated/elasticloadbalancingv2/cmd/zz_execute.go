package cmd

func Execute(args []string) error {
	if p := _elasticloadbalancingv2Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_elasticloadbalancingv2Cmd.Name()}, args...))
		return p.Execute()
	}
	_elasticloadbalancingv2Cmd.SetArgs(args)
	return _elasticloadbalancingv2Cmd.Execute()
}
