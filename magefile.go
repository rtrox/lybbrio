//go:build mage
// +build mage

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
)

const (
	repo            = "github.com/rtrox/lybbr.io"
	name            = "lybbrio"
	buildEntrypoint = "./cmd/lybbrio/main.go"
)

var (
	VersionNumber = "dev"
	Revision      = "unknown"
	LdFlags       = ""
	BuildTime     = ""
	RootPath      = ""

	// Aliases are mage aliases of targets
	Aliases = map[string]interface{}{
		"build":          Build.Build,
		"clean":          Build.Clean,
		"generate":       Build.Generate,
		"test":           Check.Test,
		"lint":           Check.Lint,
		"new-ent":        Build.NewEnt,
		"check":          Check.All,
		"check:lint":     Check.Lint,
		"check:lint-fix": Check.LintFix,
		"check:test":     Check.Test,
		"check:vet":      Check.Vet,
		"check:fmt":      Check.Fmt,
		"check:mod-tidy": Check.ModTidy,
		"check:all":      Check.All,
	}
	Default = Build.Build
)

func runCmdWithOutput(name string, arg ...string) (output []byte, err error) {
	cmd := exec.Command(name, arg...)

	cmd.Env = os.Environ()
	cmd.Dir = RootPath

	output, err = cmd.Output()
	if err != nil {
		if ee, is := err.(*exec.ExitError); is {
			return nil, fmt.Errorf("error running command: %s, %s", string(ee.Stderr), err)
		}
		return nil, fmt.Errorf("error running command: %s", err)
	}

	return output, nil
}

func runAndStreamOutput(cmd string, args ...string) {
	c := exec.Command(cmd, args...)

	c.Env = os.Environ()
	c.Dir = RootPath

	fmt.Printf("%s\n\n", c.String())

	stdout, _ := c.StdoutPipe()
	stderr, _ := c.StderrPipe()
	err := c.Start()
	if err != nil {
		fmt.Printf("Could not start: %s\n", err)
		os.Exit(1)
	}

	go func() {
		reader := bufio.NewReader(stdout)
		line, err := reader.ReadString('\n')
		for err == nil {
			fmt.Print(line)
			line, err = reader.ReadString('\n')
		}
	}()

	reader2 := bufio.NewReader(stderr)
	line, err := reader2.ReadString('\n')
	for err == nil {
		fmt.Print(line)
		line, err = reader2.ReadString('\n')
	}

	if err := c.Wait(); err != nil {
		os.Exit(1)
	}
}

func setVersion() {
	version, err := runCmdWithOutput("git", "describe", "--tags", "--always", "--abbrev=10")
	if err != nil {
		fmt.Printf("Error getting version: %s\n", err)
		os.Exit(1)
	}
	VersionNumber = strings.Trim(string(version), "\n")
	VersionNumber = strings.Replace(VersionNumber, "-", "+", 1)
	VersionNumber = strings.Replace(VersionNumber, "-g", "-", 1)

	revision := os.Getenv("GITHUB_SHA")
	if revision == "" {
		rev2, err := runCmdWithOutput("git", "rev-parse", "--short=10", "HEAD")
		revision = strings.Trim(string(rev2), "\n")
		if err != nil {
			fmt.Printf("Error getting revision: %s\n", err)
			os.Exit(1)
		}
	}
	Revision = strings.Trim(string(revision), "\n")
}

func setRootPath() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting root path: %s\n", err)
		os.Exit(1)
	}
	RootPath = pwd
}

func initVars() {
	setVersion()
	setRootPath()
	if BuildTime == "" {
		BuildTime = time.Now().Format(time.RFC3339)
	}
	LdFlags = fmt.Sprintf("-s -w -X main.version=%s -X main.revision=%s -X main.buildTime=%s", VersionNumber, Revision, BuildTime)
}

func Run() {
	mg.Deps(initVars)
	fmt.Printf("Reminder: Mage does not yet support targets with variadic arguments, so you must use a env vars with mage run.\n")
	fmt.Printf("Running %s version %s, revision %s\n", name, VersionNumber, Revision)
	runAndStreamOutput("go", "run", "-ldflags", LdFlags, buildEntrypoint)
}

type Build mg.Namespace

func (Build) Build() {
	mg.Deps(initVars)
	fmt.Printf("Building %s version %s, revision %s\n", name, VersionNumber, Revision)
	runAndStreamOutput("go", "build", "-ldflags", LdFlags, "-o", name, buildEntrypoint)
}

func (Build) Clean() {
	mg.Deps(initVars)
	if err := exec.Command("go", "clean", "./...").Run(); err != nil {
		fmt.Printf("Error cleaning: %s\n", err)
		os.Exit(1)
	}
}

func (Build) Generate() {
	mg.Deps(initVars)
	fmt.Printf("Generating %s version %s, revision %s\n", name, VersionNumber, Revision)
	runAndStreamOutput("go", "generate", ".")
}

func (Build) NewEnt(entName string) {
	mg.Deps(initVars)
	fmt.Printf("Creating new ent %s version %s, revision %s\n", name, VersionNumber, Revision)
	runAndStreamOutput("go", "run", "-mod=mod", "entgo.io/ent/cmd/ent", "new", entName, "--target", "internal/ent/schema")
}

type Check mg.Namespace

func (Check) Lint() {
	mg.Deps(initVars)
	fmt.Printf("Linting %s version %s, revision %s\n", name, VersionNumber, Revision)
	runAndStreamOutput("golangci-lint", "run", "--config", ".github/lint/golangci.yaml")
}

func (Check) LintFix() {
	mg.Deps(initVars)
	fmt.Printf("Linting %s version %s, revision %s\n", name, VersionNumber, Revision)
	runAndStreamOutput("golangci-lint", "run", "--fix")
}

func (Check) Test() {
	mg.Deps(initVars)
	fmt.Printf("Testing %s version %s, revision %s\n", name, VersionNumber, Revision)
	runAndStreamOutput("go", "test", "-cover", "./...")
}

func (Check) Vet() {
	mg.Deps(initVars)
	fmt.Printf("Vetting %s version %s, revision %s\n", name, VersionNumber, Revision)
	runAndStreamOutput("go", "vet", "./...")
}

func (Check) Fmt() {
	mg.Deps(initVars)
	fmt.Printf("Formatting %s version %s, revision %s\n", name, VersionNumber, Revision)
	runAndStreamOutput("go", "fmt", "./...")
}

func (Check) ModTidy() {
	mg.Deps(initVars)
	fmt.Printf("Tidying %s version %s, revision %s\n", name, VersionNumber, Revision)
	runAndStreamOutput("go", "mod", "tidy")
}

func (Check) All() {
	mg.Deps(initVars)
	mg.Deps(Check.Lint)
	mg.Deps(Check.Test)
	mg.Deps(Check.Vet)
	mg.Deps(Check.Fmt)
	mg.Deps(Check.ModTidy)
}
