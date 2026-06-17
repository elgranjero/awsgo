package cmd

func Execute(args []string) error {
	if p := _lexmodelbuildingserviceCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_lexmodelbuildingserviceCmd.Name()}, args...))
		return p.Execute()
	}
	_lexmodelbuildingserviceCmd.SetArgs(args)
	return _lexmodelbuildingserviceCmd.Execute()
}
