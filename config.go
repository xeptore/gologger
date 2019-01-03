package gologger

var config = configType{
	Colors: map[string]colorKind{
		"grey": colorKind{
			Dark:  238,
			Light: 249,
		},
		"red": colorKind{
			Dark:  1,
			Light: 196,
		},
		"green": colorKind{
			Dark:  28,
			Light: 40,
		},
		"yellow": colorKind{
			Dark:  3,
			Light: 11,
		},
		"blue": colorKind{
			Dark:  32,
			Light: 38,
		},
		"purple": colorKind{
			Dark:  5,
			Light: 13,
		},
		"cyan": colorKind{
			Dark:  6,
			Light: 14,
		},
		"white": colorKind{
			Dark:  7,
			Light: 15,
		},
		"black": colorKind{
			Dark:  16,
			Light: 8,
		},
		"orange": colorKind{
			Dark:  202,
			Light: 208,
		},
	},
	Styles: map[string]int{
		"reset":         0,
		"bold":          1,
		"faint":         2,
		"italic":        3,
		"underline":     4,
		"blink-slow":    5,
		"blink-rapid":   6,
		"reverse-video": 7,
		"crossed-out":   9,
		"overlined":     53,
	},
	ColorReset: 15,
}
