pragma Singleton

import Quickshell
import QtQuick

Singleton {
	id: root

	readonly property int notifCount: 5

	readonly property int radius: 12
	readonly property NebulaAnimation aniWidget: NebulaAnimation {
		duration: 150
		style: Easing.OutCubic
	}

	readonly property NebulaAnimation aniHover: NebulaAnimation {
		duration: 150
		style: Easing.InOutQuad
	}

	readonly property NebulaWorkspace workspace: NebulaWorkspace {
		columns: 3
		cellSize: 80 / this.columns
		names: ["α", "β", "δ", "ζ", "ξ", "ϟ", "λ", "π", "μ", "τ", "ω", "ϰ"]
	}
	component NebulaWorkspace: QtObject {
		required property int columns
		required property real cellSize
		required property list<string> names
	}
	component NebulaAnimation: QtObject {
		required property int duration
		required property var style
	}
}
