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

	function jaccardSimilarity(str1: string, str2: string): double {
		var words1 = str1.toLowerCase().split(/\s+/);
		var words2 = str2.toLowerCase().split(/\s+/);
		var set1 = new Set(words1);
		var set2 = new Set(words2);
		var intersection = new Set([...set1].filter(x => set2.has(x)));
		var union = new Set([...set1, ...set2]);
		return intersection.size / union.size;
	}

	function locateIcon() {
		var bestEntry = null;
		var bestSimilarity = 0;
		const initialTitle = Tiles.tiles.find(t => t.wmClass == root.app.appId).initialTitle;
		for (var entry of DesktopEntries.applications.values) {
			if (!entry.name)
				continue;
			var similarity = jaccardSimilarity(initialTitle, entry.name);
			if (similarity > bestSimilarity) {
				bestSimilarity = similarity;
				bestEntry = entry;
			}
		}
		if (bestEntry) {
			var iconFromEntry = Quickshell.iconPath(bestEntry.icon, false);
			if (iconFromEntry !== "") {
				root.entry = bestEntry;
				return bestEntry.icon;
			}
		}
		return "application-x-executable";
	}

	Button {
		id: ico
		width: 24
		height: 24
		anchors.centerIn: parent
		icon {
			name: root.entry?.icon || root.locateIcon()
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
				ToolTip.show(root.entry.name || root.app.title || root.app.appId);
			} else {
				ToolTip.hide();
			}
		}
		acceptedButtons: Qt.LeftButton | Qt.RightButton
	}
}
