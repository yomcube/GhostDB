package common

import "errors"

// https://wiki.tockdom.com/wiki/List_of_Identifiers#Courses
type CourseID byte

const (
	MarioCircuit        CourseID = 0x00
	MooMooMeadows       CourseID = 0x01
	MushroomGorge       CourseID = 0x02
	GrumbleVolcano      CourseID = 0x03
	ToadsFactory        CourseID = 0x04
	CoconutMall         CourseID = 0x05
	DKsSnowboardCross   CourseID = 0x06
	WariosGoldMine      CourseID = 0x07
	LuigiCircuit        CourseID = 0x08
	DaisyCircuit        CourseID = 0x09
	MoonviewHighway     CourseID = 0x0A
	MapleTreeway        CourseID = 0x0B
	BowsersCastle       CourseID = 0x0C
	RainbowRoad         CourseID = 0x0D
	DryDryRuins         CourseID = 0x0E
	KoopaCape           CourseID = 0x0F
	GCNPeachBeach       CourseID = 0x10
	GCNMarioCircuit     CourseID = 0x11
	GCNWaluigiStadium   CourseID = 0x12
	GCNDKMountain       CourseID = 0x13
	DSYoshiFalls        CourseID = 0x14
	DSDesertHills       CourseID = 0x15
	DSPeachGardens      CourseID = 0x16
	DSDelfinoSquare     CourseID = 0x17
	SNESMarioCircuit3   CourseID = 0x18
	SNESGhostValley2    CourseID = 0x19
	N64MarioRaceway     CourseID = 0x1A
	N64SherbetLand      CourseID = 0x1B
	N64BowsersCastle    CourseID = 0x1C
	N64DKsJungleParkway CourseID = 0x1D
	GBABowserCastle3    CourseID = 0x1E
	GBAShyGuyBeach      CourseID = 0x1F
)

func (course CourseID) ToString() (string, error) {
	switch course {
	case 0x00:
		return "Mario Circuit", nil
	case 0x01:
		return "Moo Moo Meadows", nil
	case 0x02:
		return "Mushroom Gorge", nil
	case 0x03:
		return "Grumble Volcano", nil
	case 0x04:
		return "Toad's Factory", nil
	case 0x05:
		return "Coconut Mall", nil
	case 0x06:
		return "DK's Snowboard Cross", nil
	case 0x07:
		return "Wario's Gold Mine", nil
	case 0x08:
		return "Luigi Circuit", nil
	case 0x09:
		return "Daisy Circuit", nil
	case 0x0A:
		return "Moonview Highway", nil
	case 0x0B:
		return "Maple Treeway", nil
	case 0x0C:
		return "Bowser's Castle", nil
	case 0x0D:
		return "Rainbow Road", nil
	case 0x0E:
		return "Dry Dry Ruins", nil
	case 0x0F:
		return "Koopa Cape", nil
	case 0x10:
		return "GCN Peach Beach", nil
	case 0x11:
		return "GCN Mario Circuit", nil
	case 0x12:
		return "GCN Waluigi Stadium", nil
	case 0x13:
		return "GCN DK Mountain", nil
	case 0x14:
		return "DS Yoshi Falls", nil
	case 0x15:
		return "DS Desert Hills", nil
	case 0x16:
		return "DS Peach Gardens", nil
	case 0x17:
		return "DS Delfino Square", nil
	case 0x18:
		return "SNES Mario Circuit 3", nil
	case 0x19:
		return "SNES Ghost Valley 2", nil
	case 0x1A:
		return "N64 Mario Raceway", nil
	case 0x1B:
		return "N64 Sherbet Land", nil
	case 0x1C:
		return "N64 Bowser's Castle", nil
	case 0x1D:
		return "N64 DK's Jungle Parkway", nil
	case 0x1E:
		return "GBA Bowser Castle 3", nil
	case 0x1F:
		return "GBA Shy Guy Beach", nil
	default:
		return "", errors.New("Course ID invalid")
	}
}

type CharacterID byte

const (
	Mario            CharacterID = 0x00
	BabyPeach        CharacterID = 0x01
	Waluigi          CharacterID = 0x02
	Bowser           CharacterID = 0x03
	BabyDaisy        CharacterID = 0x04
	DryBones         CharacterID = 0x05
	BabyMario        CharacterID = 0x06
	Luigi            CharacterID = 0x07
	Toad             CharacterID = 0x08
	DonkeyKong       CharacterID = 0x09
	Yoshi            CharacterID = 0x0A
	Wario            CharacterID = 0x0B
	BabyLuigi        CharacterID = 0x0C
	Toadette         CharacterID = 0x0D
	Koopa            CharacterID = 0x0E
	Daisy            CharacterID = 0x0F
	Peach            CharacterID = 0x10
	Birdo            CharacterID = 0x11
	DiddyKong        CharacterID = 0x12
	KingBoo          CharacterID = 0x13
	BowserJr         CharacterID = 0x14
	DryBowser        CharacterID = 0x15
	FunkyKong        CharacterID = 0x16
	Rosalina         CharacterID = 0x17
	MiiAMaleSmall    CharacterID = 0x18
	MiiAFemaleSmall  CharacterID = 0x19
	MiiBMaleSmall    CharacterID = 0x1A
	MiiBFemaleSmall  CharacterID = 0x1B
	MiiCMaleSmall    CharacterID = 0x1C
	MiiCFemaleSmall  CharacterID = 0x1D
	MiiAMaleMedium   CharacterID = 0x1E
	MiiAFemaleMedium CharacterID = 0x1F
	MiiBMaleMedium   CharacterID = 0x20
	MiiBFemaleMedium CharacterID = 0x21
	MiiCMaleMedium   CharacterID = 0x22
	MiiCFemaleMedium CharacterID = 0x23
	MiiAMaleHeavy    CharacterID = 0x24
	MiiAFemaleHeavy  CharacterID = 0x25
	MiiBMaleHeavy    CharacterID = 0x26
	MiiBFemaleHeavy  CharacterID = 0x27
	MiiCMaleHeavy    CharacterID = 0x28
	MiiCFemaleHeavy  CharacterID = 0x29
	MediumMii        CharacterID = 0x2A
	SmallMii         CharacterID = 0x2B
	LargeMii         CharacterID = 0x2C
	BikerPeach       CharacterID = 0x2D
	BikerDaisy       CharacterID = 0x2E
	BikerRosalina    CharacterID = 0x2F
)

func (character CharacterID) ToString() (string, error) {
	switch character {
	case 0x00:
		return "Mario", nil
	case 0x01:
		return "Baby Peach", nil
	case 0x02:
		return "Waluigi", nil
	case 0x03:
		return "Bowser", nil
	case 0x04:
		return "Baby Daisy", nil
	case 0x05:
		return "Dry Bones", nil
	case 0x06:
		return "Baby Mario", nil
	case 0x07:
		return "Luigi", nil
	case 0x08:
		return "Toad", nil
	case 0x09:
		return "Donkey Kong", nil
	case 0x0A:
		return "Yoshi", nil
	case 0x0B:
		return "Wario", nil
	case 0x0C:
		return "Baby Luigi", nil
	case 0x0D:
		return "Toadette", nil
	case 0x0E:
		return "Koopa Troopa", nil
	case 0x0F:
		return "Daisy", nil
	case 0x10:
		return "Peach", nil
	case 0x11:
		return "Birdo", nil
	case 0x12:
		return "Diddy Kong", nil
	case 0x13:
		return "King Boo", nil
	case 0x14:
		return "Bowser Jr.", nil
	case 0x15:
		return "Dry Bowser", nil
	case 0x16:
		return "Funky Kong", nil
	case 0x17:
		return "Rosalina", nil
	case 0x18:
		return "Small Mii Outfit A (Male)", nil
	case 0x19:
		return "Small Mii Outfit A (Female)", nil
	case 0x1A:
		return "Small Mii Outfit B (Male)", nil
	case 0x1B:
		return "Small Mii Outfit B (Female)", nil
	case 0x1C:
		return "Small Mii Outfit C (Male)", nil
	case 0x1D:
		return "Small Mii Outfit C (Female)", nil
	case 0x1E:
		return "Medium Mii Outfit A (Male)", nil
	case 0x1F:
		return "Medium Mii Outfit A (Female)", nil
	case 0x20:
		return "Medium Mii Outfit B (Male)", nil
	case 0x21:
		return "Medium Mii Outfit B (Female)", nil
	case 0x22:
		return "Medium Mii Outfit C (Male)", nil
	case 0x23:
		return "Medium Mii Outfit C (Female)", nil
	case 0x24:
		return "Large Mii Outfit A (Male)", nil
	case 0x25:
		return "Large Mii Outfit A (Female)", nil
	case 0x26:
		return "Large Mii Outfit B (Male)", nil
	case 0x27:
		return "Large Mii Outfit B (Female)", nil
	case 0x28:
		return "Large Mii Outfit C (Male)", nil
	case 0x29:
		return "Large Mii Outfit C (Female)", nil
	case 0x2A:
		return "Medium Mii", nil
	case 0x2B:
		return "Small Mii", nil
	case 0x2C:
		return "Large Mii", nil
	case 0x2D:
		return "Peach Biker Outfit", nil
	case 0x2E:
		return "Daisy Biker Outfit", nil
	case 0x2F:
		return "Rosalina Biker Outfit", nil
	default:
		return "", errors.New("Character ID invalid")
	}
}

type VehicleID byte

const (
	StandardKartS   VehicleID = 0
	StandardKartM   VehicleID = 1
	StandardKartL   VehicleID = 2
	BoosterSeat     VehicleID = 3
	ClassicDragster VehicleID = 4
	Offroader       VehicleID = 5
	MiniBeast       VehicleID = 6
	WildWing        VehicleID = 7
	FlameFlyer      VehicleID = 8
	CheepCharger    VehicleID = 9
	SuperBlooper    VehicleID = 10
	PiranhaProwler  VehicleID = 11
	TinyTitan       VehicleID = 12
	Daytripper      VehicleID = 13
	Jetsetter       VehicleID = 14
	BlueFalcon      VehicleID = 15
	Sprinter        VehicleID = 16
	Honeycoupe      VehicleID = 17
	StandardBikeS   VehicleID = 18
	StandardBikeM   VehicleID = 19
	StandardBikeL   VehicleID = 20
	BulletBike      VehicleID = 21
	MachBike        VehicleID = 22
	FlameRunner     VehicleID = 23
	BitBike         VehicleID = 24
	Sugarscoot      VehicleID = 25
	WarioBike       VehicleID = 26
	Quacker         VehicleID = 27
	ZipZip          VehicleID = 28
	ShootingStar    VehicleID = 29
	Magikruiser     VehicleID = 30
	Sneakster       VehicleID = 31
	Spear           VehicleID = 32
	JetBubble       VehicleID = 33
	DolphinDasher   VehicleID = 34
	Phantom         VehicleID = 35
)

func (vehicle VehicleID) ToString() (string, error) {
	switch vehicle {
	case 0x00:
		return "Standard Kart S", nil
	case 0x01:
		return "Standard Kart M", nil
	case 0x02:
		return "Standard Kart L", nil
	case 0x03:
		return "Booster Seat", nil
	case 0x04:
		return "Classic Dragster", nil
	case 0x05:
		return "Offroader", nil
	case 0x06:
		return "Mini Beast", nil
	case 0x07:
		return "Wild Wing", nil
	case 0x08:
		return "Flame Flyer", nil
	case 0x09:
		return "Cheep Charger", nil
	case 0x0A:
		return "Super Blooper", nil
	case 0x0B:
		return "Piranha Prowler", nil
	case 0x0C:
		return "Tiny Titan", nil
	case 0x0D:
		return "Daytripper", nil
	case 0x0E:
		return "Jetsetter", nil
	case 0x0F:
		return "Blue Falcon", nil
	case 0x10:
		return "Sprinter", nil
	case 0x11:
		return "Honeycoupe", nil
	case 0x12:
		return "Standard Bike S", nil
	case 0x13:
		return "Standard Bike M", nil
	case 0x14:
		return "Standard Bike L", nil
	case 0x15:
		return "Bullet Bike", nil
	case 0x16:
		return "Mach Bike", nil
	case 0x17:
		return "Flame Runner", nil
	case 0x18:
		return "Bit Bike", nil
	case 0x19:
		return "Sugarscoot", nil
	case 0x1A:
		return "Wario Bike", nil
	case 0x1B:
		return "Quacker", nil
	case 0x1C:
		return "Zip Zip", nil
	case 0x1D:
		return "Shooting Star", nil
	case 0x1E:
		return "Magikruiser", nil
	case 0x1F:
		return "Sneakster", nil
	case 0x20:
		return "Spear", nil
	case 0x21:
		return "Jet Bubble", nil
	case 0x22:
		return "Dolphin Dasher", nil
	case 0x23:
		return "Phantom", nil
	default:
		return "", errors.New("Vehicle ID invalid")
	}
}

// https://wiibrew.org/wiki/Country_Codes
// https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3
type CountryCode byte

const (
	JPN CountryCode = 1
	AIA CountryCode = 8
	ATG CountryCode = 9
	ARG CountryCode = 10
	ABW CountryCode = 11
	BHS CountryCode = 12
	BRB CountryCode = 13
	BLZ CountryCode = 14
	BOL CountryCode = 15
	NRA CountryCode = 16
	VGB CountryCode = 17
	CAN CountryCode = 18
	CYM CountryCode = 19
	CHL CountryCode = 20
	COL CountryCode = 21
	CRI CountryCode = 22
	DMA CountryCode = 23
	DOM CountryCode = 24
	ECU CountryCode = 25
	SLV CountryCode = 26
	GUF CountryCode = 27
	GRD CountryCode = 28
	GLP CountryCode = 29
	GTM CountryCode = 30
	GUY CountryCode = 31
	HTI CountryCode = 32
	HND CountryCode = 33
	JAM CountryCode = 34
	MTQ CountryCode = 35
	MEX CountryCode = 36
	MSR CountryCode = 37
	ANT CountryCode = 38
	NIC CountryCode = 39
	PAN CountryCode = 40
	PRY CountryCode = 41
	PER CountryCode = 42
	KNA CountryCode = 43
	LCA CountryCode = 44
	VCT CountryCode = 45
	SUR CountryCode = 46
	TTO CountryCode = 47
	TCA CountryCode = 48
	USA CountryCode = 49
	URY CountryCode = 50
	VIR CountryCode = 51
	VEN CountryCode = 52
	ALB CountryCode = 64
	AUS CountryCode = 65
	AUT CountryCode = 66
	BEL CountryCode = 67
	NIH CountryCode = 68
	BWA CountryCode = 69
	BGR CountryCode = 70
	HRV CountryCode = 71
	CYP CountryCode = 72
	CZE CountryCode = 73
	DNK CountryCode = 74
	EST CountryCode = 75
	FIN CountryCode = 76
	FRA CountryCode = 77
	DEU CountryCode = 78
	GRC CountryCode = 79
	HUN CountryCode = 80
	ISL CountryCode = 81
	IRL CountryCode = 82
	ITA CountryCode = 83
	LVA CountryCode = 84
	LSO CountryCode = 85
	LIE CountryCode = 86
	LTU CountryCode = 87
	LUX CountryCode = 88
	MKD CountryCode = 89
	MLT CountryCode = 90
	MNE CountryCode = 91
	MOZ CountryCode = 92
	NAM CountryCode = 93
	NLD CountryCode = 94
	NZL CountryCode = 95
	NOR CountryCode = 96
	POL CountryCode = 97
	PRT CountryCode = 98
	ROU CountryCode = 99
	RUS CountryCode = 100
	SRB CountryCode = 101
	SVK CountryCode = 102
	SVN CountryCode = 103
	ZAF CountryCode = 104
	ESP CountryCode = 105
	SWZ CountryCode = 106
	SWE CountryCode = 107
	CHE CountryCode = 108
	TUR CountryCode = 109
	GBR CountryCode = 110
	ZMB CountryCode = 111
	ZWE CountryCode = 112
	AZE CountryCode = 113
	MRT CountryCode = 114
	MLI CountryCode = 115
	NER CountryCode = 116
	TCD CountryCode = 117
	SDN CountryCode = 118
	ERI CountryCode = 119
	DJI CountryCode = 120
	SOM CountryCode = 121
	TWN CountryCode = 128
	KOR CountryCode = 136
	HKG CountryCode = 144
	MAC CountryCode = 145
	IDN CountryCode = 152
	SGP CountryCode = 153
	THA CountryCode = 154
	PHL CountryCode = 155
	MYS CountryCode = 156
	CHN CountryCode = 160
	ARE CountryCode = 168
	IND CountryCode = 169
	EGY CountryCode = 170
	OMN CountryCode = 171
	QAT CountryCode = 172
	KWT CountryCode = 173
	SAU CountryCode = 174
	SYR CountryCode = 175
	BHR CountryCode = 176
	JOR CountryCode = 177
	XXX CountryCode = 0xFF
)

func (countryCode CountryCode) ToString() (string, error) {
	switch countryCode {
	case 1:
		return "Japan", nil
	case 8:
		return "Anguilla", nil
	case 9:
		return "Antigua and Barbuda", nil
	case 10:
		return "Argentina", nil
	case 11:
		return "Aruba", nil
	case 12:
		return "Bahamas", nil
	case 13:
		return "Barbados", nil
	case 14:
		return "Belize", nil
	case 15:
		return "Bolivia", nil
	case 16:
		return "Brazil", nil
	case 17:
		return "British Virgin Islands", nil
	case 18:
		return "Canada", nil
	case 19:
		return "Cayman Islands", nil
	case 20:
		return "Chile", nil
	case 21:
		return "Colombia", nil
	case 22:
		return "Costa Rica", nil
	case 23:
		return "Dominica", nil
	case 24:
		return "Dominican Republic", nil
	case 25:
		return "Ecuador", nil
	case 26:
		return "El Salvador", nil
	case 27:
		return "French Guiana", nil
	case 28:
		return "Grenada", nil
	case 29:
		return "Guadeloupe", nil
	case 30:
		return "Guatemala", nil
	case 31:
		return "Guyana", nil
	case 32:
		return "Haiti", nil
	case 33:
		return "Honduras", nil
	case 34:
		return "Jamaica", nil
	case 35:
		return "Martinique", nil
	case 36:
		return "Mexico", nil
	case 37:
		return "Montserrat", nil
	case 38:
		return "Netherlands Antilles", nil
	case 39:
		return "Nicaragua", nil
	case 40:
		return "Panama", nil
	case 41:
		return "Paraguay", nil
	case 42:
		return "Peru", nil
	case 43:
		return "St. Kitts and Nevis", nil
	case 44:
		return "St. Lucia", nil
	case 45:
		return "St. Vincent and the Grenadines", nil
	case 46:
		return "Suriname", nil
	case 47:
		return "Trinidad and Tobago", nil
	case 48:
		return "Turks and Caicos Islands", nil
	case 49:
		return "United States", nil
	case 50:
		return "Uruguay", nil
	case 51:
		return "US Virgin Islands", nil
	case 52:
		return "Venezuela", nil
	case 64:
		return "Albania", nil
	case 65:
		return "Australia", nil
	case 66:
		return "Austria", nil
	case 67:
		return "Belgium", nil
	case 68:
		return "Bosnia and Herzegovina", nil
	case 69:
		return "Botswana", nil
	case 70:
		return "Bulgaria", nil
	case 71:
		return "Croatia", nil
	case 72:
		return "Cyprus", nil
	case 73:
		return "Czech Republic", nil
	case 74:
		return "Denmark", nil
	case 75:
		return "Estonia", nil
	case 76:
		return "Finland", nil
	case 77:
		return "France", nil
	case 78:
		return "Germany", nil
	case 79:
		return "Greece", nil
	case 80:
		return "Hungary", nil
	case 81:
		return "Iceland", nil
	case 82:
		return "Ireland", nil
	case 83:
		return "Italy", nil
	case 84:
		return "Latvia", nil
	case 85:
		return "Lesotho", nil
	case 86:
		return "Liechtenstein", nil
	case 87:
		return "Lithuania", nil
	case 88:
		return "Luxembourg", nil
	case 89:
		return "F.Y.R of Macedonia", nil
	case 90:
		return "Malta", nil
	case 91:
		return "Montenegro", nil
	case 92:
		return "Mozambique", nil
	case 93:
		return "Namibia", nil
	case 94:
		return "Netherlands", nil
	case 95:
		return "New Zealand", nil
	case 96:
		return "Norway", nil
	case 97:
		return "Poland", nil
	case 98:
		return "Portugal", nil
	case 99:
		return "Romania", nil
	case 100:
		return "Russia", nil
	case 101:
		return "Serbia", nil
	case 102:
		return "Slovakia", nil
	case 103:
		return "Slovenia", nil
	case 104:
		return "South Africa", nil
	case 105:
		return "Spain", nil
	case 106:
		return "Eswatini", nil
	case 107:
		return "Sweden", nil
	case 108:
		return "Switzerland", nil
	case 109:
		return "Türkiye", nil
	case 110:
		return "United Kingdom", nil
	case 111:
		return "Zambia", nil
	case 112:
		return "Zimbabwe", nil
	case 113:
		return "Azerbaijan", nil
	case 114:
		return "Mauritania", nil
	case 115:
		return "Mali", nil
	case 116:
		return "Niger", nil
	case 117:
		return "Chad", nil
	case 118:
		return "Sudan", nil
	case 119:
		return "Eritrea", nil
	case 120:
		return "Djibouti", nil
	case 121:
		return "Somalia", nil
	case 128:
		return "Taiwan", nil
	case 136:
		return "South Korea", nil
	case 144:
		return "Hong Kong", nil
	case 145:
		return "Macao", nil
	case 152:
		return "Indonesia", nil
	case 153:
		return "Singapore", nil
	case 154:
		return "Thailand", nil
	case 155:
		return "Philippines", nil
	case 156:
		return "Malaysia", nil
	case 160:
		return "China", nil
	case 168:
		return "United Arab Emirates", nil
	case 169:
		return "India", nil
	case 170:
		return "Egypt", nil
	case 171:
		return "Oman", nil
	case 172:
		return "Qatar", nil
	case 173:
		return "Kuwait", nil
	case 174:
		return "Saudi Arabia", nil
	case 175:
		return "Syria", nil
	case 176:
		return "Bahrain", nil
	case 177:
		return "Jordan", nil
	case 0xFF:
		return "Unknown Country", nil
	default:
		return "", errors.New("Country Code invalid")
	}
}

func (countryCode CountryCode) StatecodeValid(v uint8) bool {
	if v == 0xFF {
		return true
	}
	switch countryCode {
	case XXX:
		return v == 0xFF // I'm somewhat sure that 0xFF Country Code is necessarily 0xFF State Code
	case JPN:
		return v >= 2 && v <= 48
	case AIA, ABW, BHS, BRB, VGB, CYM, DMA, GUF, GRD, GLP, MTQ, MSR, ANT, LCA, VCT, TCA, VIR, AZE, MRT, MLI, NER, TCD, SDN, ERI, DJI, SOM, HKG, MAC, SGP, QAT, KWT, BHR, JOR:
		return v == 1
	case ATG, CRI, PRT, RUS, ARE:
		return v >= 2 && v <= 8
	case ARG:
		return v >= 2 && v <= 25
	case BLZ, CYP, FIN:
		return v >= 2 && v <= 7
	case BOL, HTI, AUT, BWA, ZAF, ZMB:
		return v >= 2 && v <= 10
	case NRA, EGY:
		return v >= 2 && v <= 28
	case CAN, CHL, GRC, NAM, NZL, SAU:
		return v >= 2 && v <= 14
	case COL, IDN:
		return v >= 2 && v <= 34
	case DOM:
		return v >= 2 && v <= 31
	case ECU, GTM:
		return v >= 2 && v <= 23
	case SLV, JAM, KNA, CZE, SYR:
		return v >= 2 && v <= 15
	case GUY, PAN, SUR, LSO, LTU, ZWE:
		return v >= 2 && v <= 11
	case HND, PRY, PHL:
		return v >= 2 && v <= 19
	case MEX, CHN:
		return v >= 2 && v <= 33
	case NIC, DNK, ESP:
		return v >= 2 && v <= 18
	case PER, VEN, TWN:
		return v >= 2 && v <= 26
	case TTO, ALB, NLD:
		return v >= 2 && v <= 13
	case USA:
		return v >= 2 && v <= 53
	case URY:
		return v >= 2 && v <= 20
	case AUS, ISL, IRL, MKD, SVK, OMN:
		return v >= 2 && v <= 9
	case BEL, NIH, LUX, SRB:
		return v >= 2 && v <= 4
	case BGR:
		return v >= 2 && v <= 29
	case HRV, SWZ:
		return v >= 2 && v <= 5
	case EST:
		return v >= 2 && v <= 16
	case FRA, LVA:
		return v >= 2 && v <= 27
	case DEU, POL, KOR, MYS:
		return v >= 2 && v <= 17
	case HUN, ITA:
		return v >= 2 && v <= 21
	case LIE, MLT:
		return v >= 2 && v <= 3
	case MNE, SWE:
		return v >= 2 && v <= 22
	case MOZ:
		return v >= 2 && v <= 12
	case NOR, SVN, GBR:
		return v >= 2 && v <= 6
	case ROU:
		return v >= 2 && v <= 43
	case CHE:
		return v >= 2 && v <= 24
	case TUR:
		return v >= 2 && v <= 50
	case THA:
		return v >= 2 && v <= 77
	case IND:
		return v >= 2 && v <= 37
	default:
		return false
	}
}

type ControllerID byte

const (
	WiiWheel          ControllerID = 0
	Nunchuck          ControllerID = 1
	ClassicController ControllerID = 2
	GameCube          ControllerID = 3
)

func (controller ControllerID) ToString() (string, error) {
	switch controller {
	case 0x00:
		return "Wii Wheel", nil
	case 0x01:
		return "Wii Remote + Nunchuk", nil
	case 0x02:
		return "Classic Controller", nil
	case 0x03:
		return "GameCube Controller", nil
	default:
		return "", errors.New("Vehicle ID invalid")
	}
}
