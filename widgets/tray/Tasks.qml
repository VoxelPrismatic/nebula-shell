pragma ComponentBehavior: Bound

import QtQuick
import QtQuick.Controls
import QtQuick.Layouts
import Quickshell
import Quickshell.Wayland
import Quickshell.Services.SystemTray

import "root:/config"

Canvas {
	id: switcher

	property list<Toplevel> apps: ToplevelManager.toplevels.values.map(function (cur) {
		if (ToplevelManager.toplevels.values.find(e => e.appId == cur.appId) == cur) {
			return cur;
		}
	}).filter(e => e)

	width: Opts.workspace.cellSize * Opts.workspace.columns
	height: (Opts.workspace.cellSize * Math.max(1, Math.ceil((apps.length) / Opts.workspace.columns)))

	anchors.topMargin: 8
	anchors.top: parent.top
	anchors.horizontalCenter: parent.horizontalCenter

	GridLayout {
		id: grid
		columns: Opts.workspace.columns
		rows: Math.ceil(switcher.apps.length / Opts.workspace.columns)
		width: parent.width
		height: parent.height

		Repeater {
			model: switcher.apps

			delegate: Task {
				required property Toplevel modelData
				app: modelData
				isMax: false
			}
		}

		Repeater {
			model: 3 - switcher.apps.length
			Rectangle {
				width: Opts.workspace.cellSize
				height: this.width
				opacity: 0
			}
		}

		/*Task {
	id: plus
	idx: Hyprland.workspaces.values.length + 1
	label: "+"
	isMax: true
	visible: Hyprland.workspaces.values.reduce((acc, cur) => acc + (cur.lastIpcObject.windows == "0") ? 1 : 0, 0) == 0
	}*/

	}
}
