package cmd

func Execute(args []string) error {
	if p := _cleanroomsmlCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cleanroomsmlCmd.Name()}, args...))
		return p.Execute()
	}
	_cleanroomsmlCmd.SetArgs(args)
	return _cleanroomsmlCmd.Execute()
}
