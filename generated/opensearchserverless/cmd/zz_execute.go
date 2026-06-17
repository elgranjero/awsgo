package cmd

func Execute(args []string) error {
	if p := _opensearchserverlessCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_opensearchserverlessCmd.Name()}, args...))
		return p.Execute()
	}
	_opensearchserverlessCmd.SetArgs(args)
	return _opensearchserverlessCmd.Execute()
}
