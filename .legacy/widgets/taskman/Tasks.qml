pragma ComponentBehavior: Bound

import QtQuick
import QtQuick.Controls
import QtQuick.Layouts
import Quickshell
import Quickshell.Wayland

import "root:/config"
import "root:/widgets/spaces"

Canvas {
	id: switcher

	property list<Toplevel> apps: ToplevelManager.toplevels.values.map(function (cur) {
		if (ToplevelManager.toplevels.values.find(e => e.appId == cur.appId) == cur) {
			return cur;
		}
	}).filter(e => e)

	width: Opts.workspace.cellSize * Opts.workspace.columns
	height: childrenRect.height

	anchors.topMargin: 8
	anchors.top: parent.top
	anchors.horizontalCenter: parent.horizontalCenter

	GridLayout {
		id: grid
		columns: Opts.workspace.columns
		anchors.fill: parent

		Repeater {
			model: switcher.apps

			delegate: Task {
				required property Toplevel modelData
				app: modelData
				isMax: false
			}
		}

		Repeater {
			model: 3 + (3 - switcher.apps.length % 3)
			Rectangle {
				width: Opts.workspace.cellSize - 4
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
