package cmd

func Execute(args []string) error {
	if p := _resourcegroupstaggingapiCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_resourcegroupstaggingapiCmd.Name()}, args...))
		return p.Execute()
	}
	_resourcegroupstaggingapiCmd.SetArgs(args)
	return _resourcegroupstaggingapiCmd.Execute()
}
