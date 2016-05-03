// utils_git.go

package utils

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func GitIsValidBranchName(branch string) bool {
	cmd := exec.Command("git", "check-ref-format", fmt.Sprintf("%q", branch))
	cerr := cmd.Run()
	return cerr != nil
}

func GitIsCleanWorkingTree(path string) bool {
	// Check for differences in files
	// git diff-files --quiet --ignore-submodules
	cmd := exec.Command("git", "diff-files", "--quiet", "--ignore-submodules")
	cmd.Dir = path
	cerr := cmd.Run()
	if cerr != nil {
		if _, ok := cerr.(*exec.ExitError); ok {
			// exit error
			// there is diff in files
			return false
		}
	}

	// Check for differences in the index
	// git diff-index --cached --quiet HEAD --ignore-submodules
	cmd = exec.Command("git", "diff-index", "--cached", "--quiet", "HEAD", "--ignore-submodules")
	cmd.Dir = path
	cerr = cmd.Run()
	if cerr != nil {
		if _, ok := cerr.(*exec.ExitError); ok {
			// exit error
			// there is diff in index
			return false
		}
	}

	return true
}

func GitCreateBranchIfDoesNotExist(path string, branch string) error {
	// get list of all remotes
	rcmd := exec.Command("git", "remote")
	rcmd.Dir = path
	rout, rerr := rcmd.Output()
	if rerr != nil {
		log.Fatal(fmt.Sprintf("Unable to read git remotes in project %v.\n%s", path, rerr.Error()))
	}
	remotes := strings.Split(string(rout), "\n")
	for i, r := range remotes {
		remotes[i] = strings.Trim(r, " ")
	}

	// get list of all branches
	bcmd := exec.Command("git", "branch", "-a")
	bcmd.Dir = path
	bout, berr := bcmd.Output()
	if berr != nil {
		log.Fatal(fmt.Sprintf("Unable to read git branches in project %v.\n%s", path, berr.Error()))
	}
	branches := strings.Split(string(bout), "\n")
	for i, b := range branches {
		branches[i] = strings.Trim(b, " *")
	}

	// check branches
	for _, b := range branches {
		if b == branch {
			// local branch exists
			return nil
		}
		for _, r := range remotes {
			if b == fmt.Sprintf("remotes/%s/%s", r, branch) {
				// remote branch exists
				return nil
			}
		}
	}

	// create branch
	cmd := exec.Command("git", "branch", branch)
	cmd.Dir = path
	err := cmd.Run()
	return err
}

func GitCheckoutBranch(path string, branch string) error {
	cmd := exec.Command("git", "checkout", branch)
	cmd.Dir = path
	err := cmd.Run()
	return err
}

func GitGetBranch(path string) string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = path
	branch, err := cmd.Output()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error getting current branch in project %q.\n%s", path, err.Error()))
	}
	return string(branch)
}

func GitFetch(path string) {
	Print(fmt.Sprintf("On project %q. Fetching...", path))
	cmd := exec.Command("git", "fetch", "-p")
	cmd.Dir = path
	output, err := cmd.CombinedOutput()
	Print(string(output))
	if err != nil {
		log.Fatal(fmt.Sprintf("Error fetching project %q.\n%s", path, err.Error()))
	}
}

func GitGetUpstream(path string) string {
	cmd := exec.Command("git", "for-each-ref", "--format='%(upstream:short)'", "$(git symbolic-ref -q HEAD)")
	cmd.Dir = path
	upstream, err := cmd.Output()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error getting upstream in project %q.\n%s", path, err.Error()))
	}
	return string(upstream)
}

func GitPull(path string, branch string) error {
	cmd := exec.Command("git", "pull", "origin", branch)
	cmd.Dir = path
	output, err := cmd.CombinedOutput()
	Print(string(output))
	return err
}

func GitPush(path string, branch string) error {
	cmd := exec.Command("git", "push", "origin", branch)
	cmd.Dir = path
	output, err := cmd.CombinedOutput()
	Print(string(output))
	return err
}
