import QtQuick
import Quickshell
import Quickshell.Hyprland
import "root:/config"

Rectangle {
	required property int idx
	required property string label
	property HyprlandWorkspace ws: Hyprland.workspaces.values[idx]

	property bool oob: idx >= Hyprland.workspaces.values.length
	property bool hyprFocused: ws?.focused
	property bool hyprActive: oob ? false : ws?.active

	width: parent.width / 3 - 4
	height: parent.height / Math.ceil((Hyprland.workspaces.values.length + 1) / 3) - 4

	color: area.containsMouse ? hyprFocused ? Sakura.layerInverse : Sakura.hlMed : hyprFocused ? Sakura.textNormal : hyprActive ? Sakura.hlHigh : Sakura.layerOverlay
	radius: 4
	visible: idx <= Hyprland.workspaces.values.length

	Text {
		anchors.centerIn: parent
		text: parent.idx == Hyprland.workspaces.values.length ? "+" : parent.label
		color: parent.hyprFocused ? Sakura.layerBase : Sakura.textNormal
		font {
			family: "Ubuntu Nerd Font"
			weight: parent.hyprFocused ? 700 : 100
			pixelSize: 16
		}
	}

	MouseArea {
		id: area
		anchors.fill: parent
		cursorShape: Qt.PointingHandCursor
		hoverEnabled: true
		onClicked: {
			console.debug(area.pressedButtons);
			if (mouse.button == Qt.RightButton) {
				Hyprland.dispatch(`movetoworkspace ${parent.oob ? parent.idx + 25 : parent.ws.id}`);
			} else {
				Hyprland.dispatch(`focusworkspaceoncurrentmonitor ${parent.oob ? parent.idx + 25 : parent.ws.id}`);
			}
		}
		acceptedButtons: Qt.LeftButton | Qt.RightButton
	}
}
