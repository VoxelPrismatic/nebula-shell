import QtQuick
import QtQuick.Controls
import Quickshell
import Quickshell.Hyprland
import "root:/svc"
import "root:/config"

Rectangle {
	id: root
	required property int idx
	required property string label
	required property bool isMax
	property HyprlandWorkspace ws: Hypr.workspaces.values[idx]

	property bool hyprFocused: ws?.focused || false
	property bool hyprActive: ws?.active || false
	property bool disabled: false

	width: parent.width / 3 - 4
	height: this.width

	function calcColor() {
		Hyprland.refreshWorkspaces();
		let hoverColor = Sakura.hlMed;
		let normalColor = Sakura.layerOverlay;

		if (this.hyprFocused) {
			hoverColor = Sakura.layerInverse;
			normalColor = Sakura.textNormal;
		} else if (this.ws?.lastIpcObject.windows == "0") {
			hoverColor = Sakura.paintLove;
			normalColor = Sakura.paintRose;
		} else if (this.hyprActive) {
			normalColor = Sakura.hlHigh;
		}

		this.color = area.containsMouse ? hoverColor : normalColor;
		const reverseText = [Sakura.textNormal, Sakura.paintRose];
		if (reverseText.includes(normalColor)) {
			label_.color = Sakura.layerBase;
		} else {
			label_.color = Sakura.textNormal;
		}
	}

	color: calcColor()
	onHyprActiveChanged: calcColor()
	onHyprFocusedChanged: calcColor()

	radius: 4

	Text {
		id: label_
		anchors.centerIn: parent
		text: parent.label
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
			root.calcColor();
			this.targetId = root.ws?.id || 0;
			if (!parent.isMax)
				return;

			for (var shell of Hypr.workspaces.values) {
				if (shell.id >= this.targetId) {
					this.targetId = shell.id + 1;
				}
			}
		}
		onClicked: function (mouse) {
			if (mouse.button == Qt.RightButton) {
				Hypr.dispatch(`movetoworkspace ${this.targetId}`);
			} else {
				Hypr.dispatch(`focusworkspaceoncurrentmonitor ${this.targetId}`);
			}
		}
		acceptedButtons: Qt.LeftButton | Qt.RightButton
	}
}
