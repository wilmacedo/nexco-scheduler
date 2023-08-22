package config

import "errors"

var ErrFailedToLoadEnv = errors.New("failed to load env file")
var ErrConvertInterval = errors.New("failed to convert interval to int")
