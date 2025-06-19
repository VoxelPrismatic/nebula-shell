pragma ComponentBehavior: Bound

import QtQuick
import QtQml
import QtQuick.Effects
import QtQuick.Controls
import Quickshell

import "root:/config"
import "root:/modules"

Canvas {
	id: switcher
	width: parent.width
	height: parent.height

	default property alias children: area.children
	readonly property bool hovered: this.trigger || area.containsMouse
	readonly property Rectangle contentBox: content
	readonly property MouseArea area_: area
	required property bool trigger
	required property bool botAttached
	required property bool topAttached
	property color botColor: Sakura.layerBase
	property color topColor: Sakura.layerBase

	signal mouseOver
	signal mouseExit
	signal opened
	signal closed

	onHoveredChanged: {
		if (hovered) {
			opened();
		} else {
			closed();
		}
	}

	Rectangle {
		id: content
		width: switcher.width - Opts.radius
		height: switcher.height - Opts.radius
		anchors.bottom: switcher.bottom
		anchors.right: switcher.right
		anchors.rightMargin: switcher.hovered ? 0 : -this.width - Opts.radius
		color: Sakura.layerBase
		topLeftRadius: Opts.radius
		Behavior on anchors.rightMargin {
			NumberAnimation {
				duration: Opts.aniWidget.duration
				easing.type: Opts.aniWidget.style
			}
		}
		MouseArea {
			id: area
			hoverEnabled: true
			anchors.fill: parent
			onEntered: switcher.mouseOver()
			onExited: switcher.mouseExit()
		}
	}

	ToolTip {}
	MultiEffect {
		source: content
		anchors.fill: content
		shadowBlur: 1.0
		shadowEnabled: true
		shadowColor: Sakura.paintPine
		shadowVerticalOffset: 8
		shadowHorizontalOffset: 8
	}

	Corner {
		id: botLeftCorner
		px: Opts.radius
		color: switcher.botColor
		anchors.bottom: parent.bottom
		anchors.right: content.left
		botRight: true
		visible: switcher.botAttached
	}
	Corner {
		id: topLeftCorner
		px: Opts.radius
		color: switcher.topColor
		anchors.top: parent.top
		anchors.right: content.left
		topRight: true
		visible: switcher.topAttached
	}

	Corner {
		id: topRightCorner
		px: Opts.radius
		color: switcher.topColor
		anchors.bottom: content.top
		anchors.right: parent.right
		botRight: true
		anchors.rightMargin: Math.min(Opts.radius, content.width + content.anchors.rightMargin) - Opts.radius
		visible: !switcher.topAttached
	}
	Corner {
		id: botRightCorner
		px: Opts.radius
		color: switcher.botColor
		anchors.top: content.bottom
		anchors.right: parent.right
		topRight: true
		anchors.rightMargin: Math.min(Opts.radius, content.width + content.anchors.rightMargin) - Opts.radius
		visible: !switcher.botAttached
	}
}
