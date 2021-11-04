package handler

func ErrorData(kode string, name string, link string) map[string]string {
	errH := make(map[string]string)
	errH["kode"] = kode
	errH["name"] = name
	errH["link"] = link
	return errH
}
