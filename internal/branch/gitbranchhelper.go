package branch

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// TicketType represents the type of Rally ticket
type TicketType string

const (
	TicketTypeDefect TicketType = "Defect"
	TicketTypeStory  TicketType = "HierarchicalRequirement" // Rally's internal type for Stories/User Stories
)

// Ticket represents a resolved Rally ticket
type Ticket struct {
	ID   string
	Type TicketType
}

// Rally ticket formatted IDs are typically DE, US, or S followed by digits.
var ticketIDInBranch = regexp.MustCompile(`(?i)\b(?:DE|US|S)[0-9]+\b`)

func ResolveTicket(flagTicket, workDir string) (*Ticket, error) {
	var id string
	if s := strings.TrimSpace(flagTicket); s != "" {
		id = strings.ToUpper(s)
	} else {
		branch, err := currentGitBranch(workDir)
		if err != nil {
			return nil, fmt.Errorf("no ticket set (-d) and could not read git branch: %w", err)
		}

		id = ticketIDFromBranchName(branch)
		if id == "" {
			return nil, fmt.Errorf("no ticket set (-d) and branch %q has no ticket ID (expected DE<number>, US<number>, or S<number>, case-insensitive)", branch)
		}
	}

	ticketType := TicketTypeStory
	if strings.HasPrefix(id, "DE") {
		ticketType = TicketTypeDefect
	}

	return &Ticket{
		ID:   id,
		Type: ticketType,
	}, nil
}

func ticketIDFromBranchName(branch string) string {
	m := ticketIDInBranch.FindString(branch)
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
