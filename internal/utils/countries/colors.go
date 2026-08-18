package _countries

func ColorForISO(iso string) (string, bool) {
	color, ok := isoToColor[iso]
	return color, ok
}

// isoToColor gives each ISO code the dominant/most recognizable color from
// its actual flag (best-effort approximation, not official Pantone specs).
var isoToColor = map[string]string{
	"US": "#3C3B6E", // United States — canton blue
	"RU": "#0039A6", // Russia — flag blue
	"EG": "#CE1126", // Egypt — red
	"ZA": "#007749", // South Africa — green
	"GR": "#0D5EAF", // Greece — blue
	"NL": "#AE1C28", // Netherlands — red
	"BE": "#FDDA24", // Belgium — yellow
	"FR": "#0055A4", // France — blue
	"ES": "#FFC400", // Spain — gold
	"HU": "#CE2939", // Hungary — red
	"IT": "#009246", // Italy — green
	"RO": "#002B7F", // Romania — blue
	"CH": "#D52B1E", // Switzerland — red
	"AT": "#ED2939", // Austria — red
	"GB": "#C8102E", // United Kingdom — Union Jack red
	"DK": "#C60C30", // Denmark — red

	"SE": "#006AA7", // Sweden — blue
	"NO": "#EF2B2D", // Norway — red
	"PL": "#DC143C", // Poland — red
	"DE": "#FFCE00", // Germany — gold
	"PE": "#D91023", // Peru — red
	"MX": "#006847", // Mexico — green
	"CU": "#002A8F", // Cuba — blue
	"AR": "#75AADB", // Argentina — sky blue
	"BR": "#009739", // Brazil — green
	"CL": "#0033A0", // Chile — blue
	"CO": "#FCD116", // Colombia — yellow
	"VE": "#FFCC00", // Venezuela — yellow
	"MY": "#010066", // Malaysia — canton blue
	"AU": "#00008B", // Australia — field blue
	"ID": "#CE1126", // Indonesia — red
	"PH": "#0038A8", // Philippines — blue

	"NZ": "#00247D", // New Zealand — field blue
	"SG": "#EF3340", // Singapore — red
	"TH": "#132B75", // Thailand — blue
	"JP": "#BC002D", // Japan — red disc
	"KR": "#003478", // South Korea — taegeuk blue
	"VN": "#DA251D", // Vietnam — red
	"CN": "#DE2910", // China — red
	"TR": "#E30A17", // Turkey — red
	"IN": "#FF9933", // India — saffron
	"PK": "#01411C", // Pakistan — green
	"LK": "#8D153A", // Sri Lanka — lion maroon
	"MM": "#FECB00", // Myanmar — yellow
	"IR": "#239F40", // Iran — green
	"MA": "#C1272D", // Morocco — red
	"DZ": "#006233", // Algeria — green
	"TN": "#E70013", // Tunisia — red

	"SN": "#00853F", // Senegal — green
	"NG": "#008751", // Nigeria — green
	"KE": "#006600", // Kenya — green
	"TZ": "#1EB53A", // Tanzania — green
	"UG": "#FCD116", // Uganda — yellow
	"PT": "#DA020E", // Portugal — red
	"LU": "#00A3E0", // Luxembourg — light blue
	"IE": "#169B62", // Ireland — green
	"IS": "#02529C", // Iceland — blue
	"FI": "#003580", // Finland — blue
	"BG": "#00966E", // Bulgaria — green
	"LT": "#FDB913", // Lithuania — yellow
	"LV": "#9E3039", // Latvia — carmine
	"EE": "#0072CE", // Estonia — blue
	"UA": "#0057B7", // Ukraine — blue
	"CZ": "#11457E", // Czech Republic — blue

	"SK": "#0B4EA2", // Slovakia — blue
	"HK": "#DE2910", // Hong Kong — red
	"KH": "#E00025", // Cambodia — red
	"BD": "#006A4E", // Bangladesh — green
	"TW": "#FE0000", // Taiwan — red
	"LB": "#ED1C24", // Lebanon — red
	"JO": "#CE1126", // Jordan — red
	"SY": "#CE1126", // Syria — red
	"IQ": "#CE1126", // Iraq — red
	"KW": "#007A3D", // Kuwait — green
	"SA": "#006C35", // Saudi Arabia — green
	"YE": "#CE1126", // Yemen — red
	"OM": "#DB161B", // Oman — red
	"AE": "#FF0000", // United Arab Emirates — red
	"IL": "#0038B8", // Israel — blue
	"BH": "#CE1126", // Bahrain — red

	"QA": "#8D1B3D", // Qatar — maroon
	"BT": "#FF9E1B", // Bhutan — orange
	"MN": "#C4272F", // Mongolia — red
	"NP": "#DC143C", // Nepal — crimson
	"TJ": "#CC0000", // Tajikistan — red
	"TM": "#00843D", // Turkmenistan — green
	"AZ": "#00B5E2", // Azerbaijan — turquoise
	"GE": "#FF0000", // Georgia — red
	"KG": "#E8112D", // Kyrgyzstan — red
	"UZ": "#0099B5", // Uzbekistan — blue
}
