package localization

// TODO: Maybe do interface and have standardized ways of rendering the different
// data types as string, but leaving it as string does simplify everything significantly
type LocalizedText struct {
	Data    map[string]string
	Message string
}

type Locale map[string]LocalizedText

func InitLocale(langauge string) Locale {
	return Locale{
		"key": LocalizedText{
			Data: map[string]string{
				"VarName":  "5",
				"OtherVar": "mega",
			},
			Message: "This {{.VarName}} is that {{.OtherVar}} thing",
		},
	}
}
