package kmp

func SearchGreaterOccurrence(buffer, pattern []byte) (idxOccurrence, lenOccurrence int) {
	var table = Table{}
	var idxBuffer, lenOccPattern int
	var lenBuffer = len(buffer)
	var lenPattern = len(pattern)

	idxOccurrence = -1
	table.Build(pattern)
	for idxBuffer < lenBuffer {
		if buffer[idxBuffer] == pattern[lenOccPattern] {
			idxBuffer += 1
			lenOccPattern += 1

			if lenOccurrence < lenOccPattern {
				lenOccurrence = lenOccPattern
				idxOccurrence = idxBuffer - lenOccPattern

				if lenOccPattern == lenPattern {
					break
				}
			}
		} else {
			lenOccPattern = table.Content[lenOccPattern]

			if lenOccPattern < 0 {
				idxBuffer += 1
				lenOccPattern += 1
			}
		}
	}

	return idxOccurrence, lenOccurrence
}