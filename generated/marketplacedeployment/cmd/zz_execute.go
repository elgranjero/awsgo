package cmd

func Execute(args []string) error {
	if p := _marketplacedeploymentCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_marketplacedeploymentCmd.Name()}, args...))
		return p.Execute()
	}
	_marketplacedeploymentCmd.SetArgs(args)
	return _marketplacedeploymentCmd.Execute()
}
