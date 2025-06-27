package sakura

var Vectors = SakuraSwatch[func(uint) uint]{
	Dawn: SakuraPalette[func(uint) uint]{
		Hl: SakuraHl[func(uint) uint]{
			High: HSLVector{282.69230769231, -17.254901960784, -3.2582524271845}.Tx,
			Low:  HSLVector{-7.3076923076923, -2.3529411764706, -0.28196721311474}.Tx,
			Med:  HSLVector{-22.307692307692, -10.588235294118, -2.5094170403587}.Tx,
		},
		Layer: SakuraLayer[func(uint) uint]{
			Base:    HSLVector{0, 0, 0}.Tx,
			Overlay: HSLVector{-4.0723981900453, -3.1372549019608, 1.8247933884298}.Tx,
			Surface: HSLVector{2.6923076923076, 1.9607843137255, -0.49411764705881}.Tx,
		},
		Text: SakuraText[func(uint) uint]{
			Normal: HSLVector{0, 0, 0}.Tx,
			Muted:  HSLVector{8.974358974359, 17.254901960784, -21.322314049587}.Tx,
			Subtle: HSLVector{0.30769230769235, 10.196078431373, -11.823241693372}.Tx,
		},
		Paint: SakuraPaint[func(uint) uint]{
			Love: HSLVector{0, 0, 0}.Tx,
			Gold: HSLVector{0, 0, 0}.Tx,
			Rose: HSLVector{0, 0, 0}.Tx,
			Pine: HSLVector{0, 0, 0}.Tx,
			Foam: HSLVector{0, 0, 0}.Tx,
			Iris: HSLVector{0, 0, 0}.Tx,
			Tree: HSLVector{0, 0, 0}.Tx,
		},
	},
	Main: SakuraPalette[func(uint) uint]{
		Hl: SakuraHl[func(uint) uint]{
			High: HSLVector{-1.7307692307693, 26.274509803922, -12.810140237325}.Tx,
			Low:  HSLVector{-4.9450549450549, 3.921568627451, -5.6763285024155}.Tx,
			Med:  HSLVector{-0.6593406593407, 18.039215686275, -10.50135501355}.Tx,
		},
		Layer: SakuraLayer[func(uint) uint]{
			Base:    HSLVector{0, 0, 0}.Tx,
			Overlay: HSLVector{-1.4046822742475, 8.6274509803922, 3.544061302682}.Tx,
			Surface: HSLVector{-2.1719457013575, 3.921568627451, 0.84541062801932}.Tx,
		},
		Text: SakuraText[func(uint) uint]{
			Normal: HSLVector{0, 0, 0}.Tx,
			Muted:  HSLVector{3.1168831168831, -43.137254901961, 11.879128945437}.Tx,
			Subtle: HSLVector{2.5454545454545, -29.019607843137, 8.6306653809064}.Tx,
		},
		Paint: SakuraPaint[func(uint) uint]{
			Foam: HSLVector{-0.041095890410929, 22.352941176471, -18.134171907757}.Tx,
			Gold: HSLVector{0.34524530587521, 4.7058823529412, -26.151761517615}.Tx,
			Iris: HSLVector{-0.89760638297872, 24.313725490196, -0.10502318194625}.Tx,
			Love: HSLVector{0.10155316606932, 21.56862745098, 7.7659574468085}.Tx,
			Pine: HSLVector{0.091185410334361, 4.7058823529412, -3.7313831206961}.Tx,
			Rose: HSLVector{-0.24764962164638, 7.843137254902, -20.544285007422}.Tx,
			Tree: HSLVector{-0.12084172511982, 4.7058823529412, -3.4704246956188}.Tx,
		},
	},
	Moon: SakuraPalette[func(uint) uint]{
		Hl: SakuraHl[func(uint) uint]{
			High: HSLVector{2.8571428571429, 21.960784313725, -13.434343434343}.Tx,
			Low:  HSLVector{-0.25974025974025, 3.1372549019608, -3.4050179211469}.Tx,
			Med:  HSLVector{1.4857142857143, 14.117647058824, -11.111111111111}.Tx,
		},
		Layer: SakuraLayer[func(uint) uint]{
			Base:    HSLVector{0, 0, 0}.Tx,
			Overlay: HSLVector{2.5615763546798, 10.980392156863, -3.5230352303523}.Tx,
			Surface: HSLVector{1.7857142857143, 3.5294117647059, -0.79365079365079}.Tx,
		},
		Text: SakuraText[func(uint) uint]{
			Normal: HSLVector{0, 0, 0}.Tx,
			Muted:  HSLVector{3.1168831168831, -43.137254901961, 11.879128945437}.Tx,
			Subtle: HSLVector{2.5454545454545, -29.019607843137, 8.6306653809064}.Tx,
		},
		Paint: SakuraPaint[func(uint) uint]{
			Foam: HSLVector{-0.041095890410929, 22.352941176471, -18.134171907757}.Tx,
			Gold: HSLVector{0.34524530587521, 4.7058823529412, -26.151761517615}.Tx,
			Iris: HSLVector{-0.89760638297872, 24.313725490196, -0.10502318194625}.Tx,
			Love: HSLVector{0.10155316606932, 21.56862745098, 7.7659574468085}.Tx,
			Pine: HSLVector{0.22556390977445, 17.647058823529, -4.6929215822346}.Tx,
			Rose: HSLVector{-0.52795451468796, 7.4509803921569, -5.9252633671238}.Tx,
			Tree: HSLVector{-0.45378151260505, 17.647058823529, -4.775828460039}.Tx,
		},
	},
}
