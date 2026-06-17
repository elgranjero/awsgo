package cmd

func Execute(args []string) error {
	if p := _ec2Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ec2Cmd.Name()}, args...))
		return p.Execute()
	}
	_ec2Cmd.SetArgs(args)
	return _ec2Cmd.Execute()
}
