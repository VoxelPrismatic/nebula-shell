pragma ComponentBehavior: Bound

import QtQuick
import QtQuick.Effects
import Quickshell

import "root:/config"
import "root:/modules"

Canvas {
	id: switcher
	width: parent.width
	height: parent.height

	readonly property bool hovered: this.trigger || area.containsMouse
	readonly property Rectangle contentBox: content
	readonly property MouseArea area_: area
	required property bool trigger
	required property bool botAttached
	required property bool topAttached

	Rectangle {
		id: content
		width: parent.width - Opts.radius
		height: parent.height - Opts.radius
		anchors.bottom: parent.bottom
		anchors.right: parent.right
		anchors.rightMargin: switcher.hovered ? 0 : -this.width
		color: Sakura.layerBase
		topLeftRadius: Opts.radius
		Behavior on anchors.rightMargin {
			NumberAnimation {
				duration: Opts.aniWidget.duration
				easing.type: Opts.aniWidget.style
			}
		}
	}

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
		color: Sakura.layerBase
		anchors.bottom: parent.bottom
		anchors.right: content.left
		botRight: true
		visible: switcher.botAttached
	}
	Corner {
		id: topLeftCorner
		px: Opts.radius
		color: Sakura.layerBase
		anchors.top: parent.top
		anchors.right: content.left
		topRight: true
		visible: switcher.topAttached
	}

	Corner {
		id: topRightCorner
		px: Opts.radius
		color: Sakura.layerBase
		anchors.bottom: content.top
		anchors.right: parent.right
		botRight: true
		anchors.rightMargin: Math.min(Opts.radius, content.width+content.anchors.rightMargin) - Opts.radius
		visible: !switcher.topAttached
	}
	Corner {
		id: botRightCorner
		px: Opts.radius
		color: Sakura.layerBase
		anchors.top: content.bottom
		anchors.right: parent.right
		topRight: true
		anchors.rightMargin: Math.min(Opts.radius, content.width+content.anchors.rightMargin) - Opts.radius
		visible: !switcher.botAttached
	}

	MouseArea {
		id: area
		hoverEnabled: true
		anchors.fill: parent
	}
}
