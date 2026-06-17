package cmd

func Execute(args []string) error {
	if p := _partnercentralsellingCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_partnercentralsellingCmd.Name()}, args...))
		return p.Execute()
	}
	_partnercentralsellingCmd.SetArgs(args)
	return _partnercentralsellingCmd.Execute()
}
