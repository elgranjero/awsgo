package cmd

func Execute(args []string) error {
	if p := _route53resolverCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_route53resolverCmd.Name()}, args...))
		return p.Execute()
	}
	_route53resolverCmd.SetArgs(args)
	return _route53resolverCmd.Execute()
}
