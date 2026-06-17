package cmd

func Execute(args []string) error {
	if p := _rolesanywhereCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_rolesanywhereCmd.Name()}, args...))
		return p.Execute()
	}
	_rolesanywhereCmd.SetArgs(args)
	return _rolesanywhereCmd.Execute()
}
