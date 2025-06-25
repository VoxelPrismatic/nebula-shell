package hypripc

import (
	"bufio"
	"fmt"
	"nebula-shell/svc/hyprctl"
	"net"
	"os"
	"strings"
)

type Environment map[string]string
type GenericEventListener interface {
	update(event, value string)
}

type IpcListener struct {
	// Emitted on workspace change. Is emitted ONLY when a user requests a workspace change,
	// and is not emitted on mouse movements (see the focused monitor event)
	EvtWorkspace EventListener[*IpcWorkspace, hyprctl.HyprWorkspace]

	// Emitted on the active monitor being changed.
	EvtFocusedMonitor EventListener[*IpcFocusedMonitor, hyprctl.HyprMonitor]

	// Emitted on the active window being changed.
	EvtActiveWindow EventListener[*IpcActiveWindow, hyprctl.HyprWindow]

	// Emitted when a fullscreen status of a window changes. True/False = Enter/Leave fullscreen
	EvtFullscreen EventListener[*IpcFullscreen, bool]

	// Emitted when a monitor is removed (disconnected)
	EvtMonitorRemoved EventListener[*IpcMonitorRemoved, hyprctl.HyprMonitor]

	// Emitted when a monitor is added (connected)
	EvtMonitorAdded EventListener[*IpcMonitorAdded, hyprctl.HyprMonitor]

	// Emitted when a workspace is created
	EvtCreateWorkspace EventListener[*IpcCreateWorkspace, hyprctl.HyprWorkspace]

	// Emitted when a workspace is destroyed
	EvtDestroyWorkspace EventListener[*IpcDestroyWorkspace, hyprctl.HyprWorkspace]

	// Emitted when a workspace is moved to a different monitor
	EvtMoveWorkspace EventListener[*IpcMoveWorkspace, hyprctl.HyprMonitor]

	// Emitted when a workspace is renamed
	EvtRenameWorkspace EventListener[*IpcRenameWorkspace, hyprctl.HyprWorkspace]

	// Emitted when the special workspace opened in a monitor changes
	// (closing results in a nil Workspace reference)
	EvtActiveSpecial EventListener[*IpcActiveSpecial, hyprctl.HyprMonitor]

	// Emitted on a layout change of the active keyboard
	EvtActiveLayout EventListener[*IpcActiveLayout, hyprctl.HyprDevKeyboard]

	// Emitted when a window is opened
	EvtOpenWindow EventListener[*IpcOpenWindow, hyprctl.HyprWindow]

	// Emitted when a window is closed
	EvtCloseWindow EventListener[*IpcCloseWindow, hyprctl.HyprWindow]

	// Emitted when a window is moved to a workspace
	EvtMoveWindow EventListener[*IpcMoveWindow, hyprctl.HyprWindow]

	// Emitted when a layerSurface is mapped
	EvtOpenLayer EventListener[*IpcOpenLayer, []hyprctl.HyprLayerTarget]

	// Emitted when a layerSurface is mapped
	EvtCloseLayer EventListener[*IpcCloseLayer, []hyprctl.HyprLayerTarget]

	// Emitted when a keybind submap changes. Empty means default.
	EvtSubmap EventListener[*IpcSubmap, string]

	// Emitted when a window changes it's floating mode
	FloatChanged EventListener[*IpcFloatChanged, hyprctl.HyprWindow]

	// Emitted when a window requests an urgent state
	EvtUrgent EventListener[*IpcUrgent, hyprctl.HyprWindow]

	// Emitted when a screencopy state of a client changes.
	// Keep in mind there might be multiple separate clients.
	EvtScreencast EventListener[*IpcScreencast, CastingOwner]

	// Emitted when a window title changes
	EvtWindowTitle EventListener[*IpcWindowTitle, hyprctl.HyprWindow]

	// Emitted when togglegroup command is used
	// If a group is destroyed, all windows are returned, and the Included flag is set to False
	EvtToggleGroup EventListener[*IpcToggleGroup, hyprctl.HyprWindow]

	// Emitted when the window is merged into a group. Returns the address of a merged window
	EvtMoveIntoGroup EventListener[*IpcMoveIntoGroup, hyprctl.HyprWindow]

	// Emitted when the window is removed from a group. Returns the address of a removed window
	EvtMoveOutOfGroup EventListener[*IpcMoveOutOfGroup, hyprctl.HyprWindow]

	// Emitted when ignoregrouplock is toggled
	EvtIgnoreGroupLock EventListener[*IpcIgnoreGroupLock, bool]

	// Emitted when lockgroups is toggled
	EvtLockGroups EventListener[*IpcLockGroups, bool]

	// Emitted when the config is done reloading. This has no target value.
	EvtConfigReloaded EventListener[*IpcConfigReloaded, bool]

	// Emitted when a window is pinned or unpinned
	EvtPinWindow EventListener[*IpcPinWindow, hyprctl.HyprWindow]

	// Emitted when an external taskbar-like app requests a window to be minimized
	EvtMinimized EventListener[*IpcMinimized, hyprctl.HyprWindow]

	// Emitted when an app requests to ring the system bell via xdg-system-bell-v1.
	// Window may be nil pointer.
	EvtBell EventListener[*IpcBell, hyprctl.HyprWindow]
}

func Env() Environment {
	ret := Environment{}
	for _, line := range os.Environ() {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 1 {
			ret[parts[0]] = ""
		} else {
			ret[parts[0]] = parts[1]
		}
	}
	return ret
}

func (env Environment) Get(name string) (string, error) {
	val, ok := env[name]
	if !ok || val == "" {
		return "", fmt.Errorf("environment variable $%s not set", name)
	}
	return val, nil
}

func fire[T any](obj *T, listeners []func(T)) {
	for _, listener := range listeners {
		listener(*obj)
	}
	*obj = *new(T)
}

func mustSplitN(value string, count int) []string {
	ret := strings.SplitN(value, ",", count)
	if len(ret) != count {
		panic("ipc: changed spec")
	}
	return ret
}

func (ipc *IpcListener) Connect() error {
	env := Env()
	xdgRtDir, err := env.Get("XDG_RUNTIME_DIR")
	if err != nil {
		return err
	}

	hyprInstSig, err := env.Get("HYPRLAND_INSTANCE_SIGNATURE")
	if err != nil {
		return err
	}

	socket := fmt.Sprintf("%s/hypr/%s/.socket2.sock", xdgRtDir, hyprInstSig)
	conn, err := net.Dial("unix", socket)
	if err != nil {
		return fmt.Errorf("socket: %+v", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	handlers := map[string]GenericEventListener{
		"workspace":          &ipc.EvtWorkspace,
		"workspacev2":        &ipc.EvtWorkspace,
		"focusedmon":         &ipc.EvtFocusedMonitor,
		"focusedmonv2":       &ipc.EvtFocusedMonitor,
		"activewindow":       &ipc.EvtActiveWindow,
		"activewindowv2":     &ipc.EvtActiveWindow,
		"fullscreen":         &ipc.EvtFullscreen,
		"monitorremoved":     &ipc.EvtMonitorRemoved,
		"monitorremovedv2":   &ipc.EvtMonitorRemoved,
		"monitoradded":       &ipc.EvtMonitorAdded,
		"monitoraddedv2":     &ipc.EvtMonitorAdded,
		"createworkspace":    &ipc.EvtCreateWorkspace,
		"createworkspacev2":  &ipc.EvtCreateWorkspace,
		"destroyworkspace":   &ipc.EvtDestroyWorkspace,
		"destroyworkspacev2": &ipc.EvtDestroyWorkspace,
		"moveworkspace":      &ipc.EvtMoveWorkspace,
		"moveworkspacev2":    &ipc.EvtMoveWorkspace,
		"renameworkspace":    &ipc.EvtRenameWorkspace,
		"activespecial":      &ipc.EvtActiveSpecial,
		"activespecial2":     &ipc.EvtActiveSpecial,
		"activelayout":       &ipc.EvtActiveLayout,
		"openwindow":         &ipc.EvtOpenWindow,
		"closewindow":        &ipc.EvtCloseWindow,
		"movewindow":         &ipc.EvtMoveWindow,
		"movewindowv2":       &ipc.EvtMoveWindow,
		"openlayer":          &ipc.EvtOpenLayer,
		"closelayer":         &ipc.EvtCloseLayer,
		"submap":             &ipc.EvtSubmap,
		"changefloatingmode": &ipc.FloatChanged,
		"urgent":             &ipc.EvtUrgent,
		"screencast":         &ipc.EvtScreencast,
		"windowtitle":        &ipc.EvtWindowTitle,
		"windowtitlev2":      &ipc.EvtWindowTitle,
		"togglegroup":        &ipc.EvtToggleGroup,
		"moveintogroup":      &ipc.EvtMoveIntoGroup,
		"moveoutofgroup":     &ipc.EvtMoveOutOfGroup,
		"ignoregrouplock":    &ipc.EvtIgnoreGroupLock,
		"lockgroups":         &ipc.EvtLockGroups,
		"configreloaded":     &ipc.EvtConfigReloaded,
		"pin":                &ipc.EvtPinWindow,
		"minimized":          &ipc.EvtMinimized,
		"bell":               &ipc.EvtBell,
	}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		parts := strings.SplitN(line[:len(line)-1], ">>", 2)
		if len(parts) != 2 {
			return fmt.Errorf("spec changed; line=%s", line)
		}

		event, value := parts[0], parts[1]
		handler, ok := handlers[event]
		if !ok {
			fmt.Printf("not implemented: %s\n", event)
		} else {
			handler.update(event, value)
		}
	}
}
