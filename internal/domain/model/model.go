package model

// main models

type Chat struct {
	Id   ID     `json:"id"`
	Name string `json:"name"`
	//ApiKey      ApiKey `json:"api_key"`
	Port    int    `json:"port"`
	Version string `json:"version"`
}

type CommandName string

const (
	ChatLits   CommandName = "list"
	ChatCreate CommandName = "create"
)

type FlagKind string

const (
	FlagString FlagKind = "str"
	FlagInt    FlagKind = "int"
	FlagBool   FlagKind = "bool"
)

type Command struct {
	Name        CommandName
	Description string
	flags       []CommandFlag
}

type CommandFlag struct {
	Name     string
	Kind     FlagKind
	Required bool
}
