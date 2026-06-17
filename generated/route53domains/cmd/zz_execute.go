package cmd

func Execute(args []string) error {
	if p := _route53domainsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_route53domainsCmd.Name()}, args...))
		return p.Execute()
	}
	_route53domainsCmd.SetArgs(args)
	return _route53domainsCmd.Execute()
}
