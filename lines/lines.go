package lines

func GetAllLines() ([]Line, error) {
	var allLines []Line
	allLines = append(allLines, Line{Color: "BLUE"})
	allLines = append(allLines, Line{Color: "GREEN"})
	allLines = append(allLines, Line{Color: "ORANGE"})
	allLines = append(allLines, Line{Color: "RED"})
	allLines = append(allLines, Line{Color: "YELLOW"})
	return allLines, nil
}
