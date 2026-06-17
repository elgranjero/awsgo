package cmd

func Execute(args []string) error {
	if p := _secretsmanagerCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_secretsmanagerCmd.Name()}, args...))
		return p.Execute()
	}
	_secretsmanagerCmd.SetArgs(args)
	return _secretsmanagerCmd.Execute()
}
