package main

// TODO: Handle package name differences between package managers
// TODO: Fish shell setup

import (
	packagemanager "ffbs/packageManager"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"

	"github.com/atotto/clipboard"
)

// This is global because this is the easiest way to pass the package between subroutines when
// you're lazy, don't give a fuck and don't intend to release the package to anyone other than
// your own stupid self
var packageManagerName string

func main() {
	wg := sync.WaitGroup{}
	wg.Add(4)

	installPackages()

	home := os.Getenv("HOME")
	sshKeygen(home)
	openGithub()

	go func() {
		if err := os.MkdirAll(home, 0776); err != nil {
			fmt.Println("Unable to create conf dir structure for nvim", err.Error())
		}
		err := CopyDirectory("nvim-lua", home+"/.config/nvim")
		if err != nil {
			fmt.Println("unable to copy nvim conf", err.Error())
		}
		wg.Done()
	}()

	go func() {
		err := Copy("./floating-conf/tmux.conf", home+"/.tmux.conf")
		if err != nil {
			fmt.Println("unable to copy tmux conf", err.Error())
		}

		wg.Done()
	}()

	go func() {
		makeAlacrittyConfig(home)
		wg.Done()
	}()

	go func() {
		goFishing()
		wg.Done()
	}()

	wg.Wait()
}

func ExitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// NOTE: sshKeygen runs ssh-keygen interactively, then copy-pastes the ssh-key from a *hard-coded*
// path, so if future me decides to store the key somewhere else- this is why it won't copy
// you dumbfuck
func sshKeygen(home string) {
	cmd := exec.Command("ssh-keygen")

	// Allows go to take over cmd in & out
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	ExitOnError(cmd.Start())
	ExitOnError(cmd.Wait())

	publicKey, err := os.ReadFile(home + "/.ssh/id_ed25519.pub")
	if err != nil {
		fmt.Println("Error reading public key", err)
		return
	}

	// NOTE: I'm lazy, which is why I used a library to copy the public key to the clipboard
	if err := clipboard.WriteAll(string(publicKey)); err != nil {
		fmt.Println("Error copying public key to clipboard", err)
		return
	}

	fmt.Println("Public key copied to clipboard... Navigation to github.com/login required. Trying automatically")
}

func installPackages() {
	packageManager := getPackageManager()
	fmt.Printf("Using package manager '%s'\n", packageManager)

	/* Required terminal packages */
	if err := packageManager.Install("nvim"); err != nil {
		fmt.Println("Error installing nvim", err.Error())
		return
	}
	if err := packageManager.Install("tmux"); err != nil {
		fmt.Println("Error installing tmux", err.Error())
		return
	}
	if err := packageManager.Install("fish"); err != nil {
		fmt.Println("Error installing fish", err.Error())
		return
	}

	/* Required GUI packages */
	if err := packageManager.Install("alacritty"); err != nil {
		fmt.Println("Error installing alacritty", err.Error())
		return
	}
	if err := packageManager.Install("brave-browser"); err != nil {
		fmt.Println("Error installing brave-browser", err.Error())
		return
	}
	if err := packageManager.Install("discord"); err != nil {
		fmt.Println("Error installing discord", err.Error())
		return
	}

	/* Linux specific packages */
	if packageManagerName != "brew" {
		installLinuxSpecificDeps(packageManager)
	}
}

func installLinuxSpecificDeps(pm packagemanager.PackageManager) {
	// for clipboard
	if err := pm.Install("xclip"); err != nil {
		fmt.Println("Error installing xclip", err.Error())
		return
	}

	if err := pm.Install("make"); err != nil {
		fmt.Println("Error installing make", err.Error())
		return
	}

}

func getPackageManager() packagemanager.PackageManager {
	switch runtime.GOOS {
	case "windows":
		panic("windows not supported")
	case "linux":
		result, err := exec.Command("lsb_release", "-d").Output()
		if err != nil {
			panic(err)
		}

		strings.TrimSpace(string(result))
		osName := strings.Split(string(result), "\t")[1]

		fmt.Println("os = ", string(result))

		if strings.Contains(osName, "Ubuntu") || strings.Contains(osName, "Debian") {
			packageManagerName = "apt"
		} else if strings.Contains(osName, "Fedora") {
			packageManagerName = "dnf"
		} else if strings.Contains(osName, "Arch") {
			packageManagerName = "pacman"
		} else {
			panic("Unsupported OS " + osName)
		}

	case "darwin":
		// Check if brew is installed
		_, err := exec.Command("brew", "--version").Output()
		if err != nil {
			// Install brew
			brewInstallErr := exec.Command("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)").Run()
			if brewInstallErr != nil {
				panic(brewInstallErr)
			}
		}
		packageManagerName = "brew"
	}

	return packagemanager.New(packageManagerName)
}

func openGithub() {
	// At this point, we're either on Mac or some supported Linux distro
	// so it's safe to assume that we can use 'open' or 'xdg-open'
	var cmd *exec.Cmd
	if packageManagerName == "brew" {
		cmd = exec.Command("open", "https://github.com/login")
	} else {
		cmd = exec.Command("xdg-open", "https://github.com/login")
	}

	ExitOnError(cmd.Run())
}

// NOTE: This is dynamic because the conf is tiny and I don't know for sure what the path
// to fish will be on the target machine
func makeAlacrittyConfig(home string) {
	err := os.MkdirAll(home+"/.config/alacritty", 0776)
	if err != nil {
		fmt.Println("Unable to create conf dir structure for alacritty", err.Error())
	}

	fishPath, err := exec.LookPath("fish")
	if err != nil {
		fmt.Println("Unable to find fish shell... aborting fish conf bootstrap", err.Error())
		return
	}

	sb := strings.Builder{}
	sb.WriteString("shell = \"")
	sb.WriteString(fishPath)
	sb.WriteString("\"\n\n")
	sb.WriteString("[window]\n")
	sb.WriteString("startup_mode = \"fullscreen\"\n")
	sb.WriteString("padding = { x = 10, y = 10 }\n")

	if err := os.MkdirAll(home+"/.config/alacritty", 0776); err != nil {
		fmt.Println("Unable to create conf dir structure for alacritty", err.Error())
		return
	}

	if err := os.WriteFile(home+"/.config/alacritty/alacritty.toml", []byte(sb.String()), 0776); err != nil {
		fmt.Println("Unable to write alacritty conf", err.Error())
		return
	}
}

// Fish shell plugins setup
func goFishing() {
	fishPath, err := exec.LookPath("fish")
	if err != nil {
		fmt.Println("Unable to find fish shell... aborting fish conf bootstrap", err.Error())
		return
	}

	// Write fisher install script
	toWrite := "curl -sL https://raw.githubusercontent.com/jorgebucaran/fisher/main/functions/fisher.fish | source && fisher install jorgebucaran/fisher"
	err = os.WriteFile("/tmp/fisherInstall.sh", []byte(toWrite), 0776)
	if err != nil {
		fmt.Println("Error writing fisher install script", err.Error())
		return
	}

	// Curl fisher, install fisher, install fisher plugins
	cmd := exec.Command(fishPath, "/tmp/fisherInstall.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error installing fisher", err.Error())
		return
	}

	err = os.Remove("/tmp/fisherInstall.sh")
	if err != nil {
		fmt.Println("Error removing fisher install script @ /tmp/fisherInstall.sh", err.Error())
		return
	}

	cmd = exec.Command(fishPath, "-c", "fisher", "install", "jorgebucaran/nvm.fish")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error installing nvm.fish", err.Error())
		return
	}

	cmd = exec.Command(fishPath, "-c", "fish_vi_key_bindings")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error setting fish_vi_key_bindings", err.Error())
		return
	}

	// Auto function to load latest node version and swap caps and escape

	toWrite = `if status is-interactive
    # Commands to run in interactive sessions can go here
    setxkbmap -option caps:swapescape
    nvm use latest
end`

	cmd = exec.Command(fishPath, "-c", "echo '"+toWrite+"' > ~/.config/fish/config.fish")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error writing to fish config", err.Error())
		return
	}
}
