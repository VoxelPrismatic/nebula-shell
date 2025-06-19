import QtQuick
import QtQuick.Controls

Canvas {
	id: root
	property int size: 24
	property color pressColor
	property color hoverColor
	property color defaultColor
	readonly property color renderColor: inPress ? pressColor : inHover ? hoverColor : defaultColor
	readonly property bool inHover: area.containsMouse
	readonly property bool inPress: area.containsPress
	required property string glyph
	width: this.size
	height: this.size
	signal leave
	signal hover
	signal click
	Button {
		width: root.size
		height: root.size
		padding: 0
		flat: true
		icon {
			name: root.glyph
			width: root.size
			height: root.size
		}
	}
	MouseArea {
		id: area
		anchors.fill: parent
		hoverEnabled: true
		onEntered: root.hover()
		onExited: root.leave()
		onClicked: root.click()
	}
}
