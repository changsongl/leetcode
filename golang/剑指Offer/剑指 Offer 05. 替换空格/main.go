package main

func replaceSpace(s string) string {
	ns := make([]byte, 0, len(s)*3)
	for _, b := range s {
		if b == ' ' {
			ns = append(ns, '%', '2', '0')
		} else {
			ns = append(ns, byte(b))
		}
	}
	return string(ns)
}
