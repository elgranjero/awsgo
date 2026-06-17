package cmd

func Execute(args []string) error {
	if p := _guarddutyCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_guarddutyCmd.Name()}, args...))
		return p.Execute()
	}
	_guarddutyCmd.SetArgs(args)
	return _guarddutyCmd.Execute()
}
