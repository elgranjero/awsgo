package cmd

func Execute(args []string) error {
	if p := _route53globalresolverCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_route53globalresolverCmd.Name()}, args...))
		return p.Execute()
	}
	_route53globalresolverCmd.SetArgs(args)
	return _route53globalresolverCmd.Execute()
}
