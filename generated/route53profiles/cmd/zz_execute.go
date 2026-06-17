package cmd

func Execute(args []string) error {
	if p := _route53profilesCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_route53profilesCmd.Name()}, args...))
		return p.Execute()
	}
	_route53profilesCmd.SetArgs(args)
	return _route53profilesCmd.Execute()
}
