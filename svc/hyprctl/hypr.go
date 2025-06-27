package hyprctl

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"slices"
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
	args := []string{"dispatch"}
	for _, p := range params {
		args = append(args, fmt.Sprint(p))
	}

	cmd := exec.Command("hyprctl", args...)
	stdout, stderr := cmd.CombinedOutput()
	out := string(stdout)
	return out, stderr
}
