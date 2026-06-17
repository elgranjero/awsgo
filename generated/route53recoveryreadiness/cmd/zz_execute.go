package cmd

func Execute(args []string) error {
	if p := _route53recoveryreadinessCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_route53recoveryreadinessCmd.Name()}, args...))
		return p.Execute()
	}
	_route53recoveryreadinessCmd.SetArgs(args)
	return _route53recoveryreadinessCmd.Execute()
}
