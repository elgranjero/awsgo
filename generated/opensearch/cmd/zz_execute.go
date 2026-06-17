package cmd

func Execute(args []string) error {
	if p := _opensearchCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_opensearchCmd.Name()}, args...))
		return p.Execute()
	}
	_opensearchCmd.SetArgs(args)
	return _opensearchCmd.Execute()
}
