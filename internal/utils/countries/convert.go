package _countries

import "strings"

func CountryCodeToISO(rawCode string) (string, bool) {
	code := strings.TrimPrefix(strings.TrimSpace(rawCode), "+")
	if code == "" {
		return "", false
	}
	iso, ok := countryCodeToISO[code]
	return iso, ok
}

var countryCodeToISO = map[string]string{
	"1": "US", "7": "RU", "20": "EG", "27": "ZA", "30": "GR", "31": "NL",
	"32": "BE", "33": "FR", "34": "ES", "36": "HU", "39": "IT", "40": "RO",
	"41": "CH", "43": "AT", "44": "GB", "45": "DK", "46": "SE", "47": "NO",
	"48": "PL", "49": "DE", "51": "PE", "52": "MX", "53": "CU", "54": "AR",
	"55": "BR", "56": "CL", "57": "CO", "58": "VE", "60": "MY", "61": "AU",
	"62": "ID", "63": "PH", "64": "NZ", "65": "SG", "66": "TH", "81": "JP",
	"82": "KR", "84": "VN", "86": "CN", "90": "TR", "91": "IN", "92": "PK",
	"94": "LK", "95": "MM", "98": "IR", "212": "MA", "213": "DZ", "216": "TN",
	"221": "SN", "234": "NG", "254": "KE", "255": "TZ", "256": "UG",
	"351": "PT", "352": "LU", "353": "IE", "354": "IS", "358": "FI",
	"359": "BG", "370": "LT", "371": "LV", "372": "EE", "380": "UA",
	"420": "CZ", "421": "SK", "852": "HK", "855": "KH", "880": "BD",
	"886": "TW", "961": "LB", "962": "JO", "963": "SY", "964": "IQ",
	"965": "KW", "966": "SA", "967": "YE", "968": "OM", "971": "AE",
	"972": "IL", "973": "BH", "974": "QA", "975": "BT", "976": "MN",
	"977": "NP", "992": "TJ", "993": "TM", "994": "AZ", "995": "GE",
	"996": "KG", "998": "UZ",
}
