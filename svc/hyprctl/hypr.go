package hyprctl

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"slices"
	"strings"
)

func Call[T any](args ...string) (*T, error) {
	args = slices.Concat([]string{"-j"}, args)
	cmd := exec.Command("hyprctl", args...)
	stdout, stderr := cmd.CombinedOutput()
	if stderr != nil {
		return nil, stderr
	}

	ret := new(T)
	err := json.Unmarshal(stdout, ret)
	return ret, err
}

func Splash() (string, error) {
	cmd := exec.Command("hyprctl", "splash")
	stdout, stderr := cmd.CombinedOutput()
	out := string(stdout)
	return out, stderr
}

func Dispatch(params ...any) (string, error) {
	if len(params) == 0 {
		return "empty", nil
	}

	args := make([]string, len(params)+1)
	args[0] = "dispatch"
	for i, p := range params {
		args[i+1] = fmt.Sprint(p)
	}

	cmd := exec.Command("hyprctl", args...)
	stdout, stderr := cmd.CombinedOutput()
	out := string(stdout)
	return out, stderr
}

func BatchDispatch(params ...[]any) (string, error) {
	batches := []string{}
	for _, b := range params {
		if len(b) == 0 {
			continue
		}
		batch := make([]string, len(b)+1)
		batch[0] = "dispatch"
		for j, p := range b {
			batch[j+1] = fmt.Sprint(p)
		}
		batches = append(batches, strings.Join(batch, " "))
	}
	if len(batches) == 0 {
		return "", nil
	}

	args := []string{"--batch", strings.Join(batches, "; ")}

	cmd := exec.Command("hyprctl", args...)
	stdout, stderr := cmd.CombinedOutput()
	out := string(stdout)
	return out, stderr
}
