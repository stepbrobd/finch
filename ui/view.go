package ui

func MinMax(array []float32) (float32, float32) {
	var max float32 = array[0]
	var min float32 = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func GetColorZones(slice [][]float32) []float32 {
	result := []float32{}

	if len(slice) == 0 {
		return []float32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	}

	min, max := MinMax(slice[len(slice)-1])
	for _, row := range slice {
		rowMin, rowMax := MinMax(row)
		if rowMin < min {
			min = rowMin
		}
		if rowMax > max {
			max = rowMax
		}
	}

	zoneSize := (max - min) / 11
	for i := 0; i < 11; i++ {
		result = append(result, min+zoneSize*float32(i))
	}

	return result
}

func MkViewRed(slice [][]float32, index int) string {
	result := ""

	if len(slice) == 0 {
		return result
	}

	colorZones := GetColorZones(slice)

	for _, value := range slice[index] {
		for i := 0; i < 11; i++ {
			// less or equal than zone 0
			if value <= colorZones[0] {
				result += StyleRed050.Render("█")
				break
			}
			// larger than zone 0, less or equal than zone 1
			if value > colorZones[0] && value <= colorZones[1] {
				result += StyleRed100.Render("█")
				break
			}
			// larger than zone 1, less or equal than zone 2
			if value > colorZones[1] && value <= colorZones[2] {
				result += StyleRed200.Render("█")
				break
			}
			// larger than zone 2, less or equal than zone 3
			if value > colorZones[2] && value <= colorZones[3] {
				result += StyleRed300.Render("█")
				break
			}
			// larger than zone 3, less or equal than zone 4
			if value > colorZones[3] && value <= colorZones[4] {
				result += StyleRed400.Render("█")
				break
			}
			// larger than zone 4, less or equal than zone 5
			if value > colorZones[4] && value <= colorZones[5] {
				result += StyleRed500.Render("█")
				break
			}
			// larger than zone 5, less or equal than zone 6
			if value > colorZones[5] && value <= colorZones[6] {
				result += StyleRed600.Render("█")
				break
			}
			// larger than zone 6, less or equal than zone 7
			if value > colorZones[6] && value <= colorZones[7] {
				result += StyleRed700.Render("█")
				break
			}
			// larger than zone 7, less or equal than zone 8
			if value > colorZones[7] && value <= colorZones[8] {
				result += StyleRed800.Render("█")
				break
			}
			// larger than zone 8, less or equal than zone 9
			if value > colorZones[8] && value <= colorZones[9] {
				result += StyleRed900.Render("█")
				break
			}
			// larger than zone 9, less or equal than zone 10
			if value > colorZones[9] && value <= colorZones[10] {
				result += StyleRed950.Render("█")
				break
			}
			// larger than zone 10
			if value > colorZones[10] {
				result += StyleRed950.Render("█")
				break
			}
		}
	}

	return result
}

func MkViewGreen(slice [][]float32, index int) string {
	result := ""

	if len(slice) == 0 {
		return result
	}

	colorZones := GetColorZones(slice)

	for _, value := range slice[index] {
		for i := 0; i < 11; i++ {
			// less or equal than zone 0
			if value <= colorZones[0] {
				result += StyleGreen050.Render("█")
				break
			}
			// larger than zone 0, less or equal than zone 1
			if value > colorZones[0] && value <= colorZones[1] {
				result += StyleGreen100.Render("█")
				break
			}
			// larger than zone 1, less or equal than zone 2
			if value > colorZones[1] && value <= colorZones[2] {
				result += StyleGreen200.Render("█")
				break
			}
			// larger than zone 2, less or equal than zone 3
			if value > colorZones[2] && value <= colorZones[3] {
				result += StyleGreen300.Render("█")
				break
			}
			// larger than zone 3, less or equal than zone 4
			if value > colorZones[3] && value <= colorZones[4] {
				result += StyleGreen400.Render("█")
				break
			}
			// larger than zone 4, less or equal than zone 5
			if value > colorZones[4] && value <= colorZones[5] {
				result += StyleGreen500.Render("█")
				break
			}
			// larger than zone 5, less or equal than zone 6
			if value > colorZones[5] && value <= colorZones[6] {
				result += StyleGreen600.Render("█")
				break
			}
			// larger than zone 6, less or equal than zone 7
			if value > colorZones[6] && value <= colorZones[7] {
				result += StyleGreen700.Render("█")
				break
			}
			// larger than zone 7, less or equal than zone 8
			if value > colorZones[7] && value <= colorZones[8] {
				result += StyleGreen800.Render("█")
				break
			}
			// larger than zone 8, less or equal than zone 9
			if value > colorZones[8] && value <= colorZones[9] {
				result += StyleGreen900.Render("█")
				break
			}
			// larger than zone 9, less or equal than zone 10
			if value > colorZones[9] && value <= colorZones[10] {
				result += StyleGreen950.Render("█")
				break
			}
			// larger than zone 10
			if value > colorZones[10] {
				result += StyleGreen950.Render("█")
				break
			}
		}
	}

	return result
}

func MkViewBlue(slice [][]float32, index int) string {
	result := ""

	if len(slice) == 0 {
		return result
	}

	colorZones := GetColorZones(slice)

	for _, value := range slice[index] {
		for i := 0; i < 11; i++ {
			// less or equal than zone 0
			if value <= colorZones[0] {
				result += StyleBlue050.Render("█")
				break
			}
			// larger than zone 0, less or equal than zone 1
			if value > colorZones[0] && value <= colorZones[1] {
				result += StyleBlue100.Render("█")
				break
			}
			// larger than zone 1, less or equal than zone 2
			if value > colorZones[1] && value <= colorZones[2] {
				result += StyleBlue200.Render("█")
				break
			}
			// larger than zone 2, less or equal than zone 3
			if value > colorZones[2] && value <= colorZones[3] {
				result += StyleBlue300.Render("█")
				break
			}
			// larger than zone 3, less or equal than zone 4
			if value > colorZones[3] && value <= colorZones[4] {
				result += StyleBlue400.Render("█")
				break
			}
			// larger than zone 4, less or equal than zone 5
			if value > colorZones[4] && value <= colorZones[5] {
				result += StyleBlue500.Render("█")
				break
			}
			// larger than zone 5, less or equal than zone 6
			if value > colorZones[5] && value <= colorZones[6] {
				result += StyleBlue600.Render("█")
				break
			}
			// larger than zone 6, less or equal than zone 7
			if value > colorZones[6] && value <= colorZones[7] {
				result += StyleBlue700.Render("█")
				break
			}
			// larger than zone 7, less or equal than zone 8
			if value > colorZones[7] && value <= colorZones[8] {
				result += StyleBlue800.Render("█")
				break
			}
			// larger than zone 8, less or equal than zone 9
			if value > colorZones[8] && value <= colorZones[9] {
				result += StyleBlue900.Render("█")
				break
			}
			// larger than zone 9, less or equal than zone 10
			if value > colorZones[9] && value <= colorZones[10] {
				result += StyleBlue950.Render("█")
				break
			}
			// larger than zone 10
			if value > colorZones[10] {
				result += StyleBlue950.Render("█")
				break
			}
		}
	}

	return result
}
