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
		WlrLayershell.layer: WlrLayer.Bottom
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
			WlrLayershell.layer: WlrLayer.Overlay

			Workspaces {
				id: workspaces
				screen: root.modelData.name
			}
			Line {
				id: workspace_separator
				anchors.topMargin: 8
				anchors.top: workspaces.bottom
			}
			Tasks {

				anchors.topMargin: 8
				anchors.top: workspace_separator.bottom
			}
			Clock {
				id: clockWidget
			}
		}

		Corner {
			id: lowerCorner
			px: Opts.radius
			color: Sakura.layerBase
			botRight: true
			anchors.right: parent.right
			anchors.rightMargin: dock.width
			anchors.bottom: parent.bottom
		}

		Corner {
			id: upperCorner
			px: Opts.radius
			color: lowerCorner.color
			topRight: true
			anchors.right: lowerCorner.anchors.right
			anchors.rightMargin: lowerCorner.anchors.rightMargin
			anchors.top: parent.top
		}

		PanelWindow {
			id: widgets
			anchors {
				top: true
				bottom: true
				right: true
			}

			implicitWidth: 512

			WlrLayershell.exclusionMode: ExclusionMode.Ignore
			color: "transparent"
			screen: root.modelData
			aboveWindows: this.shown
			readonly property bool shown: notifWidget.hovered
			margins.right: dock.implicitWidth

			Behavior on aboveWindows {
				NumberAnimation {
duration: widgets.shown ? Opts.aniWidget.duration + 50 : 0
				}
			}

			Notifications {
				id: notifWidget
				trigger: clockWidget.area_.containsMouse
			}
		}
	}
}
