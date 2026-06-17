package cmd

func Execute(args []string) error {
	if p := _clouddirectoryCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_clouddirectoryCmd.Name()}, args...))
		return p.Execute()
	}
	_clouddirectoryCmd.SetArgs(args)
	return _clouddirectoryCmd.Execute()
}
