pragma ComponentBehavior: Bound

import QtQuick
import QtQuick.Controls
import QtQuick.Layouts
import Quickshell
import Quickshell.Hyprland

import "root:/config"

Canvas {
	id: switcher

	required property string screen

	width: Opts.workspace.cellSize * Opts.workspace.columns
	height: Opts.workspace.cellSize * Math.ceil((Hyprland.workspaces.values.length + (plus.visible ? 1 : 0)) / Opts.workspace.columns)

	anchors.topMargin: 8
	anchors.top: parent.top
	anchors.horizontalCenter: parent.horizontalCenter

	GridLayout {
		id: grid
		columns: Opts.workspace.columns
		rows: Math.ceil(Hyprland.workspaces.values.length / Opts.workspace.columns)
		width: parent.width
		height: parent.height

		Repeater {
			model: Hyprland.workspaces.values.length

			Workspace {
				required property int modelData
				idx: modelData
				label: Opts.workspace.names[modelData] || modelData
				isMax: false
				mon: Hyprland.monitors.values.find(e => e.name == switcher.screen)
			}
		}

		Workspace {
			id: plus
			idx: Hyprland.workspaces.values.length + 1
			label: "+"
			isMax: true
			visible: Hyprland.workspaces.values.reduce((acc, cur) => acc + (cur.lastIpcObject.windows == "0") ? 1 : 0, 0) == 0
		}
	}

	Timer {
		interval: 100
		repeat: true
		running: true

		onTriggered: {
			Hyprland.refreshWorkspaces();
			Hyprland.refreshMonitors();
		}
	}
}
