package cmd

func Execute(args []string) error {
	if p := _connectcampaignsv2Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_connectcampaignsv2Cmd.Name()}, args...))
		return p.Execute()
	}
	_connectcampaignsv2Cmd.SetArgs(args)
	return _connectcampaignsv2Cmd.Execute()
}
