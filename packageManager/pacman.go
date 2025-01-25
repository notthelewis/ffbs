package packagemanager

// TODO: Flesh out required methods

type Pacman struct {
}

func (a Pacman) String() string {
	return "pacman"
}

func (a Pacman) Install(packageName string, opts ...string) error {
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

func (a Pacman) IsInstalled(packageName string) bool {
	return false
}

func (a Pacman) Remove(packageName string, opts ...string) error {

	return nil
}

func (a Pacman) Update(packageName string, opts ...string) error {

	return nil
}

func (a Pacman) Upgrade(packageName string, opts ...string) error {
	return nil
}

func (a Pacman) Version() string {
	return ""
}
