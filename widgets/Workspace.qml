import QtQuick
import QtQuick.Controls
import Quickshell
import Quickshell.Hyprland
import "root:/config"

Rectangle {
	id: root
	required property int idx
	required property string label
	required property bool isMax
	property HyprlandMonitor mon
	property HyprlandWorkspace ws: Hyprland.workspaces.values[idx]

	property bool hyprFocused: (ws && mon?.activeWorkspace == ws) || false
	property bool hyprActive: ws?.active || false
	property bool disabled: false

	width: parent.width / 3 - 4
	height: this.width

	property var hoverColor: hyprFocused ? Sakura.layerInverse : ws?.lastIpcObject.windows == "0" ? Sakura.paintLove : Sakura.hlMed
	property var normalColor: hyprFocused ? Sakura.textNormal : ws?.lastIpcObject.windows == "0" ? Sakura.paintRose : hyprActive ? Sakura.hlHigh : Sakura.layerOverlay

	color: area.containsMouse ? hoverColor : normalColor

	radius: 4

	Text {
		id: label_
		anchors.centerIn: parent
		text: parent.label
		color: ([Sakura.textNormal, Sakura.paintRose]).includes(parent.normalColor) ? Sakura.layerBase : Sakura.textNormal
		font {
			family: "Ubuntu Nerd Font"
			weight: parent.hyprFocused ? 700 : 100
			pixelSize: 16
		}
	}

	ToolTip {}

	MouseArea {
		id: area
		anchors.fill: parent
		cursorShape: Qt.PointingHandCursor
		hoverEnabled: true
		property int targetId
		onHoveredChanged: {
			this.targetId = root.ws?.id || 0;
			if (!parent.isMax)
				return;

			for (var shell of Hyprland.workspaces.values) {
				if (shell.id >= this.targetId) {
					this.targetId = shell.id + 1;
				}
			}
		}
		onClicked: function (mouse) {
			if (mouse.button == Qt.RightButton) {
				Hyprland.dispatch(`movetoworkspace ${this.targetId}`);
			} else {
				Hyprland.dispatch(`focusworkspaceoncurrentmonitor ${this.targetId}`);
			}
			Hyprland.refreshWorkspaces();
			Hyprland.refreshMonitors();
		}
		acceptedButtons: Qt.LeftButton | Qt.RightButton
	}
}
