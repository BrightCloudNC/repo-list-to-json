package dataX

type DataX struct{
	CurrentName string `json:"currentName"`
	NewName string `json:"newName"`
}

func DataXList(data [][]string) []DataX{
	var dataXList []DataX
	for i, line := range data{
		var dataXRecord DataX
		if i == 0 {
			continue
		}
		for j, v := range line{
			switch j {
			case 0:
				dataXRecord.CurrentName = v
			case 2:
				dataXRecord.NewName = v
			}
		}
		dataXList = append(dataXList, dataXRecord)
	}
	return dataXList
}