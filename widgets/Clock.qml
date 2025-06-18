import Quickshell
import Quickshell.Io
import QtQuick

import "root:/config"

Canvas {
	id: clock

	width: 96
	height: 96

	// anchors.bottomMargin: 16
	anchors.bottom: parent.bottom
	anchors.horizontalCenter: parent.horizontalCenter
	// anchors.centerIn: parent

	property int hour: 0
	property int minute: 0
	property int milli: 0
	property int month: 0
	property int day: 0

	Timer {
		interval: 100
		running: true
		repeat: true
		onTriggered: {
			var now = new Date();
			clock.hour = now.getHours() % 12;
			clock.minute = now.getMinutes();
			clock.month = now.getMonth();
			clock.day = now.getDate();
			clock.milli = now.getMilliseconds() + now.getSeconds() * 1000;
		}
	}

	Rectangle {
		id: pin
		width: clockBody.handWidth * 2
		height: clockBody.handWidth * 2
		radius: clockBody.handWidth * 2
		color: Sakura.textNormal
		anchors.centerIn: parent
		z: 1000
		opacity: area.containsMouse ? 0.1 : 1

		Behavior on opacity {
			NumberAnimation {
				duration: Opts.aniHover.duration
				easing.type: Opts.aniHover.style
			}
		}
	}

	Text {
		anchors.horizontalCenter: parent.horizontalCenter
		anchors.bottom: pin.top
		anchors.bottomMargin: 0
		text: clock.hour.toString().padStart(2, "0")
		color: area.containsMouse ? Sakura.textMuted : Sakura.layerBase
		font {
			family: "Ubuntu Nerd Font"
			pixelSize: 18
			weight: 800
		}

		Behavior on color {
			ColorAnimation {
				duration: Opts.aniHover.duration
				easing.type: Opts.aniHover.style
			}
		}
	}

	Text {
		id: minuteLabel
		anchors.horizontalCenter: parent.horizontalCenter
		anchors.top: pin.bottom
		anchors.topMargin: 0
		text: clock.minute.toString().padStart(2, "0")
		color: area.containsMouse ? Sakura.textMuted : Sakura.layerBase
		font {
			family: "Ubuntu Nerd Font"
			pixelSize: 18
			weight: 800
		}

		Behavior on color {
			ColorAnimation {
				duration: Opts.aniHover.duration
				easing.type: Opts.aniHover.style
			}
		}
	}

	Text {
		anchors.verticalCenter: parent.verticalCenter
		anchors.right: pin.left
		anchors.rightMargin: 2
		property list<string> months: ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]
		text: months[clock.month]
		color: Sakura.paintIris
		font {
			family: "Ubuntu Nerd Font"
			pixelSize: 12
		}
		opacity: area.containsMouse ? 1 : 0

		Behavior on opacity {
			NumberAnimation {
				duration: Opts.aniHover.duration
				easing.type: Opts.aniHover.style
			}
		}
	}
	Text {
		anchors.verticalCenter: parent.verticalCenter
		anchors.right: clockBody.right
		anchors.rightMargin: 8
		text: clock.day.toString().padStart(2, "0")
		color: Sakura.paintIris
		font {
			family: "Ubuntu Nerd Font"
			pixelSize: 12
			weight: area.containsMouse ? 400 : 800
		}

		Behavior on font.weight {
			NumberAnimation {
				duration: Opts.aniHover.Duration
				easing.type: Opts.aniHover.style
			}
		}
	}

	Rectangle {
		id: clockBody
		antialiasing: true

		property int handWidth: 2

		anchors.centerIn: parent
		width: 64
		height: clockBody.width
		color: "transparent"
		border.color: Sakura.textNormal

		radius: clockBody.width
		opacity: area.containsMouse ? 0.1 : 1

		Behavior on opacity {
			NumberAnimation {
				duration: Opts.aniHover.duration
				easing.type: Opts.aniHover.style
			}
		}

		Rectangle {
			id: hourHand
			width: clockBody.handWidth
			height: clockBody.height / 4
			antialiasing: clockBody.antialiasing
			radius: 100

			color: Sakura.paintPine
			anchors.bottom: parent.verticalCenter
			anchors.horizontalCenter: parent.horizontalCenter

			transform: Rotation {
				origin.x: hourHand.width / 2
				origin.y: hourHand.height
				angle: (clock.hour + clock.minute / 60) * 30
			}
		}

		Rectangle {
			id: minuteHand
			width: clockBody.handWidth
			height: clockBody.height / 2
			antialiasing: clockBody.antialiasing

			color: Sakura.paintPine
			anchors.bottom: parent.verticalCenter
			anchors.horizontalCenter: parent.horizontalCenter
			radius: 100

			transform: Rotation {
				origin.x: minuteHand.width / 2
				origin.y: minuteHand.height
				angle: (clock.minute + clock.milli / 60000) * 6
			}
		}

		Rectangle {
			id: secondHand
			width: clockBody.handWidth / 2
			height: clockBody.height / 2

			color: Sakura.paintLove
			anchors.bottom: parent.verticalCenter
			anchors.horizontalCenter: parent.horizontalCenter
			antialiasing: clockBody.antialiasing
			radius: 100

			transform: Rotation {
				origin.x: secondHand.width / 2
				origin.y: secondHand.height
				angle: (clock.milli / 1000) * 6
			}
		}
	}

	MouseArea {
		id: area
		hoverEnabled: true
		anchors.fill: clock
	}

	readonly property MouseArea area_: area
}
