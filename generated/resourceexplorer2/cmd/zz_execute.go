package cmd

func Execute(args []string) error {
	if p := _resourceexplorer2Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_resourceexplorer2Cmd.Name()}, args...))
		return p.Execute()
	}
	_resourceexplorer2Cmd.SetArgs(args)
	return _resourceexplorer2Cmd.Execute()
}
