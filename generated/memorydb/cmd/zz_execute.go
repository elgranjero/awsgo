package cmd

func Execute(args []string) error {
	if p := _memorydbCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_memorydbCmd.Name()}, args...))
		return p.Execute()
	}
	_memorydbCmd.SetArgs(args)
	return _memorydbCmd.Execute()
}
