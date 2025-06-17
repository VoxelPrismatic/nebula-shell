pragma ComponentBehavior: Bound

import QtQuick
import QtQuick.Controls
import QtQuick.Layouts
import Quickshell
import Quickshell.Hyprland

import "root:/config"
import "root:/svc"

Canvas {
	id: switcher

	width: Opts.workspace.cellSize * Opts.workspace.columns
	height: Opts.workspace.cellSize * Math.ceil((Hyprland.workspaces.values.length + (plus.visible ? 1 : 0)) / Opts.workspace.columns)

	anchors.topMargin: 8
	anchors.top: parent.top
	anchors.horizontalCenter: parent.horizontalCenter

	GridLayout {
		id: grid
		columns: Opts.workspace.columns
		rows: Math.ceil(Hypr.workspaces.values.length / Opts.workspace.columns)
		width: parent.width
		height: parent.height

		Repeater {
			model: Hypr.workspaces.values.length

			Workspace {
				required property int modelData
				idx: modelData
				label: Opts.workspace.names[modelData] || modelData
				isMax: false
			}
		}

		Workspace {
			id: plus
			idx: Hypr.workspaces.values.length + 1
			label: "+"
			isMax: true
			visible: Hypr.workspaces.values.reduce((acc, cur) => acc + (cur.lastIpcObject.windows == "0") ? 1 : 0, 0) == 0
		}

		ToolTip {
			id: tooltip_
		}
	}
}
