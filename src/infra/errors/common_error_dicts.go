package errors

const (

	// client error
	UNAUTHORIZED       ErrorCode = 1001 // Authorization errors
	DATA_INVALID       ErrorCode = 1002 // Validation errors
	USER_ALREADY_EXIST ErrorCode = 1003 // Unique Data Errors

	// server error
	UNKNOWN_ERROR         ErrorCode = 2000 // Unknown server errors
	FAILED_RETRIEVE_DATA  ErrorCode = 2001 // Data retrieval errors
	FAILED_FORWARD_DATA   ErrorCode = 2002 // Data forwarding errors
	FAILED_CREATE_DATA    ErrorCode = 2003 // Data creation errors
	STATUS_PAGE_NOT_FOUND ErrorCode = 2004 // Not found errors
)

var errorCodes = map[ErrorCode]*CommonError{
	UNKNOWN_ERROR: {
		ClientMessage: "Unknown error.",
		SystemMessage: "Unknown error.",
		ErrorCode:     UNKNOWN_ERROR,
	},
	DATA_INVALID: {
		ClientMessage: "Invalid Data Request",
		SystemMessage: "Some of query params has invalid value.",
		ErrorCode:     DATA_INVALID,
	},
	FAILED_RETRIEVE_DATA: {
		ClientMessage: "Failed to retrieve Data.",
		SystemMessage: "Something wrong happened while retrieve Data.",
		ErrorCode:     FAILED_RETRIEVE_DATA,
	},
	STATUS_PAGE_NOT_FOUND: {
		ClientMessage: "Invalid Status Page.",
		SystemMessage: "Status Page Email Address not found.",
		ErrorCode:     STATUS_PAGE_NOT_FOUND,
	},
	UNAUTHORIZED: {
		ClientMessage: "Unauthorized",
		SystemMessage: "Unauthorized",
		ErrorCode:     UNAUTHORIZED,
	},
	FAILED_FORWARD_DATA: {
		ClientMessage: "Failed to forward data.",
		SystemMessage: "Something wrong happened while forwarding data.",
		ErrorCode:     FAILED_FORWARD_DATA,
	},

	FAILED_CREATE_DATA: {
		ClientMessage: "Failed to create data.",
		SystemMessage: "Something wrong happened while create data.",
		ErrorCode:     FAILED_CREATE_DATA,
	},
	USER_ALREADY_EXIST: {
		ClientMessage: "Username Already Exist.",
		SystemMessage: "Username Already Exist.",
		ErrorCode:     USER_ALREADY_EXIST,
	},
}
