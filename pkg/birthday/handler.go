package birthday

func BuildHandlers() map[string]func() string {
	return map[string]func() string{
		"all": AllHandler,
	}
}

func AllHandler() string {
	return ""
}
