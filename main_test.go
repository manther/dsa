package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseSttring(t *testing.T) {
	// some test code
	require.Equal(t, "dlrow olleh", reverseString("hello world"))
}

func TestReverseSttring2(t *testing.T) {
	// some test code
	require.Equal(t, "dlrow olleh", reverseString2("hello world"))
}

func TestReverseSttring3(t *testing.T) {
	// some test code
	require.Equal(t, []byte("dlrow olleh"), reverseString3([]byte("hello world")))
}

func TestReverseString4(t *testing.T) {
	// some test code
	s := reverseString3([]byte("hello world"))
	reverseString4(s)
	require.Equal(t, []byte("dlrow olleh"), s)
}

func TestReverseString5(t *testing.T) {
	// some test code
	s := reverseString3([]byte{104, 101, 108, 108, 111})
	reverseString4(s)
	require.Equal(t, []byte("olleh"), s)
}

func TestPalandrone(t *testing.T) {
	// some test code
	b := palandrome("hello")
	require.Equal(t, false, b)
	b = palandrome("racecar")
	require.Equal(t, true, b)
}

func TestPalandrone2(t *testing.T) {
	// some test code
	b := palandrome2("hello")
	require.Equal(t, false, b)
	b = palandrome2("racecar")
	require.Equal(t, true, b)
	b = palandrome2("abba")
	require.Equal(t, true, b)
	b = palandrome2("race car")
	require.Equal(t, true, b)
	b = palandrome2("A man, a plan, a canal: Panama")
	require.Equal(t, true, b)
}

func TestPalandroneLLint(t *testing.T) {
	// some test code
	b := palandromeLLInt(&ListNode{Val: 'h', Next: &ListNode{Val: 'e', Next: &ListNode{Val: 'l', Next: &ListNode{Val: 'l', Next: &ListNode{Val: 'o'}}}}})
	require.Equal(t, false, b)
	b = palandromeLLInt(&ListNode{Val: 'r', Next: &ListNode{Val: 'a', Next: &ListNode{Val: 'c', Next: &ListNode{Val: 'e', Next: &ListNode{Val: 'c', Next: &ListNode{Val: 'a', Next: &ListNode{Val: 'r'}}}}}}})
	require.Equal(t, true, b)
}

func TestReverseLL(t *testing.T) {
	// some test code
	l := reverseList(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}})
	require.Equal(t, &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}, l)
	l = reverseList(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}})
	require.Equal(t, &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}, l)
	require.Nil(t, reverseList(nil))
}

func TestCountWords(t *testing.T) {
	// some test code
	words1 := []string{"apple", "banana", "cherry"}
	words2 := []string{"date", "elderberry", "fig"}
	require.Equal(t, 0, countWords(words1, words2))
	words1 = []string{"plsfzfgawptsrhtppzhlcjqbcmi", "gsccupsxkshslvrixubnxdhcsjshp", "s", "ezrflcmchmhhvojohew", "lmskv", "fxhnharndihttfrnccktwbmxzpqee", "hd", "mcrcjipns", "e", "qrnqz", "csbtrpsmhdnafekmns", "wjb", "xmibfbabpomcgyhb", "hruigjocaexurmfljszxvlxkyolqky", "kiigmajzxjocnixufcy", "ckklctinivzecvhocjbsxymrt", "kxfzrrv", "qjceabulyrwatm", "ojafmsymfkotalkvsufnprdu", "wmkpuftthlvycyqmzozgvb", "umhjoqoenipkhumbikswzheqe", "hgyvfnvjosgxefphhzr", "rqxhyspbloycsxovno", "paicai", "zpnmsfkdljsscaddgwhirtqntv", "fysxtnbcohdlcpibrddtoa", "icwlkeyaiuzasuntg", "wmloodipkkplkbsico", "tbffw", "whlygbzsrewchnbtt", "gip", "psqxgfhpioqyw", "uwbnogpqsu", "ioibrgngksybilgr", "zojhxjnxqfojsfafkuvsiyimewsepo", "vfjvmojjizmvfzbc", "wkjnunyavkdomqcsrowiw", "nbbuavuhhomg", "qsjvrgkypqsxioefxkwh", "boejyfmqyvsbpugbqyrkpl", "efwlbrnpkfc", "huiaolundrnxezqqzyiqauboqm", "jmkjyizpmdumpw", "oykzgyoocwavrjum", "f", "dqurh", "cwguhiozmaeozcqwnffgexmfnku", "jxvlggrhculruzgfkokpmgrbyftdo", "qvrtqvchslzbykfwuxayywnracnlq", "daelyspayiondzyq", "eqghkkijnripfzksjrwulmxbaa", "zmyvzmwwhbgyfcqbbt", "mjucbpastaer", "phyagrxzqrkn", "nnavmxqyfczdd", "irqxbfyarmjeoiomprfjequksm", "ggdht", "hofvrhohgxvevxgjtgfprsmooo", "ntztieup", "ueflxsfystwopaopgldaimuipsaful", "iydkzacinibnpjgljls", "ihsropbisotqwhpndcqq", "lahooipcdsetztoncvjauzcnqh", "sobdrebzxpwxgvqveayc", "wcnddeqeqtdrerbbdwnpquhaxbbpe", "ithgkrvauylxnecjzkbroiaptu", "duceyy", "fhcyatyjbcmbftigjihvpceit", "ygtswpbjneglmhdgwtgrv", "g", "nzogophfgwdotbttenoxy", "lju", "kfb", "sgktjimgogrwvigpfunvfnxc", "rpavqupjhpuj", "vs", "hygvurijfrsrazwvxp", "rablosnrbxhf", "njiswypyccwn", "btvymedpiunurhaqtoxzxaxnorrf", "evxjylrnnbwlocmduito", "zjdduajyynlsuyymugcuwg", "qzesdxpfsj", "a", "iyouogtkdygmmqwmzsjdeppyoskh", "mwsvjmtbrfnm", "kz", "cvvqjhlrhzzmmntmbjyza", "wwaazasyy", "tyzeebwhqmcobmakafgebbwtigbi", "ginky", "dtalssxkpyfqaurgz", "yafm", "qhxehrmsljaufkvhvuwvmnunjno", "y", "jywujvcljmoqghtz", "bwjjtjrtcwlv", "iclszrajbgqvmsuh", "hdwofeucbvxxrsfaiikghvjcahwvy", "hjruyhqev", "drcgjidxzlucrucxdxd", "impxzwjxmiosor", "mkgdfgguom", "mmlbqulasotujisgwpkzoejsc", "lpckqobfpjzvy", "eeabxbb", "gmrjtyfkwgzaiabkirptgnfutztd", "hechoogbsykpoplww", "tyjlifgtizyvpqisgipjk", "mzzudpnujooifzqgatprzfejhmi", "cudoqctbacrxrewlrscddzgqvnyb", "hzhpnymwjhuaoflycsfopmvqcxxv", "pjgs", "ltnsokdyld", "nwzqpqymrxtmbdf", "nxqcezojdgswdx", "iimkiwkpnvuufvfgpmrpydwem", "lzzzxnztaigeuexwjkuve", "vzvmnh", "hqtkmfaguqowsmoqymilhuyqli", "jkj", "xmynmjfuvfjzwbyzllzuepkftsk", "adchcrgqoz", "iekocgvlolololo", "ozgzvbokbauwkipdeyfunptjco", "sdvtnigwgleiral", "ngvwegcsfwlzedmwyzghygj", "acbiywsbpifabqrtasbsiypwtm", "kextotfhtjl", "bgbeuwgfmru", "sprleyhjucfcli", "jchcophgucybaaexrbtfnfmcats", "cqosqjuxlzmseaweieeaposqpaa", "oalrhiwqccgsphcgwjyx", "uhakpqbmflimcknywbzcrwukzmkum", "iwaassehsaojeheeld", "brly", "kq", "wzdvbyfyx", "gzjngsmlxhu", "ofrupfvujzjexteznqu", "tzhwyzpyi", "cnjjpjz", "sfwedbvjwuxkhqzfglqhutznqvx", "xnuopczllpxqga", "gvzapdgeqezouspqbohbgsvodjdid", "aqyfe", "l", "uqiwis", "s", "jxgkgjgzxaunyledp", "bfwivzblomvixmxhrdxomldst", "ejqny", "rjizxyahdjoqnxvalzzx", "narkxounvlpmovyunkl", "uxadyjuoyhvybacvaffsfhspl", "dpjitixb", "mkqiiunsffxpvlysnqpbkrutsoo", "lnrrvbxxkqxdivr", "bhqkobemsduydnsea", "emyejsthfl", "cchkfx", "armbwbclk", "w", "wlzdkkmrmnhedtmipy", "pnymgukydtluubrwkngaohqp", "xuirrfqfyoimwm", "xllhirakkfdajyglr", "gcempkdx", "ahqgqykzitbqphjrorfsdknalmf", "krdbk", "ornevzpgknjx", "vnnvictiqsxwbujcoilmgmywh", "lyohgxylwgrlgsdhagtfgh", "swjetqfxmhbbsplionglesgwtobjig", "yvjyxcpfqzaiwgqw", "uanxy", "chggkz", "chnpnzeizdnliczrmvssguap", "zrrttunymnppvlj", "nrpdmrhwmhzeukjwiiersmlhohimmv", "lzhzr", "ywkymfcgittyxujrjdyvxoy", "agtpspxhilmmlspz", "orocovqkhxrafbjogqfizqe", "apdfsxnjmmqpymeiejrf", "gomvicmc", "et", "ccq", "yxqrmxtrbmrb", "ufniqsmdruabolytejvcgfzw", "vilollcowereynhexxslxggbdqdiv", "ewhszmrpvwkfvhub", "ropevek", "oycxkv", "qsdppfiwxlqahncomkhla", "dtnaflx", "vfseztyahalmrywwvxworxown", "chgbawrqo", "rvfwrqqgmredhitfuieplodgkbqke", "yrj", "tzuplosuzflzjbwl", "volgapwsdopoakj", "ymsllv", "exwegnmdksbt", "vkddjfkssoxfczokjroijjxvjxfz", "fuzpgwk", "byglewgcvgmboxjarhurztb", "hdyijxwbyymq", "bdow", "fijvnajwmxpfvxhfmdvgvzal", "bpxhuqjzfuxijqtpbczpmpo", "mmjt", "tvrwttkamwhwwwe", "eglznvgvwobujenmfhmbvutwydouy", "alvmrqlkvn", "grbhrsyxraezwvzjjlpyqvovtinb", "aijr", "xv", "wiefix", "ifvnxukkdkoal", "nifizflkpvhpdefbqeg", "ksnz", "yemju", "b", "beezbouznuehxuoxrwcxppzzfdii", "bmgmuaplw", "yorkpqtf", "gx", "pnjujio", "nuablp", "ehcwkkzztgavpuczgmbylahctkmahj", "kipvctrf", "omuowicgxorgvjposylpf", "sqpvrxaf", "fxahaigkuljflmymfdvxydncsc", "koz", "yhunffzpa", "lxsdrtfimvmp", "wbboihhzzsyngjw", "pxlbtcjcyeyltgjqmiaplldoih", "tvfgucmqsqawkoulnw", "zgikecpxgkvtyttkblqpgnypemhc", "kpxiescejxnzjbzmhjpadoygpufudg", "pzzakdgffmupv", "jrrlutbebsxzmxwlhykulprzv", "zcerueeudguhodsq", "pddqgygxxd", "fhtgwy", "lihwn", "kvbhujftweg", "mmj", "yuxrpaachruyjki", "nuaowzzownvzofuoovcbl", "jijoiywepebahhqdv", "bhdkwiqchmrqjbsdidxct", "onh", "snelkpzexsfub", "rhjednmlueduletulbfqywwjylyna", "lpbfygesivkzwveukcm", "epyrzjqyafhxddrxwgsowmsmyevkdw", "uktpvfwcxpdsbcab", "vlfhtlqdkfggxusma", "hwpqpwnyhlsokrnmqsjoiwzzlbaylv", "v", "chlkdhpklasay", "xxmzssfoezczvzgpfo", "jpjtklksindyjpvxhejtucset", "fvpnqduebposjklmptdgkaq", "fmcuvpikdrkmizblymnnmwkf", "niqlyrl", "qyvabwvkhxntbacvhh", "vtvvkozwytlhotpop", "lmcxjhudvumnlkqqatflyy", "lgxzmasivwpcnnywnuufsbqesoscr", "yjsfppmproitzv", "lsbkldsem", "tcttcfjhgqnqhvqzyrzjxfn", "udmole", "tgyjtnxdghchkxkuoghexfeydburw", "yxllvjhndb", "qietoq", "uzhhtul", "exzqsigutm", "as", "viamlaohkbtrqmdnunrdowqtgsufkr", "oylmvpsskwlpdtmztdklelxvjux", "gubgujoxhecsq", "efalejqivgmq", "w", "vjxgyddxmq", "ujrpd", "fszpdvkmtqvmilcdv", "xjminawbblb", "ibr", "fewgitzw", "hprffroywugiqixdknxtpjaemi", "tzoyuqczdnfemsgsh", "phumhrexialxyxvm", "wganyhzv", "jzgcvsbkwiuxwxgylygy", "pipxxfxwabozs", "uzbdpprpidfpxcxvlbwqxxc", "dexkfmajwahmpjflbpgkyxr", "hngxdllbufoxleyupznzbczfegpuf", "rbwvfflgygzp", "lerks", "ugquinvxvmunbobyrcijzdoicfmc", "bkhxnvdpugculmh", "melfmzenlvvwpfkjnxofddey", "bogwobnsdavmjsxmxjwyleun", "jq", "ytxacmohwkofts", "jdsmtiozbthrloguxrzjqfhxxme", "bzr", "fyukylveigjhlfiossgg", "kiipzcvlkhohfq", "fybxdscdofqimy", "m", "dzdjmuttofxznkvsod", "rwu", "ndxzwenrdcs", "uskykkpcc", "fkxcdrlkdgtmfaiagufsj", "mcray", "lginz", "cizjbmlxvnzsjrftbtskactlhmz", "daeqjfvdrxiatudqtvxkmxphr", "imrtbxgvhnritxey", "hozwsvgpzkxtqxryw", "aazwqjdjnmkslyxghvne", "eeixrpaitag", "diey", "kfiedkessuqkzgkngjmkrf", "zezwofckwwaomftf", "ueezorhcwakyvxc", "yjxdpjirmlu", "xjytzoijlxbn", "gzaviosoghcbbctkiamrqjgkuu", "tdeqplbibvuk", "jdbxxvn", "p", "zpbwccizb", "edyiiiyytlmcuqpwfmxxsohsoc", "v", "hykbrkqquxzallurqolcmxmlaob", "vgfiaxcyebvbgadcgo", "qzdlklmxudpjxayxiayjwtatagnxn", "mxbgmhobasjyizza", "cqsgdpdcvguopyjedi", "vwpmuqsyjwhwz", "ymzf", "dvj", "qtvyfppmtiwxziyzxkdvohelxmky", "ghomebkvdmegwqmgvpghtmvqx", "eirhdshiwbyiquhyyobnewvtplfgos", "bbjpktlbsckegcveathdoaprpaiiqm", "ujyeuhxiynwwhqblghcphvjkkvrlgh", "qp", "efcvlwujvbokzhnedwmlmceasa", "elnqxwypigkcrwksxzivm", "jviuoywbndunuufqqze", "bgxab", "iirugjnftmpue", "kec", "zzlstxdpqgqbyzsfdfrcbwse", "ulbicasvbqyzpxxlhqilwhqv", "dswdttupegpi", "joervjpln", "fzr", "oyljhnpab", "zfcefvmfvwwqgredyvlqcpftryx", "ozco", "iwihjdwncjpcqayvjaxpqozef", "aimibrckokvzu", "kwokiotuincktuihdh", "xpcervhuhnonfltyoqo", "emtvfafnrxkskzluueftozdc", "nzzcrp", "qgxevxqqpoapwzvomniwwvjecoddqg", "qmqhnhlhpzzsilzszmxgiywwabedx", "oledlvkixxlniku", "cw", "nxrwiqjfpf", "aecjgrlshvckeairimjzmt", "wlivvdjvxhagdrvjgna", "ts", "hxzxuyswcslfjurtgsk", "hcoomavhnegnek", "zxtctesskgziqnmgobjfexjpzbehj", "rekengyardexvhpavvtrgjyk", "upsvbkohtcmhb", "r", "ivqcjggr", "xjmajfjgbztfbuclpkjgroieqaympy", "icymijcffnwjqqye", "yllrgl", "vf", "fnnlgrgcedbdjckyzpdglanozd", "krmkafyuszgiacaxgvauohxd", "jahf", "zpcgg", "lnwsexzg", "ol", "qslmoxaukcqcxmafx", "couixcjyadowkyjrqpjqf", "nrspfiy", "gyzmorecjuqhnsswlfakkucooq", "mtxnwhkukkzbmwinkempus", "oxu", "xpgauwpehnopfj", "tjojwxuptjy", "whyjcgipmzjuz", "ajqfwhhmbmilkuyonqm", "lzvlzatwyfpzjykqyqffdo", "qmzvwfdnqaeeqwnq", "uppfsekumxkqxfacnwerhh", "rwhdsiatdjmefqqpvrigyqx", "nmydtkrjdsxgongzvqhguy", "maicftfg", "lqxbuixfbtn", "s", "alnvwoelsmgudfngj", "qgybscbkzwacbmewgxp", "fhqkckrjo", "sxhcrebvnesoombuojxegnv", "iswgoi", "yr", "qbggupjlgozeeaixzkxjttbwqbzur", "rydgau", "hxumnhwkvnumr", "cczsldgguk", "qdldwaeagjxwrlakmiloqzan", "xxduxsvxbcghbqlowcfosciscmp", "rbtkcvtolrnloqdwmckflkp", "prmp", "whplwmndffhqobx", "d", "rncewagdkchrhtjrnzmadpcwfjb", "eobptzai", "rwk", "jegkfltfhuwrtnizdnwiws", "klrpovcwydaemnohsie", "ayzvgoibazfskub", "iofebmcsyeetbpwotiwh", "fuferesjlbfttbaewyshhrtbxzubwz", "rrtkenfavrcylu"}

	words2 = []string{"kqn", "eb", "gxfghgnbhtovustrozx", "odhlkiozl", "epercrztztzezrtbl", "flxtxebczkieasv", "xmqasqngd", "wnfzotywjgcdxhwpvna", "sxdagxl", "dpzvqzk", "ecjkfoscbegkehjivusighoqqnk", "rqfourzdrcznxk", "oqru", "y", "ung", "qsrqlewfsxvalrqwkwxgvijexcz", "cmsxlyopwcccjcj", "jxymbeionjjkdtqedtgj", "iaozuvwrdf", "cfouijvujkydd", "pgoxrmocffagasup", "dnokakeyuzkfq", "hfulhdmo", "rdxztedplabiuztycfdfgeq", "vhxxgnz", "kekefyp", "enoxfxeokmkqjsia", "wtdnmypkd", "xkeka", "st", "tontipglxozzdbisvivvo", "nokgykgyasqdjawogrzwhyt", "zbvdxczpsvnmdfxhnfvkqbagllqy", "n", "fobiqnitfwnkk", "e", "flirzsqhedcvzqkdolizgkpvd", "maolhnokjhfhfmagsguzyidsw", "meuiinfkihlhvpnpuqzsozsgavmxym", "duifeolcntqb", "lgwfxarv", "bwmxltx", "edbjcvtuckvlrhlzzxzhpzuascoc", "vk", "fqdsbpcettrclxm", "tlsvzglozwlxopzfr", "ecozqjgaqkydphispvthfk", "mfwnnoorheu", "hfcjipwjeydmrmmsaepln", "ddrjgzbnhcjkyjjpobadrfpczsler", "dotnrswglfbgvxzcm", "ojanpanmxlgijinesftli", "sosrajezvi", "zntoxwfbsqmwbiuinxjmwjlvjnglj", "rjgzbwbopuuguvomdrhjd", "hlcjjedjuthesdtkxulbkf", "fqquhhppufiwk", "yllulqmjkbborqdjadtnotyy", "kkoioemifyifbk", "icw", "oyirnuaicvdtdojtac", "jbrgthooif", "hkmbqlyxxqixpip", "sajllhmynbyohpszhlrgsywb", "fwoflyrshsnvenegpftkowsw", "xeosmulijkgkvssmyccxzqrzx", "yunftqycasmvzaodanzvetqudxi", "muusipbyc", "cbatbdbvzkuinz", "afhnfwiuxyohxsr", "skyhv", "adqulaxacdapzwum", "ociilqnruscpgfi", "vedudlxcoybmrkkhsuw", "esnhkkqzyzpat", "kmwdlpfkurztgmvkjxjaaot", "eduwvewmpmxi", "gutstyzyvypmnvbvkpiewqiwgbr", "jyktjhztps", "keocndbcjlqhowjfx", "lavyidmwcttjbjoxyuoloaamjjijp", "oymkjiodvdtgoxvrwtgyuktbecdmo", "whldzabrvr", "ptsyvusbhkwuxctf", "zabbnaqzyangne", "almivgofy", "voopqlzufjwzgtpavyxhreg", "whxcsjcbznwzzxavafgtg", "kxxzrheqylitupzrqnctohtvvqcab", "px", "eccqatnkwbvsrgmd", "ivczfuwypdqtsmbtgmxxcxaus", "bjf", "hbtohelg", "mrwtmatce", "pp", "wpycdwttrapxai", "jdymvpzptouy", "wagbsctivcseqtjaumhbflt", "upuwtlqugbhvxsxgtwrb", "yepspbp", "olcrdqavrthqdbasfnt", "mfkfoczefkaczylkgif", "pd", "rpppwxnupezk", "ivs", "buvpbsxftxdfldspyfxrsbwl", "uhvdkwlwjdvznntftlgii", "lblsoyfgy", "msmfswsnjksztmvgdypjajdczvn", "bqobaouojjm", "ojintbkd", "tc", "anucgoziakuo", "dfxovsgwghydesjme", "s", "vavapblsflgxkvcm", "rmledxxatstcryjtib", "lbdli", "wm", "cvzcdkzz", "dorqwghzzbkangcrttdrdjovp", "gfcrufdagwkz", "hmvbhdsgsrfmi", "vgqcsarsbsjovmggclemfatvoisdu", "sdbonteic", "obomsujilwtpxtrchum", "tipq", "ajbngjobftuhd", "mvrew", "uqcytkmzmjfpsxwlfgqejcxies", "pvbqmajwbrphpz", "affbrxzgfoohlyc", "gotqcpj", "iqzbkpkhxlbkrxuabhd", "sf", "tqguluxwrqghwfwfzqb", "hkudatagdnzqboafneftqjxcrhxgw", "uhdeuebnhmvsevaxqiseqw", "djgajmhofotlmnozgdi", "odwqtgqtbtwrmuxhuepnmlyqzkzpiu", "lnqdzftjftokwgn", "rydjeovuhwgbvmidzxrocdldrmxsmb", "xfhnzkxgedxdupwtptff", "nxuxx", "hghjisddfamudmh", "hlyitt", "dw", "pesyfl", "ccnaipacnankyiupbhjivra", "abajywpwvdxuqsggpe", "msrskqfaaagkwhsvkqdelx", "yx", "ihryvvwttkzjdxyqv", "wodzvtpxclcdfawdkndmzpfvgzbp", "feznmuumsdsrqirfpvj", "ycdffiqqpyay", "obodlrehapwcblfsuuufjvbgbidu", "b", "awgjwfednpqrvnjozfqvmuydzcjed", "nxsry", "budezoplspawxzd", "soetyzvkcugqtodojbfelwjr", "zchiualewvmwosfagzulhkulixwa", "avaududitkvqkagyqqlumjsfbcuoq", "cueajr", "vsxphjelkkesegrout", "yos", "fndafnvxixwoyqjhrlzvxclgwhsfwl", "bqtzvrkl", "nwmwydrcmvcmkrur", "pwryqfsjhdwzoyfyotctb", "tutmcouczivtctrwlcyte", "qmgoohfiamcmqoyvrzzncvzqvuzcrk", "htogy", "bwworpdpnqnxjedrcteiebghqdh", "orkvieoajrxhqgydmjzchwcskksjex", "ic", "t", "pamgcrdezjeaslpmqqjncmt", "uvdmyvz", "jcnhfayreaylnjxpj", "efh", "rxknljpg", "yuwtalfcixee", "ycovnbwpiyejawmmi", "ybprv", "qlrcigxwljoiaxwlssbuozndmcx", "kyhygahrnkieuozpapvyfcc", "pwzfgkhddzcz", "yptmxlydnwyh", "mihukhfshtbp", "dvvcejirjqhpjjbohvtgs", "dxqsvykwhoqdurbnddzclrr", "dpwphfhpgrjvaapyvqnubscaq", "rqmv", "txqolpuiwbjltrnsqwhadglbntxhe", "crqcvylhjslhwhdqs", "izvlucjobdwiayaqvqpvov", "fyvnzliuzbqlwswrhsrzngiyn", "fhqrsqqoow", "lqwdmytnhca", "dljiw", "vttothnvznobyixmrknylgqni", "xnwzsfoscdxwaevehhsgliqlnwwp", "goie", "skobcntqszen", "viydnkccafcrlmfawortdzxui", "ipjaffchpzhxnlnbqow", "trofqukw", "yruowdttgp", "txnbtlyzmbzrfuecbufw", "vixmylkmlrgzjqpyutvqcelk", "miwezqzr", "duuiiknfcrm", "ffdsvnywehqyuzyfbrkhuiozfxorpe", "hufrczlkitvx", "bmeocwxatfxohxfqfa", "lz", "cjlflqcmybtpfcrnrtxthivanswj", "avwyurfesefufzbtendcs", "uxjxofpxncbadyypntfiqqabww", "pxfsqhyyqriawwayomeghlnygl", "hytdwdobtwhcjh", "ceayzx", "k", "udaztje", "yhyymkpidfmasoqvpdminjillmm", "iqfnbhiqsdbrbggrdywzmurx", "qpoag", "lsnmcrlqdrjmzqx", "kphyex", "oidlopayreexiyfwptnsxacxt", "lzuuwzenlwdklpohb", "jlnufqqme", "nlvwzlxhockyoucdjynulddzpylt", "tybiksxwuumpxwnuxfeqgizkf", "clohfofkkkfoffnh", "gqff", "twqyle", "ug", "fclmkugcmtqv", "kxjfiwxpmhmthijjiexbdgcpujkbn", "pdmlmppbrgi", "udvigkpzkssgasnhwhkivxd", "ytuhsbggfcxfzalzacxrrkmibx", "trplynrpejbuvxulbcrbnxueyhnxxd", "zhbimidnledyheunkguehanhomnjnl", "aervcdjl", "ixpjfbaly"}
	require.Equal(t, 3, countWords(words1, words2))
	words1 = []string{"apple", "banana", "cherry"}
	words2 = []string{"apple", "banana", "cherry"}
	require.Equal(t, 3, countWords(words1, words2))
	words1 = []string{"a", "ab"}
	words2 = []string{"a", "a", "a", "ab"}
	require.Equal(t, 1, countWords(words1, words2))
}

func TestReverseInt(t *testing.T) {
	var n int
	n = 15
	reverseInt(&n)
	require.Equal(t, n, 51)
	n = 981
	reverseInt(&n)
	require.Equal(t, n, 189)
	n = 500
	reverseInt(&n)
	require.Equal(t, n, 5)
	n = -15
	reverseInt(&n)
	require.Equal(t, n, -51)
	n = -90
	reverseInt(&n)
	require.Equal(t, n, -9)
}

func TestReverseInt2(t *testing.T) {
	var n int
	n = 15
	reverseInt2(&n)
	require.Equal(t, n, 51)
	n = 981
	reverseInt2(&n)
	require.Equal(t, n, 189)
	n = 500
	reverseInt2(&n)
	require.Equal(t, n, 5)
	n = -15
	reverseInt2(&n)
	require.Equal(t, n, -51)
	n = -90
	reverseInt2(&n)
	require.Equal(t, n, -9)
}

func TestMaxChar(t *testing.T) {
	require.Equal(t, 'c', maxChar("abcccccccd"))
	require.Equal(t, '1', maxChar("apple 1231111"))
}

func TestFizzBuzz(t *testing.T) {
	require.Equal(t, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz"}, fizzBuzz(10))
	require.Equal(t, []string{"1", "2", "Fizz"}, fizzBuzz(3))
	require.Equal(t, []string{"1", "2", "Fizz", "4", "Buzz"}, fizzBuzz(5))
	require.Equal(t, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}, fizzBuzz(15))
}

func TestChunkSlice(t *testing.T) {
	require.Equal(t, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}}, chunkSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3))
	require.Equal(t, [][]int{{1, 2, 3, 4, 5}}, chunkSlice([]int{1, 2, 3, 4, 5}, 10))
	require.Equal(t, [][]int{{1}, {2}, {3}, {4}, {5}}, chunkSlice([]int{1, 2, 3, 4, 5}, 1))
	require.Equal(t, [][]int{{1}, {2}, {3}, {4}, {5}}, chunkSlice([]int{1, 2, 3, 4, 5}, 1))
	require.Equal(t, [][]int{{1, 2}, {3, 4}, {5}}, chunkSlice([]int{1, 2, 3, 4, 5}, 2))
	require.Equal(t, [][]int{{1, 2, 3, 4, 5}}, chunkSlice([]int{1, 2, 3, 4, 5}, 6))
	require.Equal(t, [][]int{{1, 2, 3, 4, 5}}, chunkSlice([]int{1, 2, 3, 4, 5}, 5))
	require.Equal(t, [][]int{{1, 2, 3, 4}, {5}}, chunkSlice([]int{1, 2, 3, 4, 5}, 4))
}

func TestChunkSlice2(t *testing.T) {
	require.Equal(t, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}}, chunkSlice2([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3))
	require.Equal(t, [][]int{{1, 2, 3, 4, 5}}, chunkSlice2([]int{1, 2, 3, 4, 5}, 10))
	require.Equal(t, [][]int{{1}, {2}, {3}, {4}, {5}}, chunkSlice2([]int{1, 2, 3, 4, 5}, 1))
	require.Equal(t, [][]int{{1}, {2}, {3}, {4}, {5}}, chunkSlice2([]int{1, 2, 3, 4, 5}, 1))
	require.Equal(t, [][]int{{1, 2}, {3, 4}, {5}}, chunkSlice2([]int{1, 2, 3, 4, 5}, 2))
	require.Equal(t, [][]int{{1, 2, 3, 4, 5}}, chunkSlice2([]int{1, 2, 3, 4, 5}, 6))
	require.Equal(t, [][]int{{1, 2, 3, 4, 5}}, chunkSlice2([]int{1, 2, 3, 4, 5}, 5))
	require.Equal(t, [][]int{{1, 2, 3, 4}, {5}}, chunkSlice2([]int{1, 2, 3, 4, 5}, 4))
}

func TestAnagrams(t *testing.T) {
	require.Equal(t, true, anagrams("rail safety", "fairy tales"))
	require.Equal(t, true, anagrams("RAIL! SAFETY!", "fairy tales"))
	require.Equal(t, false, anagrams("Hi there", "Bye there"))
}

func TestBananagrams(t *testing.T) {
	require.Equal(t, true, bananagrams("rail safety", "fairy tales"))
	require.Equal(t, true, bananagrams("RAIL! SAFETY!", "fairy tales"))
	require.Equal(t, false, bananagrams("Hi there", "Bye there"))
}

func TestCapitalize(t *testing.T) {
	require.Equal(t, "Hi There", capitalize("hi there"))
	require.Equal(t, "Hi There, How Is It Going?", capitalize("hi there, how is it going?"))
}

func TestCapitalize2(t *testing.T) {
	require.Equal(t, "Hi There", capitalize2("hi there"))
	require.Equal(t, "Hi There, How Is It Going?", capitalize("hi there, how is it going?"))
}

func TestSteps(t *testing.T) {
	// Save the original stdout
	originalStdout := os.Stdout

	// Create a pipe for stdout
	r, w, _ := os.Pipe()
	// Set stdout to the write side of the pipe
    os.Stdout = w
	// Call the function that writes to stdout
	steps(10)
	// Close the write side of the pipe when done
	// Can't defer this because we need to read from the pipe
	w.Close()
	// Read the output from stdout
	out, _ := io.ReadAll(r)
	// Restore originalStdout
    os.Stdout = originalStdout

	// Check the contents of the buffer
	require.Equal(t, "'#         '\n'##        '\n'###       '\n'####      '\n'#####     '\n'######    '\n'#######   '\n'########  '\n'######### '\n'##########'\n", string(out))
}
