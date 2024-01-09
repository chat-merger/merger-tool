package model

import "fmt"

// value models

type ID struct {
	value string
}

func (r ID) String() string {
	return fmt.Sprintf("%s", r.value)
}

func (r ID) Value() string {
	return r.value
}

func NewID(val string) ID {
	return ID{val}
}

type ApiKey struct {
	value string
}

func (r ApiKey) String() string {
	return fmt.Sprintf("%s", r.value)
}

func (r ApiKey) Value() string {
	return r.value
}

func NewApiKey(val string) ApiKey {
	return ApiKey{val}
}
