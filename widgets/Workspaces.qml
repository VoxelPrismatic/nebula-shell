pragma ComponentBehavior: Bound

import QtQuick
import QtQuick.Layouts
import Quickshell
import Quickshell.Hyprland

Canvas {
	id: switcher

	width: 80
	height: 26 * Math.ceil((Hyprland.workspaces.values.length + 1) / 3)

	anchors.topMargin: 8
	anchors.top: parent.top
	anchors.horizontalCenter: parent.horizontalCenter

	GridLayout {
		columns: 3
		rows: Math.ceil(Hyprland.workspaces.values.length / 3)
		width: parent.width
		height: parent.height

		Workspace {
			idx: 0
			label: "α"
		}

		Workspace {
			idx: 1
			label: "β"
		}

		Workspace {
			idx: 2
			label: "δ"
		}

		Workspace {
			idx: 3
			label: "ζ"
		}

		Workspace {
			idx: 4
			label: "ξ"
		}

		Workspace {
			idx: 5
			label: "ϟ"
		}

		Workspace {
			idx: 6
			label: "λ"
		}

		Workspace {
			idx: 7
			label: "π"
		}

		Workspace {
			idx: 8
			label: "μ"
		}

		Workspace {
			idx: 9
			label: "τ"
		}

		Workspace {
			idx: 10
			label: "ω"
		}

		Workspace {
			idx: 11
			label: "ϰ"
		}
	}
}
