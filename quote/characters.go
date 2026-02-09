package quote

type Character string

const (
	GroupVoices Character = "00"
	Kinzo       Character = "01"
	Krauss      Character = "02"
	Natsuhi     Character = "03"
	Jessica     Character = "04"
	Eva         Character = "05"
	Hideyoshi   Character = "06"
	George      Character = "07"
	Rudolf      Character = "08"
	Kyrie       Character = "09"
	Battler     Character = "10"
	Ange        Character = "11"
	Rosa        Character = "12"
	Maria       Character = "13"
	Genji       Character = "14"
	Shannon     Character = "15"
	Kanon       Character = "16"
	Gohda       Character = "17"
	Kumasawa    Character = "18"
	Nanjo       Character = "19"
	Amakusa     Character = "20"
	Okonogi     Character = "21"
	Kasumi      Character = "22"
	Professor   Character = "23"
	Kawabata    Character = "24"
	NanjoSon    Character = "25"
	KumasawaSon Character = "26"
	Beatrice    Character = "27"
	Bernkastel  Character = "28"
	Lambdadelta Character = "29"
	Virgilia    Character = "30"
	Ronove      Character = "31"
	Gaap        Character = "32"
	Sakutaro    Character = "33"
	EvaBeatrice Character = "34"
	Chiester45  Character = "35"
	Chiester410 Character = "36"
	Lucifer     Character = "38"
	Leviathan   Character = "39"
	Satan       Character = "40"
	Belphegor   Character = "41"
	Mammon      Character = "42"
	Beelzebub   Character = "43"
	Asmodeus    Character = "44"
	Goat        Character = "45"
	Erika       Character = "46"
	Dlanor      Character = "47"
	Gertrude    Character = "48"
	Cornelia    Character = "49"
	Featherine  Character = "50"
	Zepar       Character = "51"
	Furfur      Character = "52"
	Lion        Character = "53"
	Will        Character = "54"
	Clair       Character = "55"
	KinzoYoung  Character = "58"
)

var characters = map[string]Character{
	"beatrice":             Beatrice,
	"beatrice2":            Beatrice,
	"beatrice3":            Beatrice,
	"virgilia":             Virgilia,
	"ronove":               Ronove,
	"will":                 Will,
	"lion":                 Lion,
	"krauss":               Krauss,
	"natsuhi":              Natsuhi,
	"rudolf":               Rudolf,
	"kyrie":                Kyrie,
	"hideyoshi":            Hideyoshi,
	"gohda":                Gohda,
	"nanjo":                Nanjo,
	"nanjoson":             NanjoSon,
	"kumasawa":             Kumasawa,
	"kumasawason":          KumasawaSon,
	"sakutaro":             Sakutaro,
	"luci":                 Lucifer,
	"satan":                Satan,
	"levi":                 Leviathan,
	"mammon":               Mammon,
	"clair vaux bernardus": Clair,
	"zepar":                Zepar,
	"kasumi":               Kasumi,
	"kawabata":             Kawabata,
	"kinzo2":               KinzoYoung,
	"okonogi":              Okonogi,
	"battler":              Battler,
	"bernkastel":           Bernkastel,
	"lambdadelta":          Lambdadelta,
	"erika":                Erika,
	"ange":                 Ange,
	"eva":                  Eva,
	"eva2":                 EvaBeatrice,
	"george":               George,
	"jessica":              Jessica,
	"maria":                Maria,
	"rosa":                 Rosa,
	"kinzo":                Kinzo,
	"shannon":              Shannon,
	"kanon":                Kanon,
	"genji":                Genji,
	"dlanor":               Dlanor,
	"featherine":           Featherine,
	"gapgapgap":            Gaap,
	"asmo":                 Asmodeus,
	"beelz":                Beelzebub,
	"belphe":               Belphegor,
	"cornelia":             Cornelia,
	"amakusa":              Amakusa,
	"furfur":               Furfur,
	"gertrude":             Gertrude,
	"goat-kun":             Goat,
	"lan":                  Lambdadelta,
	"professor":            Professor,
	"00":                   GroupVoices,
	"410":                  Chiester410,
	"45":                   Chiester45,
}

func LookupCharacter(name string) (Character, bool) {
	c, ok := characters[name]
	return c, ok
}
