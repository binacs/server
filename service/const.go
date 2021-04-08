package service

import "time"

const (
	dbCheckInterval  = 15 * time.Second
	cliCheckInterval = 1 * time.Minute

	fatalPscNilInGinContext = "fatal-psc-nil-in-gin-context"
)
