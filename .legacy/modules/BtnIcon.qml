import QtQuick
import QtQuick.Controls

Canvas {
	id: root
	property int size: 24
	property color pressColor
	property color hoverColor
	property color defaultColor
	property color fill
	readonly property color renderColor: inPress ? pressColor : inHover ? hoverColor : defaultColor
	readonly property bool inHover: area.containsMouse
	readonly property bool inPress: area.containsPress
	required property string glyph
	property int cursor: Qt.PointingHandCursor
	width: this.size
	height: this.size
	signal leave
	signal hover
	signal click
	signal dblClick
	signal press
	signal release
	Button {
		width: root.size
		height: root.size
		padding: 0
		flat: true
		icon {
			name: root.glyph
			width: root.size
			height: root.size
			color: root.fill
		}
	}
	MouseArea {
		id: area
		anchors.fill: parent
		hoverEnabled: true
		onEntered: root.hover()
		onExited: root.leave()
		onClicked: root.click()
		onDoubleClicked: root.dblClick()
		onPressed: root.press()
		onReleased: root.release()
		cursorShape: root.cursor
	}
}
