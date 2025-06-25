import QtQuick

Item {
	id: root

	property double hue
	property double sat
	property double lum

	function rgb() {
		let _hue = this.hue / 360;
		let _sat = Math.min(Math.max(this.sat / 100, 0), 1);
		let _lum = Math.min(Math.max(this.lum / 100, 0), 1);

		let intPart = Math.floor(_hue * 6);
		let floatPart = _hue * 6 - intPart;

		let prime = _lum * (1 - _sat);
		let quart = _lum * (1 - floatPart * _sat);
		let third = _lum * (1 - (1 - floatPart) * _sat);

		let r = ([_lum, quart, prime, prime, third, _lum])[intPart];
		let g = ([third, _lum, _lum, quart, prime, prime])[intPart];
		let b = ([prime, prime, third, _lum, _lum, quart])[intPart];

		let rgbObj = Qt.createComponent("RGB.qml");
		if (rgbObj.status != Component.Ready) {
			console.error("Error loading RGB.qml:", rgbObj.errorString());
			return null;
		}

		return rgbObj.createObject(root, {
			red: Math.min(Math.max(Math.round(r * 255), 0), 255),
			green: Math.min(Math.max(Math.round(g * 255), 0), 255),
			blue: Math.min(Math.max(Math.round(b * 255), 0), 255)
		});
	}
}
