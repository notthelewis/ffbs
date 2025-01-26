package packagemanager

type PackageManager interface {
	Install(packageName string, opts ...string) error
	String() string
}

func New(packagemanager string) PackageManager {
	switch packagemanager {
	case "apt":
		return Apt{}
	case "dnf":
		return Dnf{}
	case "pacman":
		return Pacman{}
	case "brew":
		return Brew{}
	default:
		panic("Unsupported package manager")
	}
}
