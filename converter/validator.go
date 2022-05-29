package converter

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/p2pquake/jmaxml-seis-parser-go/epsp"
	"github.com/p2pquake/jmaxml-seis-parser-go/jmaseis"
)

type ValidationWarning string
type ValidationError string

func (e ValidationWarning) Error() string {
	return fmt.Sprintf("EPSP validation warning: %s", string(e))
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("EPSP validation error: %s", string(e))
}

func ValidateJMAQuake(filename string, report *jmaseis.Report, jmaQuake *epsp.JMAQuake) []error {
	errors := []error{}

	if report.Control.Status != "通常" {
		errors = append(errors, ValidationWarning(fmt.Sprintf("Control.Status is training or exam (%s).", report.Control.Status)))
	}

	errors = append(errors, validateJMAQuakeIssueType(filename, jmaQuake)...)
	errors = append(errors, validateJMAQuakeTsunami(filename, jmaQuake)...)
	return errors
}

func ValidateJMATsunami(filename string, report *jmaseis.Report, jmaTsunami *epsp.JMATsunami) []error {
	errors := []error{}

	if report.Control.Status != "通常" {
		errors = append(errors, ValidationWarning(fmt.Sprintf("Control.Status is training or exam (%s).", report.Control.Status)))
	}

	if jmaTsunami.Cancelled && len(jmaTsunami.Areas) > 0 {
		errors = append(errors, ValidationError(fmt.Sprintf("%s is cancelled, but has areas", filename)))
	}
	if !jmaTsunami.Cancelled && len(jmaTsunami.Areas) <= 0 {
		errors = append(errors, ValidationError(fmt.Sprintf("%s has no area, but is not cancelled", filename)))
	}

	return errors
}

func validateJMAQuakeIssueType(filename string, jmaQuake *epsp.JMAQuake) []error {
	errors := []error{}

	if filename == "" {
		return errors
	}

	if strings.Contains(filename, "VXSE51") {
		if jmaQuake.Issue.Type != "ScalePrompt" {
			errors = append(errors, ValidationError(fmt.Sprintf("%s issue type (%s) is not valid", filename, jmaQuake.Issue.Type)))
		}
	} else if strings.Contains(filename, "VXSE52") {
		if jmaQuake.Issue.Type != "Destination" {
			errors = append(errors, ValidationError(fmt.Sprintf("%s issue type (%s) is not valid", filename, jmaQuake.Issue.Type)))
		}
	} else if strings.Contains(filename, "VXSE53") {
		if jmaQuake.Issue.Type != "DetailScale" && jmaQuake.Issue.Type != "Foreign" {
			errors = append(errors, ValidationError(fmt.Sprintf("%s issue type (%s) is not valid", filename, jmaQuake.Issue.Type)))
		}
	} else {
		if jmaQuake.Issue.Type != "Other" {
			errors = append(errors, ValidationError(fmt.Sprintf("%s issue type (%s) is not valid", filename, jmaQuake.Issue.Type)))
		}
	}

	return errors
}

func validateJMAQuakeTsunami(filename string, jmaQuake *epsp.JMAQuake) []error {
	errors := []error{}

	if jmaQuake.Issue.Type == "ScalePrompt" {
		if jmaQuake.Earthquake.DomesticTsunami != "Checking" {
			errors = append(errors, ValidationWarning(fmt.Sprintf("%s (%s) domesticTsunami is not valid (%s)", filename, jmaQuake.Issue.Type, jmaQuake.Earthquake.DomesticTsunami)))
		}
		if jmaQuake.Earthquake.ForeignTsunami != "Unknown" {
			errors = append(errors, ValidationWarning(fmt.Sprintf("%s (%s) foreignTsunami is not valid (%s)", filename, jmaQuake.Issue.Type, jmaQuake.Earthquake.ForeignTsunami)))
		}
	}

	if jmaQuake.Issue.Type == "Destination" || jmaQuake.Issue.Type == "DetailScale" {
		if !regexp.MustCompile("^(None|NonEffective|Watch|Warning)$").MatchString(jmaQuake.Earthquake.DomesticTsunami) {
			errors = append(errors, ValidationWarning(fmt.Sprintf("%s (%s) domesticTsunami is not valid (%s)", filename, jmaQuake.Issue.Type, jmaQuake.Earthquake.DomesticTsunami)))
		}
		if jmaQuake.Earthquake.ForeignTsunami != "Unknown" {
			errors = append(errors, ValidationWarning(fmt.Sprintf("%s (%s) foreignTsunami is not valid (%s)", filename, jmaQuake.Issue.Type, jmaQuake.Earthquake.ForeignTsunami)))
		}
	}

	if jmaQuake.Issue.Type == "Foreign" {
		if !regexp.MustCompile("^(Checking|None|NonEffective|Watch|Warning)$").MatchString(jmaQuake.Earthquake.DomesticTsunami) {
			errors = append(errors, ValidationWarning(fmt.Sprintf("%s (%s) domesticTsunami is not valid (%s)", filename, jmaQuake.Issue.Type, jmaQuake.Earthquake.DomesticTsunami)))
		}
		if !regexp.MustCompile("^(None|Checking|NonEffectiveNearby|WarningNearby|WarningPacific|WarningPacificWide|WarningIndian|WarningIndianWide|Potential|Unknown)$").MatchString(jmaQuake.Earthquake.ForeignTsunami) {
			errors = append(errors, ValidationWarning(fmt.Sprintf("%s (%s) foreignTsunami is not valid (%s)", filename, jmaQuake.Issue.Type, jmaQuake.Earthquake.ForeignTsunami)))
		}

	}

	return errors
}
