package cmd

func Execute(args []string) error {
	if p := _route53recoverycontrolconfigCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_route53recoverycontrolconfigCmd.Name()}, args...))
		return p.Execute()
	}
	_route53recoverycontrolconfigCmd.SetArgs(args)
	return _route53recoverycontrolconfigCmd.Execute()
}
