package cmd

func Execute(args []string) error {
	if p := _resourcegroupsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_resourcegroupsCmd.Name()}, args...))
		return p.Execute()
	}
	_resourcegroupsCmd.SetArgs(args)
	return _resourcegroupsCmd.Execute()
}
