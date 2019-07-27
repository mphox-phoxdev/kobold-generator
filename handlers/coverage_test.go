package handlers

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	rc := m.Run()
	envVarName := "HANDLERS_COVERAGE"

	// rc 0 means we've passed,
	// and CoverMode will be non empty if run with -cover
	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		coveragePercentage := .80
		parsedEnv, err := strconv.ParseFloat(os.Getenv(envVarName), 64)
		if err != nil {
			fmt.Printf("Reading environment variable %s failed, using default coverage of %v%%\n", envVarName, coveragePercentage*100.0)
		} else {
			coveragePercentage = parsedEnv
		}
		if c < coveragePercentage {
			fmt.Printf("Tests passed but %s failed at %v\n", envVarName, c)
			rc = -1
		}
	}
	os.Exit(rc)
}
