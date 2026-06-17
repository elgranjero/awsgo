package cmd

func Execute(args []string) error {
	if p := _sagemakergeospatialCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_sagemakergeospatialCmd.Name()}, args...))
		return p.Execute()
	}
	_sagemakergeospatialCmd.SetArgs(args)
	return _sagemakergeospatialCmd.Execute()
}
