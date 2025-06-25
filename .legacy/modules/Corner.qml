import QtQuick
import QtQuick.Effects
import Quickshell
import Quickshell.Wayland

Item {
	id: root
	required property real px
	required property color color
	property bool topRight
	property bool topLeft
	property bool botRight
	property bool botLeft

	width: root.px * 2
	height: root.px * 2

	Rectangle {
		id: src
		width: parent.width
		height: parent.height
		color: root.color
		visible: false
		smooth: true
	}

	Rectangle {
		id: mask
		width: parent.width
		height: parent.height
		color: "white"
		topRightRadius: root.topRight ? parent.width : 0
		topLeftRadius: root.topLeft ? parent.width : 0
		bottomRightRadius: root.botRight ? parent.width : 0
		bottomLeftRadius: root.botLeft ? parent.width : 0
		smooth: true
	}

	MultiEffect {
		source: src
		anchors.fill: src
		maskEnabled: true
		smooth: true
		maskSource: ShaderEffectSource {
			sourceItem: mask
			hideSource: true
			smooth: true
			antialiasing: true
		}
		maskInverted: true
		antialiasing: true
		maskThresholdMin: 0.5
		maskThresholdMax: 1.0
		maskSpreadAtMin: 1.0
		maskSpreadAtMax: 0.0
	}
}
