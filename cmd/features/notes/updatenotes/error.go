package updatenotes

import "test/pkg/error"

var errUnexpectedError = error.New("500", "", "ERR_INVALID_GET_NOTES", "failed to update note. unexpected error")
var errNoteNotFoundError = error.New("404", "", "ERR_NOTES_NOT_FOUND", "failed to update note. data not found")
var errInvalidIDFormatError = error.New("400", "", "ERR_INVALID_ID_FORMAT", "failed to update note. invalid id format")
