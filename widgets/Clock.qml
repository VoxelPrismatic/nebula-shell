import Quickshell
import Quickshell.Io
import QtQuick

import "root:/config"

Canvas {
	id: clock

	width: 64
	height: 64

	anchors.bottomMargin: 16
	anchors.bottom: parent.bottom
	anchors.horizontalCenter: parent.horizontalCenter
	// anchors.centerIn: parent

	property int hour: 0
	property int minute: 0
	property int milli: 0

	Timer {
		interval: 100
		running: true
		repeat: true
		onTriggered: {
			var now = new Date();
			clock.hour = now.getHours() % 12;
			clock.minute = now.getMinutes();
			clock.milli = now.getMilliseconds() + now.getSeconds() * 1000;
		}
	}

	Rectangle {
		id: clockBody
		antialiasing: true

		property int handWidth: 2

		anchors.centerIn: parent
		width: clock.width
		height: clockBody.width
		color: "transparent"
		border.color: Sakura.textNormal

		radius: clockBody.width

		Rectangle {
			id: hourHand
			width: clockBody.handWidth
			height: clockBody.height / 4
			antialiasing: clockBody.antialiasing

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

			transform: Rotation {
				origin.x: secondHand.width / 2
				origin.y: secondHand.height
				angle: (clock.milli / 1000) * 6
			}
		}

		Rectangle {
			id: pin
			width: clockBody.handWidth * 2
			height: clockBody.handWidth * 2
			radius: clockBody.handWidth * 2
			color: Sakura.textNormal
			anchors.centerIn: parent
		}
	}
}
