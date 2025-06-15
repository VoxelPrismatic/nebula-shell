pragma ComponentBehavior: Bound

import QtQuick
import QtQuick.Effects
import Quickshell
import Quickshell.Wayland
import "root:/widgets"
import "root:/config"

Variants {
	id: help
	model: Quickshell.screens

	PanelWindow {
		id: root
		required property ShellScreen modelData
		screen: modelData
		WlrLayershell.exclusionMode: ExclusionMode.Ignore
		WlrLayershell.layer: WlrLayer.Background
		color: "transparent"
		anchors {
			top: true
			bottom: true
			right: true
			left: true
		}

		PanelWindow {
			id: dock
			anchors {
				top: true
				bottom: true
				right: true
			}

			implicitWidth: 96
			screen: root.modelData

			color: Sakura.layerBase

			Clock {}
		}

		Corner {
			px: 16
			color: Sakura.layerBase
			botRight: true
			anchors.right: parent.right
			anchors.rightMargin: dock.width
			anchors.bottom: parent.bottom
		}

		Corner {
			px: 16
			color: Sakura.layerBase
			topRight: true
			anchors.right: parent.right
			anchors.rightMargin: dock.width
			anchors.top: parent.top
		}
	}
}
