package listnotes

import "test/pkg/error"

var errUnexpectedError = error.New("500", "", "ERR_INVALID_GET_NOTES", "failed to get list. unexpected error")
