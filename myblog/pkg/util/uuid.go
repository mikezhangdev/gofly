package util

import "github.com/google/uuid"

func UUid()string{
	str := uuid.New().String()
	return str
}
