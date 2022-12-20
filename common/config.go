package common

var version = "1.8.2"
var Userdict = map[string][]string{
	"ftp":        {"ftp", "admin", "www", "web", "root", "db", "wwwroot", "data"},
	"mysql":      {"root", "mysql"},
	"mssql":      {"sa", "sql"},
	"smb":        {"administrator", "admin", "guest"},
	"rdp":        {"administrator", "admin", "guest"},
	"postgresql": {"postgres", "admin"},
	"ssh":        {"root", "admin"},
	"mongodb":    {"root", "admin"},
	"oracle":     {"sys", "system", "admin", "test", "web", "orcl"},
}

var Passwords = []string{
	"123456",
	"123456789",
	"12345678",
	"654321",
	"1234567890",
	"woaini",
	"password",
	"zxcvbnm",
	"147258369",
	"147258",
	"987654321",
	"1qaz2wsx",
	"qazwsx",
	"qwqwqw",
	"123123",
	"tangkai",
	"qwertyuiop",
	"super123",
	"asd123",
	"abcd1234",
	"wodima123",
	"qazwsxedc",
	"741852963",
	"dedewang",
	"asdfasdf",
	"963852741",
	"qweasd",
	"qwe123456",
	"qweasdzxc",
	"zxcvbn",
	"qwer1234",
	"woaiwojia",
	"qaz123",
	"qyahzn",
	"qweqwe",
	"tianya",
	"zxcvbnm123",
	"147852369",
	"147369",
	"xiaoxiao",
	"woshishui",
	"asdf1234",
	"caonima",
	"abcdefg",
	"woaini123",
	"163.com",
	"buzhidao",
	"yangyang",
	"zhang123",
	"wang123",
	"qwertyui",
	"jingjing",
	"hao123456",
	"zinch",
	"3344521",
	"12345",
	"sier66I9jR",
	"123456123456",
	"loveyou",
	"888666",
	"585858",
	"111111a",
	"19491001",
	"369258147",
	"zxczxc",
	"852456",
	"lz110110",
	"qq123123",
	"9876543210",
	"521314",
	"tiantian",
	"9erXdo44zY",
	"ssssss",
	"11112222",
	"120120",
	"C2j5jbvr2F",
	"nihaoma",
	"12345678910",
	"abcdef",
	"qazwsx123",
	"aptx4869",
	"wocaonima",
	"222333",
	"ww6082516",
	"774517397",
	"qwaszx",
	"520123",
	"147896325",
	"weiwei",
	"aaa111",
	"jjcG16dj5K",
	"a000000",
	"q1w2e3",
	"555888",
	"a00000",
	"999888",
	"maomao",
	"123456qwe",
	"Khgdajyw2a",
	"112421",
	"686868",
	"159951",
	"shanghai",
	"kawiYU78",
	"sunshine",
	"787878",
	"121121",
	"ms0083jxj",
	"7788521",
	"woailaopo",
	"12qwaszx",
	"qazxsw",
	"majiajun",
	"goodday85S",
	"deeven",
	"234567",
	"xxxxxx",
	"565656",
	"212121",
	"159159",
	"1233211234567",
	"sgdHhfC4x2",
	"a5201314",
	"5203344",
	"qaz123456",
	"123456aaa",
	"wodemima",
	"qqwwee",
	"223344",
	"123456654321",
	"55555555",
	"235689",
	"goodluck",
	"332211",
	"b85Re4oloV",
	"98765",
	"12345679",
	"321456",
	"aini1314",
	"woaini521",
	"521521521",
	"1234566",
	"aaaaaaa",
	"l123456",
	"mnbvcxz",
	"kb9zc8uxtx",
	"5201314520",
	"86rzoGzb8V",
	"NBvBB32fa9",
	"123567",
	"0",
	"584131420",
	"zxzxzx",
	"123456711",
	"ApjSqpM844",
	"181818",
	"147147",
	"zxc123456",
	"superman",
	"123520",
	"a1b2c3d4",
	"11235813",
	"556677",
	"898989",
	"112112",
	"wojiushiwo",
	"feifei",
	"1a2b3c4d",
	"5841314520",
	"33333333",
	"122333",
	"aaa123456",
	"woainima",
	"168888",
	"999",
	"beyond",
	"a123456a",
	"635241",
	"998877",
	"uifKjhF522",
	"wangwei",
	"13141314",
	"650829yjm",
	"computer",
	"12345600",
	"JxsGx2Yd87",
	"100200300",
	"d123456",
	"ndaCebx2wx",
	"12345qwert",
	"a11111",
	"142857",
	"258456",
	"goyj2010",
	"hao123",
	"14789632",
	"jiajia",
	"7894561230",
	"zhangwei",
	"qwerasdf",
	"admin",
	"admin1234",
	"admin12345",
	"",
	"admin123",
	"root",
	"{user}12345",
	"{user}1234",
	"pass123",
	"pass@123",
	"111111",
	"123",
	"1",
	"admin@123",
	"admin123!@#",
	"{user}",
	"{user}1",
	"{user}111",
	"{user}123",
	"{user}@123",
	"{user}_123",
	"{user}#123",
	"{user}@111",
	"{user}@2019",
	"{user}@123#4",
	"P@ssw0rd!",
	"P@ssw0rd",
	"Passw0rd",
	"qwe123",
	"test",
	"test123",
	"123qwe",
	"123qwe!@#",
	"123321",
	"666666",
	"a123456.",
	"123456~a",
	"123456!a",
	"000000",
	"00000",
	"8888888",
	"!QAZ2wsx",
	"abc123",
	"abc123456",
	"1qaz@WSX",
	"a12345",
	"Aa1234",
	"Aa1234.",
	"Aa12345",
	"a123456",
	"a123123",
	"Aa123123",
	"Aa123456",
	"Aa12345.",
	"sysadmin",
	"system",
	"1qaz!QAZ",
	"2wsx@WSX",
	"qwe123!@#",
	"Aa123456!",
	"A123456s!",
	"sa123456",
	"1q2w3e",
	"Charge123",
	"Aa123456789",
}
var PORTList = map[string]int{
	"ftp":         21,
	"ssh":         22,
	"findnet":     135,
	"netbios":     139,
	"smb":         445,
	"mssql":       1433,
	"oracle":      1521,
	"mysql":       3306,
	"rdp":         3389,
	"psql":        5432,
	"redis":       6379,
	"fcgi":        9000,
	"mem":         11211,
	"mgo":         27017,
	"ms17010":     1000001,
	"cve20200796": 1000002,
	"web":         1000003,
	"webonly":     1000003,
	"webpoc":      1000003,
	"smb2":        1000004,
	"wmiexec":     1000005,
	"all":         0,
	"portscan":    0,
	"icmp":        0,
	"main":        0,
}

var Outputfile = "result.txt"
var SucOutputfile = "sucresult.txt"
var IsSave = true
var Webport = "80,81,82,83,84,85,86,87,88,89,90,91,92,98,99,443,800,801,808,880,888,889,1000,1010,1080,1081,1082,1099,1118,1888,2008,2020,2100,2375,2379,3000,3008,3128,3505,5555,6080,6648,6868,7000,7001,7002,7003,7004,7005,7007,7008,7070,7071,7074,7078,7080,7088,7200,7680,7687,7688,7777,7890,8000,8001,8002,8003,8004,8006,8008,8009,8010,8011,8012,8016,8018,8020,8028,8030,8038,8042,8044,8046,8048,8053,8060,8069,8070,8080,8081,8082,8083,8084,8085,8086,8087,8088,8089,8090,8091,8092,8093,8094,8095,8096,8097,8098,8099,8100,8101,8108,8118,8161,8172,8180,8181,8200,8222,8244,8258,8280,8288,8300,8360,8443,8448,8484,8800,8834,8838,8848,8858,8868,8879,8880,8881,8888,8899,8983,8989,9000,9001,9002,9008,9010,9043,9060,9080,9081,9082,9083,9084,9085,9086,9087,9088,9089,9090,9091,9092,9093,9094,9095,9096,9097,9098,9099,9100,9200,9443,9448,9800,9981,9986,9988,9998,9999,10000,10001,10002,10004,10008,10010,10250,12018,12443,14000,16080,18000,18001,18002,18004,18008,18080,18082,18088,18090,18098,19001,20000,20720,21000,21501,21502,28018,20880"
var DefaultPorts = "21,22,80,81,135,139,443,445,1433,1521,3306,5432,6379,7001,8000,8080,8089,9000,9200,11211,27017"

type HostInfo struct {
	Host    string
	Ports   string
	Url     string
	Infostr []string
}

type PocInfo struct {
	Target  string
	PocName string
}

var (
	Path        string
	Scantype    string
	Command     string
	SshKey      string
	Domain      string
	Username    string
	Password    string
	Proxy       string
	Timeout     int64 = 3
	WebTimeout  int64 = 5
	TmpSave     bool
	NoPing      bool
	Ping        bool
	Pocinfo     PocInfo
	IsWebCan    bool
	IsBrute     bool
	RedisFile   string
	RedisShell  string
	Userfile    string
	Passfile    string
	HostFile    string
	PortFile    string
	PocPath     string
	Threads     int
	URL         string
	UrlFile     string
	Urls        []string
	NoPorts     string
	NoHosts     string
	SC          string
	PortAdd     string
	UserAdd     string
	PassAdd     string
	BruteThread int
	LiveTop     int
	Socks5Proxy string
	Hash        string
	HashBytes   []byte
	HostPort    []string
	IsWmi       bool
)

var (
	UserAgent  = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36"
	Accept     = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	DnsLog     bool
	PocNum     int
	PocFull    bool
	CeyeDomain string
	ApiKey     string
	Cookie     string
)
