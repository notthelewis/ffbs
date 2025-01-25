package packagemanager

// TODO: Flesh out required methods

type Brew struct {
}

func (a Brew) String() string {
	return "brew"
}

func (a Brew) Install(packageName string, opts ...string) error {
	return nil
	// args := make([]string, 0, 2+len(opts))
	// args = append(args, "install")
	// args = append(args, packageName)
	// args = append(args, opts...)
	// args = append(args, "-y")
	//
	// return exec.Command("apt", args...).Run()
}

// TODO: Flesh out required methods

func (a Brew) IsInstalled(packageName string) bool {
	return false
}

func (a Brew) Remove(packageName string, opts ...string) error {

	return nil
}

func (a Brew) Update(packageName string, opts ...string) error {

	return nil
}

func (a Brew) Upgrade(packageName string, opts ...string) error {
	return nil
}

func (a Brew) Version() string {
	return ""
}
