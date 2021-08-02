package handler

import "errors"

// ErrPartnerAlreadyExists is throwed when a document (partner unique field) already exists in the repository
var ErrPartnerDocumentConflict = errors.New("partner document already taken")
