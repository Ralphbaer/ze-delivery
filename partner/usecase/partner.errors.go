package usecase

import "errors"

// ErrReasonIsEmpty is throwed when startDate or endDate (or both) are informed and reason not
var ErrReasonIsEmpty = errors.New("reason is empty")

// ErrOutOfSchedule is throwed when startDate or endDate (or both) are out of our Schedule (08AM - 08PM)
var ErrOutOfSchedule = errors.New("startDate and/or EndDate is/are out out of our Schedule (08AM - 08PM)")

// ErrInvalidReasonID is throwed the ReasonId informed is not valid
var ErrInvalidReasonID = errors.New("reasonId is not valid")
