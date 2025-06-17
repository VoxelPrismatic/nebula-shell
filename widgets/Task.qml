import QtQuick
import QtQuick.Controls
import Quickshell
import Quickshell.Wayland
import "root:/config"

Rectangle {
	id: root
	required property bool isMax
	required property Toplevel app

	property DesktopEntry entry: DesktopEntries.byId(app.appId)

	width: parent.width / 3 - 4
	height: this.width

	color: Sakura.layerBase
	radius: 4

	Rectangle {
		id: bg
		anchors.fill: parent
		color: Sakura.paintRose
		opacity: area.containsMouse ? 0 : 0.2
		radius: parent.radius
	}

	function iconLocator() {
	}

	Button {
		id: ico
		width: 24
		height: 24
		anchors.centerIn: parent
		icon {
			name: root.entry?.icon.trim()
			width: 24
			height: 24
		}
		flat: true
		padding: 0
		smooth: false
	}

	Rectangle {
		id: hover
		anchors.fill: parent
		color: Sakura.layerInverse
		opacity: area.containsMouse ? 0.2 : 0
		radius: parent.radius
	}

	ToolTip {}

	MouseArea {
		id: area
		anchors.fill: parent
		cursorShape: Qt.PointingHandCursor
		hoverEnabled: true
		property int targetId
		onHoveredChanged: function () {
			if (area.containsMouse) {
				ToolTip.show(ico.icon.name);
			} else {
				ToolTip.hide();
			}
		}
		acceptedButtons: Qt.LeftButton | Qt.RightButton
	}
}
