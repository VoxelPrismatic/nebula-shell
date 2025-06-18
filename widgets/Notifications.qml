pragma ComponentBehavior: Bound

import QtQuick
import QtQuick.Effects
import QtQuick.Controls
import QtQuick.Layouts
import Quickshell
import Quickshell.Hyprland

import "root:/config"
import "root:/modules"

Canvas {
	id: switcher

	width: parent.width
	height: 368

	readonly property bool hovered: this.trigger || content.area_.containsMouse
	required property bool trigger

	anchors.bottom: parent.bottom
	anchors.right: parent.right
	Cute {
		id: content
		trigger: parent.trigger
		anchors.fill: parent
		botAttached: true
		topAttached: false
	}
}
