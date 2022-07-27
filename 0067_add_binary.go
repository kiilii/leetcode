package main

// func main() {
// 	addBinary("hello", "world")
// }

func addBinary(a string, b string) string {
	var lena = len(a) - 1
	var lenb = len(b) - 1
	var min int
	var str []byte
	if lena < lenb {
		min = lena
	} else {
		min = lenb
	}

	for i := 0; i <= (min); i++ {
		if lena <= i && lenb <= i {
			if a[i] != b[i] {
				str = append(str, byte(a[lena]))
			} else {
				if a[i] > b[i] {
					str = append(str, byte(a[i]))
				}
			}
		}
	}
	return string(str)
}
