package cmd

func Execute(args []string) error {
	if p := _auditmanagerCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_auditmanagerCmd.Name()}, args...))
		return p.Execute()
	}
	_auditmanagerCmd.SetArgs(args)
	return _auditmanagerCmd.Execute()
}
