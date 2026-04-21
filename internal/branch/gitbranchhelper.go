package branch

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// Rally defect formatted IDs are typically DE followed by digits.
var defectIDInBranch = regexp.MustCompile(`(?i)\bDE[0-9]+\b`)

func ResolveDefectID(flagDefect, workDir string) (string, error) {
	if s := strings.TrimSpace(flagDefect); s != "" {
		return s, nil
	}
	branch, err := currentGitBranch(workDir)
	if err != nil {
		return "", fmt.Errorf("no defect set (-d) and could not read git branch: %w", err)
	}
	id := defectIDFromBranchName(branch)
	if id == "" {
		return "", fmt.Errorf("no defect set (-d) and branch %q has no defect ID (expected DE<number>, case-insensitive)", branch)
	}
	return id, nil
}

func defectIDFromBranchName(branch string) string {
	m := defectIDInBranch.FindString(branch)
	if m == "" {
		return ""
	}
	return strings.ToUpper(m)
}

func currentGitBranch(workDir string) (string, error) {
	dir := workDir
	if dir == "" {
		var err error
		dir, err = os.Getwd()
		if err != nil {
			return "", err
		}
	}

	out, err := exec.Command("git", "-C", dir, "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
