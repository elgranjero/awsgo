package cmd

func Execute(args []string) error {
	if p := _ec2instanceconnectCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ec2instanceconnectCmd.Name()}, args...))
		return p.Execute()
	}
	_ec2instanceconnectCmd.SetArgs(args)
	return _ec2instanceconnectCmd.Execute()
}
