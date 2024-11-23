package playstation

type Region string

func SupportedRegions() []Region {
	return Regions
}

var (
	RegionUS   Region = "us"
	RegionCA   Region = "ca"
	RegionMX   Region = "mx"
	RegionCL   Region = "cl"
	RegionPE   Region = "pe"
	RegionAR   Region = "ar"
	RegionCO   Region = "co"
	RegionBR   Region = "br"
	RegionGB   Region = "gb"
	RegionIE   Region = "ie"
	RegionBE   Region = "be"
	RegionLU   Region = "lu"
	RegionNL   Region = "nl"
	RegionFR   Region = "fr"
	RegionDE   Region = "de"
	RegionAT   Region = "at"
	RegionCH   Region = "ch"
	RegionIT   Region = "it"
	RegionPT   Region = "pt"
	RegionDK   Region = "dk"
	RegionFI   Region = "fi"
	RegionNO   Region = "no"
	RegionSE   Region = "se"
	RegionAU   Region = "au"
	RegionNZ   Region = "nz"
	RegionES   Region = "es"
	RegionRU   Region = "ru"
	RegionAE   Region = "ae"
	RegionZA   Region = "za"
	RegionPL   Region = "pl"
	RegionGR   Region = "gr"
	RegionSA   Region = "sa"
	RegionCZ   Region = "cz"
	RegionBG   Region = "bg"
	RegionHR   Region = "hr"
	RegionRO   Region = "ro"
	RegionSI   Region = "si"
	RegionHU   Region = "hu"
	RegionSK   Region = "sk"
	RegionTR   Region = "tr"
	RegionBH   Region = "bh"
	RegionKW   Region = "kw"
	RegionLB   Region = "lb"
	RegionOM   Region = "om"
	RegionQA   Region = "qa"
	RegionIL   Region = "il"
	RegionMT   Region = "mt"
	RegionIS   Region = "is"
	RegionCY   Region = "cy"
	RegionIN   Region = "in"
	RegionUA   Region = "ua"
	RegionHK   Region = "hk"
	RegionTW   Region = "tw"
	RegionSG   Region = "sg"
	RegionMY   Region = "my"
	RegionID   Region = "id"
	RegionTH   Region = "th"
	RegionJP   Region = "jp"
	RegionKR   Region = "kr"
	RegionENUS Region = "en-US"
)

// known Sony regions
var Regions = []Region{
	RegionUS,
	RegionCA,
	RegionMX,
	RegionCL,
	RegionPE,
	RegionAR,
	RegionCO,
	RegionBR,
	RegionGB,
	RegionIE,
	RegionBE,
	RegionLU,
	RegionNL,
	RegionFR,
	RegionDE,
	RegionAT,
	RegionCH,
	RegionIT,
	RegionPT,
	RegionDK,
	RegionFI,
	RegionNO,
	RegionSE,
	RegionAU,
	RegionNZ,
	RegionES,
	RegionRU,
	RegionAE,
	RegionZA,
	RegionPL,
	RegionGR,
	RegionSA,
	RegionCZ,
	RegionBG,
	RegionHR,
	RegionRO,
	RegionSI,
	RegionHU,
	RegionSK,
	RegionTR,
	RegionBH,
	RegionKW,
	RegionLB,
	RegionOM,
	RegionQA,
	RegionIL,
	RegionMT,
	RegionIS,
	RegionCY,
	RegionIN,
	RegionUA,
	RegionHK,
	RegionTW,
	RegionSG,
	RegionMY,
	RegionID,
	RegionTH,
	RegionJP,
	RegionKR,
	RegionENUS,
}
