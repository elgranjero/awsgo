package cmd

func Execute(args []string) error {
	if p := _acmpcaCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_acmpcaCmd.Name()}, args...))
		return p.Execute()
	}
	_acmpcaCmd.SetArgs(args)
	return _acmpcaCmd.Execute()
}
