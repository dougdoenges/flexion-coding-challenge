package main

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientEndToEnd(t *testing.T) {
	worksheetTemplate := "--worksheet=../../test/data/%s"
	responseTemplate := "--responses=../../test/data/%s"

	cmd := exec.Command("go", "run", "./main.go", fmt.Sprintf(worksheetTemplate, "validWs.csv"))
	_, err := cmd.CombinedOutput()
	assert.Error(t, err)

	cmd = exec.Command("go", "run", "./main.go", fmt.Sprintf(responseTemplate, "validResponses.csv"))
	_, err = cmd.CombinedOutput()
	assert.Error(t, err)

	cmd = exec.Command("go", "run", "./main.go", fmt.Sprintf(worksheetTemplate, "validWs.csv"),
		fmt.Sprintf(responseTemplate, "validResponses.csv"))
	_, err = cmd.CombinedOutput()
	assert.Error(t, err)

	cmd = exec.Command("go", "run", "./main.go", fmt.Sprintf(worksheetTemplate, "validWs.csv"),
		fmt.Sprintf(responseTemplate, "validResponses.csv"),
		"--output=../../test/data/validResults.csv")
	_, err = cmd.CombinedOutput()
	assert.Nil(t, err)

	cmd = exec.Command("go", "run", "./main.go", fmt.Sprintf(worksheetTemplate, "invalidWs.csv"),
		fmt.Sprintf(responseTemplate, "validResponses.csv"),
		"--output=../../test/data/validResults.csv")
	_, err = cmd.CombinedOutput()
	assert.Error(t, err)

	cmd = exec.Command("go", "run", "./main.go", fmt.Sprintf(worksheetTemplate, "validWs.csv"),
		fmt.Sprintf(responseTemplate, "invalidResponses.csv"),
		"--output=../../test/data/validResults.csv")
	_, err = cmd.CombinedOutput()
	assert.Error(t, err)

	cmd = exec.Command("go", "run", "./main.go", fmt.Sprintf(worksheetTemplate, "invalidWs.csv"),
		fmt.Sprintf(responseTemplate, "invalidResponses.csv"),
		"--output=../../test/data/validResults.csv")
	_, err = cmd.CombinedOutput()
	assert.Error(t, err)

	cmd = exec.Command("go", "run", "./main.go", fmt.Sprintf(worksheetTemplate, "validWs.csv"),
		fmt.Sprintf(responseTemplate, "doesnotexist.csv"),
		"--output=../../test/data/validResults.csv")
	_, err = cmd.CombinedOutput()
	assert.Error(t, err)
}
