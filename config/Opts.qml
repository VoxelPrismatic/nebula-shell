pragma Singleton

import Quickshell
import QtQuick

Singleton {
	id: root

	readonly property NebulaWorkspace workspace: NebulaWorkspace {}
	component NebulaWorkspace: QtObject {
		readonly property int columns: 3
		readonly property real cellSize: 80 / this.columns
		readonly property list<string> names: ["α", "β", "δ", "ζ", "ξ", "ϟ", "λ", "π", "μ", "τ", "ω", "ϰ"]
	}
}
