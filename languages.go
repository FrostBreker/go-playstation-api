package playstation

type Language string

func SupportedLanguages() []Language {
	return Languages
}

var Languages = []Language{
	LangENUS,
	LangJA,
	LangEN,
	LangENGB,
	LangFR,
	LangES,
	LangESMX,
	LangDE,
	LangIT,
	LangNL,
	LangPT,
	LangPTBR,
	LangRU,
	LangPL,
	LangFI,
	LangDA,
	LangNO,
	LangSV,
	LangTR,
	LangKO,
	LangZHCN,
	LangZHTW,
}

var (
	LangJA   Language = "ja"
	LangEN   Language = "en"
	LangENGB Language = "en-GB"
	LangENUS Language = "en-US"
	LangFR   Language = "fr"
	LangES   Language = "es"
	LangESMX Language = "es-MX"
	LangDE   Language = "de"
	LangIT   Language = "it"
	LangNL   Language = "nl"
	LangPT   Language = "pt"
	LangPTBR Language = "pt-BR"
	LangRU   Language = "ru"
	LangPL   Language = "pl"
	LangFI   Language = "fi"
	LangDA   Language = "da"
	LangNO   Language = "no"
	LangSV   Language = "sv"
	LangTR   Language = "tr"
	LangKO   Language = "ko"
	LangZHCN Language = "zh-CN"
	LangZHTW Language = "zh-TW"
)
