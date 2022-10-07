package cmd

var root = NewRootCommand()

func Run() error {
	return root.Execute()
}
