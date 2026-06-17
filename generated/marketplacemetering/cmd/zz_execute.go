package cmd

func Execute(args []string) error {
	if p := _marketplacemeteringCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_marketplacemeteringCmd.Name()}, args...))
		return p.Execute()
	}
	_marketplacemeteringCmd.SetArgs(args)
	return _marketplacemeteringCmd.Execute()
}
