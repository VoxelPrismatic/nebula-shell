import QtQuick
import "root:/config"

Rectangle {
	id: root
	required property int btnSize
	property int iconSize: btnSize / 1.5
	required property string glyph
	property color fill
	property color pressColor: Sakura.hlHigh
	property color hoverColor: Sakura.hlMed
	property color defaultColor: "transparent"
	property int cursor

	signal click
	width: btnSize
	height: btnSize

	color: btn.renderColor

	BtnIcon {
		id: btn
		size: root.iconSize
		glyph: root.glyph
		fill: root.fill
		cursor: root.cursor

		anchors.centerIn: parent
		pressColor: root.pressColor
		hoverColor: root.hoverColor
		defaultColor: root.defaultColor
		onClick: root.click()
	}
}
