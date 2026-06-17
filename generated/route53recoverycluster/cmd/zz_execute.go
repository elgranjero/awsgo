package cmd

func Execute(args []string) error {
	if p := _route53recoveryclusterCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_route53recoveryclusterCmd.Name()}, args...))
		return p.Execute()
	}
	_route53recoveryclusterCmd.SetArgs(args)
	return _route53recoveryclusterCmd.Execute()
}
