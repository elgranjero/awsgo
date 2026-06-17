package cmd

func Execute(args []string) error {
	if p := _imagebuilderCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_imagebuilderCmd.Name()}, args...))
		return p.Execute()
	}
	_imagebuilderCmd.SetArgs(args)
	return _imagebuilderCmd.Execute()
}
