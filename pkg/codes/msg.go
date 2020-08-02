package codes

var ErrMsgFlags = map[uint8]string{
	UnknownError:        "Something went wrong.",
	NotEnoughPermission: "You do not have sufficient permission to access the resource.",
	InvalidFormData:     "Failed to validate form data.",
	InvalidURIParam:     "Invalid uri param.",
}

func GetMsg(code uint8) (int, string) {
	msg := ErrMsgFlags[code]
	return GetHttpCode(code), msg
}
