package cmd

func Execute(args []string) error {
	if p := _cleanroomsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cleanroomsCmd.Name()}, args...))
		return p.Execute()
	}
	_cleanroomsCmd.SetArgs(args)
	return _cleanroomsCmd.Execute()
}
