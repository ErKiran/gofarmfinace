package codes

const (
	success         = 200
	internalError   = 500
	invalidParams   = 400
	noAuthorization = 401
	noPermission    = 403
	pageNotFound    = 404
)

const (
	Ok uint8 = iota
	IsRequiredField
	UnknownError
	InvalidFormData
	NotEnoughPermission
	InvalidURIParam
	NotFound
)

var httpCodeFlags = map[uint8]int{
	Ok:                  success,
	IsRequiredField:     invalidParams,
	UnknownError:        internalError,
	InvalidFormData:     invalidParams,
	NotEnoughPermission: noPermission,
	InvalidURIParam:     invalidParams,
	NotFound:            pageNotFound,
}

func GetHttpCode(msgCode uint8) int {
	if code, ok := httpCodeFlags[msgCode]; ok {
		return code
	}
	return internalError
}
