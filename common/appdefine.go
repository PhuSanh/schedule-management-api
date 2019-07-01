package common

const (
	//mesg
	ERROR_NO_ERORR                  = 0
	ERROR_GET_ROW_FROM_DB           = 1
	ERROR_GET_REQUEST_DATA          = 2
	ERROR_REQUEST_DATA_INVALID      = 3
	ERROR_INSERT_ROW_TO_DB          = 4
	ERROR_DELETE_ROW_IN_DB          = 5
	ERROR_UPDATE_ROW_IN_DB          = 6
	ERROR_REQUEST_PARAMETER_INVALID = 7

	MSG_SUCEESS                   = "success"
	MSG_GET_ROW_FROM_DB           = "get data from db error"
	MSG_INSERT_ROW_TO_DB          = "insert row to data error"
	MSG_DELETE_ROW_TO_DB          = "delete row in data error"
	MSG_UPDATE_ROW_TO_DB          = "update row in data error"
	MSG_GET_REQUEST_DATA          = "get request data error"
	MSG_REQUEST_DATA_INVALID      = "request data invalid"
	MSG_REQUEST_PARAMETER_INVALID = "invalid parameter"
	MSG_NOT_EXIST                 = "not exist"

	//paging

	PAGING_OFFSET = "offset"
	PAGING_LIMIT  = "limit"
)
