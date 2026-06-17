package cmd

func Execute(args []string) error {
	if p := _route53Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_route53Cmd.Name()}, args...))
		return p.Execute()
	}
	_route53Cmd.SetArgs(args)
	return _route53Cmd.Execute()
}
