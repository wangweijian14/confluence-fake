package utility

func SplitByBatch(data []string, batch int) [][]string {
	var result [][]string
	length := len(data)
	i := 0
	for ; i < length; i = i + batch {
		end := i + batch
		if end > length {
			end = length
		}
		result = append(result, data[i:end])
	}
	return result
}
