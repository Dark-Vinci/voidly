package utils

import "errors"

var ErrorAlreadyExist = errors.New("User already exists")
var HashingError = errors.New("error occured when hashing password")
var UnableToInsert = errors.New("error occured when inserting user")
var NotFound = errors.New("error occured when inserting user")
var InvalidCredentials = errors.New("error occured when inserting user")
var UnableToPerformOperation = errors.New("error occured when inserting user")

var ErrInvalidTokenHeaderFormat = errors.New("invalid token header format")
var ErrInvalidToken = errors.New("invalid token")
