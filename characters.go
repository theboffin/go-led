package main

var characterSet = map[rune][][]int{
	// 12x7 matrix for each character (some characters are wider than others and thats ok)
	' ': {
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	},
	'A': {
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 1, 1, 1, 0, 0},
		{0, 1, 1, 0, 1, 1, 0},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
	},
	'B': {
		{1, 1, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 1, 1, 0},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 1, 0},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 0, 0},
	},
	'C': {
		{0, 1, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 1, 1, 0},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 0, 0},
	},
	'D': {
		{1, 1, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 1, 1, 0},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 0, 0},
	},
	'E': {
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
	},
	'F': {
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 1, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
	},
	'G': {
		{0, 1, 1, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 1, 1, 1, 1},
		{1, 1, 0, 0, 1, 1, 1, 1},
		{1, 1, 0, 0, 0, 1, 1, 0},
		{1, 1, 0, 0, 0, 1, 1, 0},
		{1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 0, 0},
	},
	'H': {
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
	},
	'I': {
		{1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
	},
	'J': {
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{1, 1, 0, 1, 1, 0, 0},
		{1, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 1, 0, 0, 0},
	},
	'K': {
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 1, 1, 0},
		{1, 1, 0, 1, 1, 0, 0},
		{1, 1, 1, 1, 0, 0, 0},
		{1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 0, 0},
		{1, 1, 0, 1, 1, 1, 0},
		{1, 1, 0, 0, 1, 1, 0},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
	},
	'L': {
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
	},
	'M': {
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 1, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
	},
	'N': {
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 1, 0, 0, 0, 1, 1},
		{1, 1, 1, 0, 0, 0, 1, 1},
		{1, 1, 1, 1, 0, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 1},
		{1, 1, 0, 0, 1, 1, 1, 1},
		{1, 1, 0, 0, 0, 1, 1, 1},
		{1, 1, 0, 0, 0, 1, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
	},
	'O': {
		{0, 1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 0},
	},
	'P': {
		{1, 1, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 1, 1, 0},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
	},
	'Q': {
		{0, 1, 1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 1},
		{1, 1, 0, 0, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 0, 0, 1, 1},
	},
	'R': {
		{1, 1, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 1, 1, 0},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 1, 0},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
	},
	'S': {
		{0, 0, 1, 1, 1, 0, 0},
		{0, 1, 1, 1, 1, 1, 0},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 0, 0},
	},
	'T': {
		{1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
	},
	'U': {
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 0},
	},
	'V': {
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1, 1},
		{0, 1, 1, 0, 1, 1, 0},
		{0, 1, 1, 0, 1, 1, 0},
		{0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
	},
	'W': {
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 1, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
	},
	'X': {
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{0, 1, 1, 0, 0, 1, 1, 0},
		{0, 1, 1, 0, 0, 1, 1, 0},
		{0, 0, 1, 1, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 1, 1, 1, 1, 0, 0},
		{0, 1, 1, 0, 0, 1, 1, 0},
		{0, 1, 1, 0, 0, 1, 1, 0},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
	},
	'Y': {
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{0, 1, 1, 0, 0, 1, 1, 0},
		{0, 1, 1, 0, 0, 1, 1, 0},
		{0, 0, 1, 1, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0},
	},
	'Z': {
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 1, 1, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
	},
}

func getLedLines(message string, ledOn string, ledOff string) ([]string, int) {
	var lines []string
	var ledWidth = 0
	for _, char := range message {
		if matrix, ok := characterSet[char]; ok {
			for i, row := range matrix {
				var line string
				for _, col := range row {
					if col == 1 {
						line += ledOn
					} else {
						line += ledOff
					}
					ledWidth++
				}
				line += ledOff + ledOff // Add a space between characters
				ledWidth++              // increment the width for the space
				if i > len(lines)-1 {
					lines = append(lines, line)
				} else {
					lines[i] += line
				}
			}
		}
	}
	return lines, ledWidth
}
