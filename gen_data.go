package z18n

import "golang.org/x/text/language"

var mapping = map[language.Tag]*localize{
	language.MustParse("af"):         l_af,
	language.MustParse("af-NA"):      l_af_NA,
	language.MustParse("am"):         l_am,
	language.MustParse("ar"):         l_ar,
	language.MustParse("ar-AE"):      l_ar_AE,
	language.MustParse("ar-BH"):      l_ar_AE, // Link
	language.MustParse("ar-DJ"):      l_ar_AE, // Link
	language.MustParse("ar-DZ"):      l_ar_DZ,
	language.MustParse("ar-EG"):      l_ar_AE, // Link
	language.MustParse("ar-IL"):      l_ar_IL,
	language.MustParse("ar-IQ"):      l_ar_IQ,
	language.MustParse("ar-JO"):      l_ar_JO,
	language.MustParse("ar-KM"):      l_ar_KM,
	language.MustParse("ar-KW"):      l_ar_AE, // Link
	language.MustParse("ar-LB"):      l_ar_LB,
	language.MustParse("ar-LY"):      l_ar_AE, // Link
	language.MustParse("ar-MA"):      l_ar_MA,
	language.MustParse("ar-MR"):      l_ar_MR,
	language.MustParse("ar-OM"):      l_ar_AE, // Link
	language.MustParse("ar-PS"):      l_ar_LB, // Link
	language.MustParse("ar-QA"):      l_ar_AE, // Link
	language.MustParse("ar-SA"):      l_ar_SA,
	language.MustParse("ar-SD"):      l_ar_AE, // Link
	language.MustParse("ar-SY"):      l_ar_JO, // Link
	language.MustParse("ar-TN"):      l_ar_TN,
	language.MustParse("ar-YE"):      l_ar_SA, // Link
	language.MustParse("as"):         l_as,
	language.MustParse("az"):         l_az,
	language.MustParse("be"):         l_be,
	language.MustParse("bg"):         l_bg,
	language.MustParse("bn"):         l_bn,
	language.MustParse("bn-IN"):      l_bn_IN,
	language.MustParse("bs"):         l_bs,
	language.MustParse("ca"):         l_ca,
	language.MustParse("ca-AD"):      l_ca_AD,
	language.MustParse("ca-FR"):      l_ca_AD, // Link
	language.MustParse("ca-IT"):      l_ca_AD, // Link
	language.MustParse("cs"):         l_cs,
	language.MustParse("cy"):         l_cy,
	language.MustParse("da"):         l_da,
	language.MustParse("de"):         l_de,
	language.MustParse("de-AT"):      l_de_AT,
	language.MustParse("de-IT"):      l_de_AT, // Link
	language.MustParse("el"):         l_el,
	language.MustParse("en"):         l_en,
	language.MustParse("en-001"):     l_en_001,
	language.MustParse("en-150"):     l_en_150,
	language.MustParse("en-AE"):      l_en_AE,
	language.MustParse("en-AG"):      l_en_AG,
	language.MustParse("en-AI"):      l_en_150, // Link
	language.MustParse("en-AS"):      l_en_AS,
	language.MustParse("en-AT"):      l_en_150, // Link
	language.MustParse("en-AU"):      l_en_AU,
	language.MustParse("en-BB"):      l_en_001, // Link
	language.MustParse("en-BE"):      l_en_BE,
	language.MustParse("en-BI"):      l_en_BI,
	language.MustParse("en-BM"):      l_en_001, // Link
	language.MustParse("en-BS"):      l_en_AG,  // Link
	language.MustParse("en-BW"):      l_en_BW,
	language.MustParse("en-BZ"):      l_en_BZ,
	language.MustParse("en-CA"):      l_en_CA,
	language.MustParse("en-CC"):      l_en_150, // Link
	language.MustParse("en-CH"):      l_en_150, // Link
	language.MustParse("en-CK"):      l_en_150, // Link
	language.MustParse("en-CM"):      l_en_150, // Link
	language.MustParse("en-CX"):      l_en_150, // Link
	language.MustParse("en-CY"):      l_en_001, // Link
	language.MustParse("en-DE"):      l_en_150, // Link
	language.MustParse("en-DG"):      l_en_150, // Link
	language.MustParse("en-DK"):      l_en_DK,
	language.MustParse("en-DM"):      l_en_AG,  // Link
	language.MustParse("en-ER"):      l_en_001, // Link
	language.MustParse("en-FI"):      l_en_DK,  // Link
	language.MustParse("en-FJ"):      l_en_001, // Link
	language.MustParse("en-FK"):      l_en_150, // Link
	language.MustParse("en-FM"):      l_en_001, // Link
	language.MustParse("en-GB"):      l_en_GB,
	language.MustParse("en-GD"):      l_en_001, // Link
	language.MustParse("en-GG"):      l_en_150, // Link
	language.MustParse("en-GH"):      l_en_001, // Link
	language.MustParse("en-GI"):      l_en_150, // Link
	language.MustParse("en-GM"):      l_en_001, // Link
	language.MustParse("en-GU"):      l_en_AS,  // Link
	language.MustParse("en-GY"):      l_en_001, // Link
	language.MustParse("en-HK"):      l_en_HK,
	language.MustParse("en-IE"):      l_en_IE,
	language.MustParse("en-IL"):      l_en_IL,
	language.MustParse("en-IM"):      l_en_150, // Link
	language.MustParse("en-IN"):      l_en_IN,
	language.MustParse("en-IO"):      l_en_150, // Link
	language.MustParse("en-JE"):      l_en_150, // Link
	language.MustParse("en-JM"):      l_en_AG,  // Link
	language.MustParse("en-KE"):      l_en_IL,  // Link
	language.MustParse("en-KI"):      l_en_001, // Link
	language.MustParse("en-KN"):      l_en_001, // Link
	language.MustParse("en-KY"):      l_en_001, // Link
	language.MustParse("en-LC"):      l_en_001, // Link
	language.MustParse("en-LR"):      l_en_001, // Link
	language.MustParse("en-LS"):      l_en_001, // Link
	language.MustParse("en-MG"):      l_en_150, // Link
	language.MustParse("en-MH"):      l_en_AS,  // Link
	language.MustParse("en-MO"):      l_en_AG,  // Link
	language.MustParse("en-MS"):      l_en_150, // Link
	language.MustParse("en-MT"):      l_en_MT,
	language.MustParse("en-MU"):      l_en_150, // Link
	language.MustParse("en-MW"):      l_en_001, // Link
	language.MustParse("en-MY"):      l_en_001, // Link
	language.MustParse("en-NA"):      l_en_001, // Link
	language.MustParse("en-NF"):      l_en_150, // Link
	language.MustParse("en-NG"):      l_en_150, // Link
	language.MustParse("en-NL"):      l_en_150, // Link
	language.MustParse("en-NR"):      l_en_150, // Link
	language.MustParse("en-NU"):      l_en_150, // Link
	language.MustParse("en-NZ"):      l_en_NZ,
	language.MustParse("en-PG"):      l_en_001, // Link
	language.MustParse("en-PH"):      l_en_AG,  // Link
	language.MustParse("en-PK"):      l_en_PK,
	language.MustParse("en-PN"):      l_en_150, // Link
	language.MustParse("en-PR"):      l_en_AS,  // Link
	language.MustParse("en-PW"):      l_en_001, // Link
	language.MustParse("en-RW"):      l_en_150, // Link
	language.MustParse("en-SB"):      l_en_001, // Link
	language.MustParse("en-SC"):      l_en_150, // Link
	language.MustParse("en-SD"):      l_en_SD,
	language.MustParse("en-SE"):      l_en_SE,
	language.MustParse("en-SG"):      l_en_SG,
	language.MustParse("en-SH"):      l_en_150, // Link
	language.MustParse("en-SI"):      l_en_150, // Link
	language.MustParse("en-SL"):      l_en_001, // Link
	language.MustParse("en-SS"):      l_en_001, // Link
	language.MustParse("en-SX"):      l_en_150, // Link
	language.MustParse("en-SZ"):      l_en_001, // Link
	language.MustParse("en-TC"):      l_en_001, // Link
	language.MustParse("en-TK"):      l_en_150, // Link
	language.MustParse("en-TO"):      l_en_001, // Link
	language.MustParse("en-TT"):      l_en_AG,  // Link
	language.MustParse("en-TV"):      l_en_150, // Link
	language.MustParse("en-TZ"):      l_en_150, // Link
	language.MustParse("en-UG"):      l_en_150, // Link
	language.MustParse("en-UM"):      l_en_AS,  // Link
	language.MustParse("en-VC"):      l_en_001, // Link
	language.MustParse("en-VG"):      l_en_001, // Link
	language.MustParse("en-VI"):      l_en_AS,  // Link
	language.MustParse("en-VU"):      l_en_001, // Link
	language.MustParse("en-WS"):      l_en_AG,  // Link
	language.MustParse("en-ZA"):      l_en_ZA,
	language.MustParse("en-ZM"):      l_en_001, // Link
	language.MustParse("en-ZW"):      l_en_ZW,
	language.MustParse("es"):         l_es,
	language.MustParse("es-419"):     l_es_419,
	language.MustParse("es-AR"):      l_es_419, // Link
	language.MustParse("es-BO"):      l_es_BO,
	language.MustParse("es-BR"):      l_es_BR,
	language.MustParse("es-BZ"):      l_es_BR, // Link
	language.MustParse("es-CL"):      l_es_CL,
	language.MustParse("es-CO"):      l_es_CO,
	language.MustParse("es-CR"):      l_es_419, // Link
	language.MustParse("es-CU"):      l_es_419, // Link
	language.MustParse("es-DO"):      l_es_DO,
	language.MustParse("es-EC"):      l_es_419, // Link
	language.MustParse("es-GT"):      l_es_GT,
	language.MustParse("es-HN"):      l_es_HN,
	language.MustParse("es-MX"):      l_es_MX,
	language.MustParse("es-NI"):      l_es_BR, // Link
	language.MustParse("es-PA"):      l_es_PA,
	language.MustParse("es-PE"):      l_es_PE,
	language.MustParse("es-PH"):      l_es_PH,
	language.MustParse("es-PR"):      l_es_PA, // Link
	language.MustParse("es-PY"):      l_es_PY,
	language.MustParse("es-SV"):      l_es_BR, // Link
	language.MustParse("es-US"):      l_es_US,
	language.MustParse("es-UY"):      l_es_UY,
	language.MustParse("es-VE"):      l_es_VE,
	language.MustParse("et"):         l_et,
	language.MustParse("eu"):         l_eu,
	language.MustParse("fa"):         l_fa,
	language.MustParse("fa-AF"):      l_fa_AF,
	language.MustParse("fi"):         l_fi,
	language.MustParse("fil"):        l_fil,
	language.MustParse("fr"):         l_fr,
	language.MustParse("fr-BE"):      l_fr_BE,
	language.MustParse("fr-CA"):      l_fr_CA,
	language.MustParse("fr-CH"):      l_fr_CH,
	language.MustParse("fr-CM"):      l_fr_CM,
	language.MustParse("fr-DJ"):      l_fr_DJ,
	language.MustParse("fr-DZ"):      l_fr_DJ, // Link
	language.MustParse("fr-MA"):      l_fr_MA,
	language.MustParse("fr-ML"):      l_fr_ML,
	language.MustParse("fr-MR"):      l_fr_MR,
	language.MustParse("fr-SY"):      l_fr_DJ, // Link
	language.MustParse("fr-TD"):      l_fr_MR, // Link
	language.MustParse("fr-TN"):      l_fr_MR, // Link
	language.MustParse("fr-VU"):      l_fr_MR, // Link
	language.MustParse("ga"):         l_ga,
	language.MustParse("gl"):         l_gl,
	language.MustParse("gu"):         l_gu,
	language.MustParse("he"):         l_he,
	language.MustParse("hi"):         l_hi,
	language.MustParse("hr"):         l_hr,
	language.MustParse("hr-BA"):      l_hr_BA,
	language.MustParse("hu"):         l_hu,
	language.MustParse("hy"):         l_hy,
	language.MustParse("id"):         l_id,
	language.MustParse("is"):         l_is,
	language.MustParse("it"):         l_it,
	language.MustParse("it-CH"):      l_it_CH,
	language.MustParse("ja"):         l_ja,
	language.MustParse("jv"):         l_jv,
	language.MustParse("ka"):         l_ka,
	language.MustParse("kk"):         l_kk,
	language.MustParse("km"):         l_km,
	language.MustParse("kn"):         l_kn,
	language.MustParse("ko"):         l_ko,
	language.MustParse("ky"):         l_ky,
	language.MustParse("lo"):         l_lo,
	language.MustParse("lt"):         l_lt,
	language.MustParse("lv"):         l_lv,
	language.MustParse("mk"):         l_mk,
	language.MustParse("ml"):         l_ml,
	language.MustParse("mn"):         l_mn,
	language.MustParse("mr"):         l_mr,
	language.MustParse("ms"):         l_ms,
	language.MustParse("ms-BN"):      l_ms_BN,
	language.MustParse("ms-ID"):      l_ms_ID,
	language.MustParse("ms-SG"):      l_ms_SG,
	language.MustParse("my"):         l_my,
	language.MustParse("nb"):         l_nb,
	language.MustParse("ne"):         l_ne,
	language.MustParse("ne-IN"):      l_ne_IN,
	language.MustParse("nl"):         l_nl,
	language.MustParse("nl-BE"):      l_nl_BE,
	language.MustParse("nn"):         l_nn,
	language.MustParse("or"):         l_or,
	language.MustParse("pa"):         l_pa,
	language.MustParse("pl"):         l_pl,
	language.MustParse("ps"):         l_ps,
	language.MustParse("ps-PK"):      l_ps_PK,
	language.MustParse("pt"):         l_pt,
	language.MustParse("pt-AO"):      l_pt_AO,
	language.MustParse("pt-CH"):      l_pt_CH,
	language.MustParse("pt-CV"):      l_pt_AO, // Link
	language.MustParse("pt-GQ"):      l_pt_AO, // Link
	language.MustParse("pt-GW"):      l_pt_AO, // Link
	language.MustParse("pt-LU"):      l_pt_CH, // Link
	language.MustParse("pt-MO"):      l_pt_MO,
	language.MustParse("pt-MZ"):      l_pt_AO, // Link
	language.MustParse("pt-PT"):      l_pt_AO, // Link
	language.MustParse("pt-ST"):      l_pt_AO, // Link
	language.MustParse("pt-TL"):      l_pt_AO, // Link
	language.MustParse("ro"):         l_ro,
	language.MustParse("ro-MD"):      l_ro_MD,
	language.MustParse("root"):       l_root,
	language.MustParse("ru"):         l_ru,
	language.MustParse("sd"):         l_sd,
	language.MustParse("si"):         l_si,
	language.MustParse("sk"):         l_sk,
	language.MustParse("sl"):         l_sl,
	language.MustParse("so"):         l_so,
	language.MustParse("so-DJ"):      l_so_DJ,
	language.MustParse("so-ET"):      l_so_ET,
	language.MustParse("so-KE"):      l_so_KE,
	language.MustParse("sq"):         l_sq,
	language.MustParse("sq-MK"):      l_sq_MK,
	language.MustParse("sq-XK"):      l_sq_MK, // Link
	language.MustParse("sr"):         l_sr,
	language.MustParse("sr-Cyrl-BA"): l_sr_Cyrl_BA,
	language.MustParse("sr-Cyrl-ME"): l_sr_Cyrl_ME,
	language.MustParse("sr-Cyrl-XK"): l_sr_Cyrl_XK,
	language.MustParse("sr-Latn"):    l_sr_Latn,
	language.MustParse("sr-Latn-BA"): l_sr_Latn_BA,
	language.MustParse("sr-Latn-ME"): l_sr_Latn_ME,
	language.MustParse("sr-Latn-XK"): l_sr_Latn_XK,
	language.MustParse("sv"):         l_sv,
	language.MustParse("sv-AX"):      l_sv_AX,
	language.MustParse("sv-FI"):      l_sv_AX, // Link
	language.MustParse("sw"):         l_sw,
	language.MustParse("sw-KE"):      l_sw_KE,
	language.MustParse("ta"):         l_ta,
	language.MustParse("ta-LK"):      l_ta_LK,
	language.MustParse("ta-SG"):      l_ta_SG,
	language.MustParse("te"):         l_te,
	language.MustParse("th"):         l_th,
	language.MustParse("tk"):         l_tk,
	language.MustParse("tr"):         l_tr,
	language.MustParse("tr-CY"):      l_tr_CY,
	language.MustParse("uk"):         l_uk,
	language.MustParse("ur"):         l_ur,
	language.MustParse("ur-IN"):      l_ur_IN,
	language.MustParse("uz"):         l_uz,
	language.MustParse("vi"):         l_vi,
	language.MustParse("yue"):        l_yue,
	language.MustParse("yue-Hant"):   l_yue, // Link
	language.MustParse("zh"):         l_zh,
	language.MustParse("zh-Hans-HK"): l_zh_Hans_HK,
	language.MustParse("zh-Hans-MO"): l_zh_Hans_HK, // Link
	language.MustParse("zh-Hans-SG"): l_zh_Hans_SG,
	language.MustParse("zh-Hant"):    l_zh_Hant,
	language.MustParse("zh-Hant-HK"): l_zh_Hant_HK,
	language.MustParse("zh-Hant-MO"): l_zh_Hant_HK, // Link
	language.MustParse("zu"):         l_zu,
}

var l_af = &localize{
	am:              "vm.",
	pm:              "nm.",
	firstDay:        6,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan.", "Feb.", "Mrt.", "Apr.", "Mei", "Jun.", "Jul.", "Aug.", "Sep.", "Okt.", "Nov.", "Des."},
		long:  [12]string{"Januarie", "Februarie", "Maart", "April", "Mei", "Junie", "Julie", "Augustus", "September", "Oktober", "November", "Desember"},
	},
	days: days{
		short: [7]string{"So.", "Ma.", "Wo.", "Di.", "Do.", "Vr.", "Sa."},
		long:  [7]string{"Sondag", "Maandag", "Woensdag", "Dinsdag", "Donderdag", "Vrydag", "Saterdag"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 02 January 2006", "02 January 2006", "02 Jan 2006", "2006-01-02"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 02 January 2006 15:04:05 MST -0700", "02 January 2006 15:04:05 MST", "02 Jan 2006 15:04:05", "2006-01-02 15:04"},
	},
}
var l_af_NA = &localize{
	am:              "vm.",
	pm:              "nm.",
	firstDay:        6,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan.", "Feb.", "Mrt.", "Apr.", "Mei", "Jun.", "Jul.", "Aug.", "Sep.", "Okt.", "Nov.", "Des."},
		long:  [12]string{"Januarie", "Februarie", "Maart", "April", "Mei", "Junie", "Julie", "Augustus", "September", "Oktober", "November", "Desember"},
	},
	days: days{
		short: [7]string{"So.", "Ma.", "Wo.", "Di.", "Do.", "Vr.", "Sa."},
		long:  [7]string{"Sondag", "Maandag", "Woensdag", "Dinsdag", "Donderdag", "Vrydag", "Saterdag"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 02 January 2006", "02 January 2006", "02 Jan 2006", "2006-01-02"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday 02 January 2006 3:04:05 PM MST -0700", "02 January 2006 3:04:05 PM MST", "02 Jan 2006 3:04:05 PM", "2006-01-02 3:04 PM"},
	},
}
var l_am = &localize{
	am:              "ጥዋት",
	pm:              "ከሰዓት",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ጃንዩ", "ፌብሩ", "ማርች", "ኤፕሪ", "ሜይ", "ጁን", "ጁላይ", "ኦገስ", "ሴፕቴ", "ኦክቶ", "ኖቬም", "ዲሴም"},
		long:  [12]string{"ጃንዩወሪ", "ፌብሩወሪ", "ማርች", "ኤፕሪል", "ሜይ", "ጁን", "ጁላይ", "ኦገስት", "ሴፕቴምበር", "ኦክቶበር", "ኖቬምበር", "ዲሴምበር"},
	},
	days: days{
		short: [7]string{"እሑድ", "ሰኞ", "ረቡዕ", "ማክሰ", "ሐሙስ", "ዓርብ", "ቅዳሜ"},
		long:  [7]string{"እሑድ", "ሰኞ", "ረቡዕ", "ማክሰኞ", "ሐሙስ", "ዓርብ", "ቅዳሜ"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006 January 2, Monday", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"2006 January 2, Monday 3:04:05 PM MST -0700", "2 January 2006 3:04:05 PM MST", "2 Jan 2006 3:04:05 PM", "02/01/2006 3:04 PM"},
	},
}
var l_ar = &localize{
	am:              "ص",
	pm:              "م",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"يناير", "فبراير", "مارس", "أبريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
		long:  [12]string{"يناير", "فبراير", "مارس", "أبريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
	},
	days: days{
		short: [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
		long:  [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday، 2 January 2006", "2 January 2006", "02\u200f/01\u200f/2006", "2\u200f/1\u200f/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday، 2 January 2006 في 3:04:05 PM MST -0700", "2 January 2006 في 3:04:05 PM MST", "02\u200f/01\u200f/2006, 3:04:05 PM", "2\u200f/1\u200f/2006, 3:04 PM"},
	},
}
var l_ar_AE = &localize{
	am:              "ص",
	pm:              "م",
	firstDay:        6,
	twentyFourHours: false,
	months: months{
		short: [12]string{"يناير", "فبراير", "مارس", "أبريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
		long:  [12]string{"يناير", "فبراير", "مارس", "أبريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
	},
	days: days{
		short: [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
		long:  [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday، 2 January 2006", "2 January 2006", "02\u200f/01\u200f/2006", "2\u200f/1\u200f/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday، 2 January 2006 في 3:04:05 PM MST -0700", "2 January 2006 في 3:04:05 PM MST", "02\u200f/01\u200f/2006, 3:04:05 PM", "2\u200f/1\u200f/2006, 3:04 PM"},
	},
}
var l_ar_DZ = &localize{
	am:              "ص",
	pm:              "م",
	firstDay:        6,
	twentyFourHours: false,
	months: months{
		short: [12]string{"جانفي", "فيفري", "مارس", "أفريل", "ماي", "جوان", "جويلية", "أوت", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
		long:  [12]string{"جانفي", "فيفري", "مارس", "أفريل", "ماي", "جوان", "جويلية", "أوت", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
	},
	days: days{
		short: [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
		long:  [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday، 2 January 2006", "2 January 2006", "02\u200f/01\u200f/2006", "2\u200f/1\u200f/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday، 2 January 2006 في 3:04:05 PM MST -0700", "2 January 2006 في 3:04:05 PM MST", "02\u200f/01\u200f/2006, 3:04:05 PM", "2\u200f/1\u200f/2006, 3:04 PM"},
	},
}
var l_ar_IL = &localize{
	am:              "ص",
	pm:              "م",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"يناير", "فبراير", "مارس", "أبريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
		long:  [12]string{"يناير", "فبراير", "مارس", "أبريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
	},
	days: days{
		short: [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
		long:  [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday، 2 January 2006", "2 January 2006", "02\u200f/01\u200f/2006", "2\u200f/1\u200f/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday، 2 January 2006 في 15:04:05 MST -0700", "2 January 2006 في 15:04:05 MST", "02\u200f/01\u200f/2006, 15:04:05", "2\u200f/1\u200f/2006, 15:04"},
	},
}
var l_ar_IQ = &localize{
	am:              "ص",
	pm:              "م",
	firstDay:        6,
	twentyFourHours: false,
	months: months{
		short: [12]string{"كانون الثاني", "شباط", "آذار", "نيسان", "أيار", "حزيران", "تموز", "آب", "أيلول", "تشرين\u00a0الأول", "تشرين الثاني", "كانون الأول"},
		long:  [12]string{"كانون الثاني", "شباط", "آذار", "نيسان", "أيار", "حزيران", "تموز", "آب", "أيلول", "تشرين الأول", "تشرين الثاني", "كانون الأول"},
	},
	days: days{
		short: [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
		long:  [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday، 2 January 2006", "2 January 2006", "02\u200f/01\u200f/2006", "2\u200f/1\u200f/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday، 2 January 2006 في 3:04:05 PM MST -0700", "2 January 2006 في 3:04:05 PM MST", "02\u200f/01\u200f/2006, 3:04:05 PM", "2\u200f/1\u200f/2006, 3:04 PM"},
	},
}
var l_ar_JO = &localize{
	am:              "ص",
	pm:              "م",
	firstDay:        6,
	twentyFourHours: false,
	months: months{
		short: [12]string{"كانون الثاني", "شباط", "آذار", "نيسان", "أيار", "حزيران", "تموز", "آب", "أيلول", "تشرين الأول", "تشرين الثاني", "كانون الأول"},
		long:  [12]string{"كانون الثاني", "شباط", "آذار", "نيسان", "أيار", "حزيران", "تموز", "آب", "أيلول", "تشرين الأول", "تشرين الثاني", "كانون الأول"},
	},
	days: days{
		short: [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
		long:  [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday، 2 January 2006", "2 January 2006", "02\u200f/01\u200f/2006", "2\u200f/1\u200f/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday، 2 January 2006 في 3:04:05 PM MST -0700", "2 January 2006 في 3:04:05 PM MST", "02\u200f/01\u200f/2006, 3:04:05 PM", "2\u200f/1\u200f/2006, 3:04 PM"},
	},
}
var l_ar_KM = &localize{
	am:              "ص",
	pm:              "م",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"يناير", "فبراير", "مارس", "أبريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
		long:  [12]string{"يناير", "فبراير", "مارس", "أبريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
	},
	days: days{
		short: [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
		long:  [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday، 2 January 2006", "2 January 2006", "02\u200f/01\u200f/2006", "2\u200f/1\u200f/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday، 2 January 2006 في 15:04:05 MST -0700", "2 January 2006 في 15:04:05 MST", "02\u200f/01\u200f/2006, 15:04:05", "2\u200f/1\u200f/2006, 15:04"},
	},
}
var l_ar_LB = &localize{
	am:              "ص",
	pm:              "م",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"كانون الثاني", "شباط", "آذار", "نيسان", "أيار", "حزيران", "تموز", "آب", "أيلول", "تشرين الأول", "تشرين الثاني", "كانون الأول"},
		long:  [12]string{"كانون الثاني", "شباط", "آذار", "نيسان", "أيار", "حزيران", "تموز", "آب", "أيلول", "تشرين الأول", "تشرين الثاني", "كانون الأول"},
	},
	days: days{
		short: [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
		long:  [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday، 2 January 2006", "2 January 2006", "02\u200f/01\u200f/2006", "2\u200f/1\u200f/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday، 2 January 2006 في 3:04:05 PM MST -0700", "2 January 2006 في 3:04:05 PM MST", "02\u200f/01\u200f/2006, 3:04:05 PM", "2\u200f/1\u200f/2006, 3:04 PM"},
	},
}
var l_ar_MA = &localize{
	am:              "ص",
	pm:              "م",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"يناير", "فبراير", "مارس", "أبريل", "ماي", "يونيو", "يوليوز", "غشت", "شتنبر", "أكتوبر", "نونبر", "دجنبر"},
		long:  [12]string{"يناير", "فبراير", "مارس", "أبريل", "ماي", "يونيو", "يوليوز", "غشت", "شتنبر", "أكتوبر", "نونبر", "دجنبر"},
	},
	days: days{
		short: [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
		long:  [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday، 2 January 2006", "2 January 2006", "02\u200f/01\u200f/2006", "2\u200f/1\u200f/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday، 2 January 2006 في 15:04:05 MST -0700", "2 January 2006 في 15:04:05 MST", "02\u200f/01\u200f/2006, 15:04:05", "2\u200f/1\u200f/2006, 15:04"},
	},
}
var l_ar_MR = &localize{
	am:              "ص",
	pm:              "م",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"يناير", "فبراير", "مارس", "إبريل", "مايو", "يونيو", "يوليو", "أغشت", "شتمبر", "أكتوبر", "نوفمبر", "دجمبر"},
		long:  [12]string{"يناير", "فبراير", "مارس", "إبريل", "مايو", "يونيو", "يوليو", "أغشت", "شتمبر", "أكتوبر", "نوفمبر", "دجمبر"},
	},
	days: days{
		short: [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
		long:  [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday، 2 January 2006", "2 January 2006", "02\u200f/01\u200f/2006", "2\u200f/1\u200f/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday، 2 January 2006 في 3:04:05 PM MST -0700", "2 January 2006 في 3:04:05 PM MST", "02\u200f/01\u200f/2006, 3:04:05 PM", "2\u200f/1\u200f/2006, 3:04 PM"},
	},
}
var l_ar_SA = &localize{
	am:              "ص",
	pm:              "م",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"يناير", "فبراير", "مارس", "أبريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
		long:  [12]string{"يناير", "فبراير", "مارس", "أبريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
	},
	days: days{
		short: [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
		long:  [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday، 2 January 2006", "2 January 2006", "02\u200f/01\u200f/2006", "2\u200f/1\u200f/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday، 2 January 2006 في 3:04:05 PM MST -0700", "2 January 2006 في 3:04:05 PM MST", "02\u200f/01\u200f/2006, 3:04:05 PM", "2\u200f/1\u200f/2006, 3:04 PM"},
	},
}
var l_ar_TN = &localize{
	am:              "ص",
	pm:              "م",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"جانفي", "فيفري", "مارس", "أفريل", "ماي", "جوان", "جويلية", "أوت", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
		long:  [12]string{"جانفي", "فيفري", "مارس", "أفريل", "ماي", "جوان", "جويلية", "أوت", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
	},
	days: days{
		short: [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
		long:  [7]string{"الأحد", "الاثنين", "الأربعاء", "الثلاثاء", "الخميس", "الجمعة", "السبت"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday، 2 January 2006", "2 January 2006", "02\u200f/01\u200f/2006", "2\u200f/1\u200f/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday، 2 January 2006 في 3:04:05 PM MST -0700", "2 January 2006 في 3:04:05 PM MST", "02\u200f/01\u200f/2006, 3:04:05 PM", "2\u200f/1\u200f/2006, 3:04 PM"},
	},
}
var l_as = &localize{
	am:              "পূৰ্বাহ্ন",
	pm:              "অপৰাহ্ন",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"জানু", "ফেব্ৰু", "মাৰ্চ", "এপ্ৰিল", "মে’", "জুন", "জুলাই", "আগ", "ছেপ্তে", "অক্টো", "নৱে", "ডিচে"},
		long:  [12]string{"জানুৱাৰী", "ফেব্ৰুৱাৰী", "মাৰ্চ", "এপ্ৰিল", "মে’", "জুন", "জুলাই", "আগষ্ট", "ছেপ্তেম্বৰ", "অক্টোবৰ", "নৱেম্বৰ", "ডিচেম্বৰ"},
	},
	days: days{
		short: [7]string{"দেও", "সোম", "বুধ", "মঙ্গল", "বৃহ", "শুক্ৰ", "শনি"},
		long:  [7]string{"দেওবাৰ", "সোমবাৰ", "বুধবাৰ", "মঙ্গলবাৰ", "বৃহস্পতিবাৰ", "শুক্ৰবাৰ", "শনিবাৰ"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January, 2006", "2 January, 2006", "02-01-2006", "2-1-2006"},
		time:     [4]string{"PM 3.04.05 MST -0700", "PM 3.04.05 MST", "PM 3.04.05", "PM 3.04"},
		datetime: [4]string{"Monday, 2 January, 2006 PM 3.04.05 MST -0700", "2 January, 2006 PM 3.04.05 MST", "02-01-2006 PM 3.04.05", "2-1-2006 PM 3.04"},
	},
}
var l_az = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"yan", "fev", "mar", "apr", "may", "iyn", "iyl", "avq", "sen", "okt", "noy", "dek"},
		long:  [12]string{"yanvar", "fevral", "mart", "aprel", "may", "iyun", "iyul", "avqust", "sentyabr", "oktyabr", "noyabr", "dekabr"},
	},
	days: days{
		short: [7]string{"B.", "B.e.", "Ç.", "Ç.a.", "C.a.", "C.", "Ş."},
		long:  [7]string{"bazar", "bazar ertəsi", "çərşənbə", "çərşənbə axşamı", "cümə axşamı", "cümə", "şənbə"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2 January 2006, Monday", "2 January 2006", "2 Jan 2006", "02.01.06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"2 January 2006, Monday 15:04:05 MST -0700", "2 January 2006 15:04:05 MST", "2 Jan 2006 15:04:05", "02.01.06 15:04"},
	},
}
var l_be = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"сту", "лют", "сак", "кра", "мая", "чэр", "ліп", "жні", "вер", "кас", "ліс", "сне"},
		long:  [12]string{"студзеня", "лютага", "сакавіка", "красавіка", "мая", "чэрвеня", "ліпеня", "жніўня", "верасня", "кастрычніка", "лістапада", "снежня"},
	},
	days: days{
		short: [7]string{"нд", "пн", "ср", "аў", "чц", "пт", "сб"},
		long:  [7]string{"нядзеля", "панядзелак", "серада", "аўторак", "чацвер", "пятніца", "субота"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006 г.", "2 January 2006 г.", "2.01.2006", "2.01.06"},
		time:     [4]string{"15:04:05, MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006 г. у 15:04:05, MST -0700", "2 January 2006 г. у 15:04:05 MST", "2.01.2006, 15:04:05", "2.01.06, 15:04"},
	},
}
var l_bg = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"яну", "фев", "март", "апр", "май", "юни", "юли", "авг", "сеп", "окт", "ное", "дек"},
		long:  [12]string{"януари", "февруари", "март", "април", "май", "юни", "юли", "август", "септември", "октомври", "ноември", "декември"},
	},
	days: days{
		short: [7]string{"нд", "пн", "ср", "вт", "чт", "пт", "сб"},
		long:  [7]string{"неделя", "понеделник", "сряда", "вторник", "четвъртък", "петък", "събота"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006 г.", "2 January 2006 г.", "2.01.2006 г.", "2.01.06 г."},
		time:     [4]string{"15:04:05 ч. MST -0700", "15:04:05 ч. MST", "15:04:05 ч.", "15:04 ч."},
		datetime: [4]string{"Monday, 2 January 2006 г., 15:04:05 ч. MST -0700", "2 January 2006 г., 15:04:05 ч. MST", "2.01.2006 г., 15:04:05 ч.", "2.01.06 г., 15:04 ч."},
	},
}
var l_bn = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"জানু", "ফেব", "মার্চ", "এপ্রিল", "মে", "জুন", "জুলাই", "আগস্ট", "সেপ্টেম্বর", "অক্টোবর", "নভেম্বর", "ডিসেম্বর"},
		long:  [12]string{"জানুয়ারী", "ফেব্রুয়ারী", "মার্চ", "এপ্রিল", "মে", "জুন", "জুলাই", "আগস্ট", "সেপ্টেম্বর", "অক্টোবর", "নভেম্বর", "ডিসেম্বর"},
	},
	days: days{
		short: [7]string{"রবি", "সোম", "বুধ", "মঙ্গল", "বৃহস্পতি", "শুক্র", "শনি"},
		long:  [7]string{"রবিবার", "সোমবার", "বুধবার", "মঙ্গলবার", "বৃহস্পতিবার", "শুক্রবার", "শনিবার"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January, 2006", "2 January, 2006", "2 Jan, 2006", "2/1/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January, 2006 3:04:05 PM MST -0700", "2 January, 2006 3:04:05 PM MST", "2 Jan, 2006 3:04:05 PM", "2/1/06 3:04 PM"},
	},
}
var l_bn_IN = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"জানু", "ফেব", "মার্চ", "এপ্রিল", "মে", "জুন", "জুলাই", "আগস্ট", "সেপ্টেম্বর", "অক্টোবর", "নভেম্বর", "ডিসেম্বর"},
		long:  [12]string{"জানুয়ারী", "ফেব্রুয়ারী", "মার্চ", "এপ্রিল", "মে", "জুন", "জুলাই", "আগস্ট", "সেপ্টেম্বর", "অক্টোবর", "নভেম্বর", "ডিসেম্বর"},
	},
	days: days{
		short: [7]string{"রবি", "সোম", "বুধ", "মঙ্গল", "বৃহস্পতি", "শুক্র", "শনি"},
		long:  [7]string{"রবিবার", "সোমবার", "বুধবার", "মঙ্গলবার", "বৃহস্পতিবার", "শুক্রবার", "শনিবার"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January, 2006", "2 January, 2006", "2 Jan, 2006", "2/1/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January, 2006 3:04:05 PM MST -0700", "2 January, 2006 3:04:05 PM MST", "2 Jan, 2006 3:04:05 PM", "2/1/06 3:04 PM"},
	},
}
var l_bs = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan", "feb", "mar", "apr", "maj", "jun", "jul", "aug", "sep", "okt", "nov", "dec"},
		long:  [12]string{"januar", "februar", "mart", "april", "maj", "juni", "juli", "august", "septembar", "oktobar", "novembar", "decembar"},
	},
	days: days{
		short: [7]string{"ned", "pon", "sri", "uto", "čet", "pet", "sub"},
		long:  [7]string{"nedjelja", "ponedjeljak", "srijeda", "utorak", "četvrtak", "petak", "subota"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2. January 2006.", "2. January 2006.", "2. Jan 2006.", "2. 1. 2006."},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2. January 2006. u 15:04:05 MST -0700", "2. January 2006. u 15:04:05 MST", "2. Jan 2006. 15:04:05", "2. 1. 2006. 15:04"},
	},
}
var l_ca = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"de gen.", "de febr.", "de març", "d’abr.", "de maig", "de juny", "de jul.", "d’ag.", "de set.", "d’oct.", "de nov.", "de des."},
		long:  [12]string{"de gener", "de febrer", "de març", "d’abril", "de maig", "de juny", "de juliol", "d’agost", "de setembre", "d’octubre", "de novembre", "de desembre"},
	},
	days: days{
		short: [7]string{"dg.", "dl.", "dc.", "dt.", "dj.", "dv.", "ds."},
		long:  [7]string{"diumenge", "dilluns", "dimecres", "dimarts", "dijous", "divendres", "dissabte"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January de 2006", "2 January de 2006", "2 Jan 2006", "2/1/06"},
		time:     [4]string{"15:04:05 (MST -0700)", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January de 2006, a les 15:04:05 (MST -0700)", "2 January de 2006, a les 15:04:05 MST", "2 Jan 2006, 15:04:05", "2/1/06 15:04"},
	},
}
var l_ca_AD = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"de gen.", "de febr.", "de març", "d’abr.", "de maig", "de juny", "de jul.", "d’ag.", "de set.", "d’oct.", "de nov.", "de des."},
		long:  [12]string{"de gener", "de febrer", "de març", "d’abril", "de maig", "de juny", "de juliol", "d’agost", "de setembre", "d’octubre", "de novembre", "de desembre"},
	},
	days: days{
		short: [7]string{"dg.", "dl.", "dc.", "dt.", "dj.", "dv.", "ds."},
		long:  [7]string{"diumenge", "dilluns", "dimecres", "dimarts", "dijous", "divendres", "dissabte"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January de 2006", "2 January de 2006", "2 Jan 2006", "2/1/06"},
		time:     [4]string{"15:04:05 (MST -0700)", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January de 2006, a les 15:04:05 (MST -0700)", "2 January de 2006, a les 15:04:05 MST", "2 Jan 2006, 15:04:05", "2/1/06 15:04"},
	},
}
var l_cs = &localize{
	am:              "dop.",
	pm:              "odp.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"led", "úno", "bře", "dub", "kvě", "čvn", "čvc", "srp", "zář", "říj", "lis", "pro"},
		long:  [12]string{"ledna", "února", "března", "dubna", "května", "června", "července", "srpna", "září", "října", "listopadu", "prosince"},
	},
	days: days{
		short: [7]string{"ne", "po", "st", "út", "čt", "pá", "so"},
		long:  [7]string{"neděle", "pondělí", "středa", "úterý", "čtvrtek", "pátek", "sobota"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2. January 2006", "2. January 2006", "2. 1. 2006", "02.01.06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 2. January 2006 15:04:05 MST -0700", "2. January 2006 15:04:05 MST", "2. 1. 2006 15:04:05", "02.01.06 15:04"},
	},
}
var l_cy = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Ion", "Chwef", "Maw", "Ebr", "Mai", "Meh", "Gorff", "Awst", "Medi", "Hyd", "Tach", "Rhag"},
		long:  [12]string{"Ionawr", "Chwefror", "Mawrth", "Ebrill", "Mai", "Mehefin", "Gorffennaf", "Awst", "Medi", "Hydref", "Tachwedd", "Rhagfyr"},
	},
	days: days{
		short: [7]string{"Sul", "Llun", "Mer", "Maw", "Iau", "Gwen", "Sad"},
		long:  [7]string{"Dydd Sul", "Dydd Llun", "Dydd Mercher", "Dydd Mawrth", "Dydd Iau", "Dydd Gwener", "Dydd Sadwrn"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006 am 15:04:05 MST -0700", "2 January 2006 am 15:04:05 MST", "2 Jan 2006 15:04:05", "02/01/06 15:04"},
	},
}
var l_da = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan.", "feb.", "mar.", "apr.", "maj", "jun.", "jul.", "aug.", "sep.", "okt.", "nov.", "dec."},
		long:  [12]string{"januar", "februar", "marts", "april", "maj", "juni", "juli", "august", "september", "oktober", "november", "december"},
	},
	days: days{
		short: [7]string{"søn.", "man.", "ons.", "tir.", "tor.", "fre.", "lør."},
		long:  [7]string{"søndag", "mandag", "onsdag", "tirsdag", "torsdag", "fredag", "lørdag"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday den 2. January 2006", "2. January 2006", "2. Jan 2006", "02.01.2006"},
		time:     [4]string{"15.04.05 MST -0700", "15.04.05 MST", "15.04.05", "15.04"},
		datetime: [4]string{"Monday den 2. January 2006 kl. 15.04.05 MST -0700", "2. January 2006 kl. 15.04.05 MST", "2. Jan 2006 15.04.05", "02.01.2006 15.04"},
	},
}
var l_de = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan.", "Feb.", "März", "Apr.", "Mai", "Juni", "Juli", "Aug.", "Sept.", "Okt.", "Nov.", "Dez."},
		long:  [12]string{"Januar", "Februar", "März", "April", "Mai", "Juni", "Juli", "August", "September", "Oktober", "November", "Dezember"},
	},
	days: days{
		short: [7]string{"So.", "Mo.", "Mi.", "Di.", "Do.", "Fr.", "Sa."},
		long:  [7]string{"Sonntag", "Montag", "Mittwoch", "Dienstag", "Donnerstag", "Freitag", "Samstag"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2. January 2006", "2. January 2006", "02.01.2006", "02.01.06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2. January 2006 um 15:04:05 MST -0700", "2. January 2006 um 15:04:05 MST", "02.01.2006, 15:04:05", "02.01.06, 15:04"},
	},
}
var l_de_AT = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jän.", "Feb.", "März", "Apr.", "Mai", "Juni", "Juli", "Aug.", "Sep.", "Okt.", "Nov.", "Dez."},
		long:  [12]string{"Jänner", "Februar", "März", "April", "Mai", "Juni", "Juli", "August", "September", "Oktober", "November", "Dezember"},
	},
	days: days{
		short: [7]string{"So.", "Mo.", "Mi.", "Di.", "Do.", "Fr.", "Sa."},
		long:  [7]string{"Sonntag", "Montag", "Mittwoch", "Dienstag", "Donnerstag", "Freitag", "Samstag"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2. January 2006", "2. January 2006", "02.01.2006", "02.01.06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2. January 2006 um 15:04:05 MST -0700", "2. January 2006 um 15:04:05 MST", "02.01.2006, 15:04:05", "02.01.06, 15:04"},
	},
}
var l_el = &localize{
	am:              "π.μ.",
	pm:              "μ.μ.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Ιαν", "Φεβ", "Μαρ", "Απρ", "Μαΐ", "Ιουν", "Ιουλ", "Αυγ", "Σεπ", "Οκτ", "Νοε", "Δεκ"},
		long:  [12]string{"Ιανουαρίου", "Φεβρουαρίου", "Μαρτίου", "Απριλίου", "Μαΐου", "Ιουνίου", "Ιουλίου", "Αυγούστου", "Σεπτεμβρίου", "Οκτωβρίου", "Νοεμβρίου", "Δεκεμβρίου"},
	},
	days: days{
		short: [7]string{"Κυρ", "Δευ", "Τετ", "Τρί", "Πέμ", "Παρ", "Σάβ"},
		long:  [7]string{"Κυριακή", "Δευτέρα", "Τετάρτη", "Τρίτη", "Πέμπτη", "Παρασκευή", "Σάββατο"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "2/1/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January 2006 - 3:04:05 PM MST -0700", "2 January 2006 - 3:04:05 PM MST", "2 Jan 2006, 3:04:05 PM", "2/1/06, 3:04 PM"},
	},
}
var l_en = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, January 2, 2006", "January 2, 2006", "Jan 2, 2006", "1/2/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, January 2, 2006 at 3:04:05 PM MST -0700", "January 2, 2006 at 3:04:05 PM MST", "Jan 2, 2006, 3:04:05 PM", "1/2/06, 3:04 PM"},
	},
}
var l_en_001 = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January 2006 at 3:04:05 PM MST -0700", "2 January 2006 at 3:04:05 PM MST", "2 Jan 2006, 3:04:05 PM", "02/01/2006, 3:04 PM"},
	},
}
var l_en_150 = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006 at 15:04:05 MST -0700", "2 January 2006 at 15:04:05 MST", "2 Jan 2006, 15:04:05", "02/01/2006, 15:04"},
	},
}
var l_en_AE = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        6,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January 2006 at 3:04:05 PM MST -0700", "2 January 2006 at 3:04:05 PM MST", "2 Jan 2006, 3:04:05 PM", "02/01/2006, 3:04 PM"},
	},
}
var l_en_AG = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January 2006 at 3:04:05 PM MST -0700", "2 January 2006 at 3:04:05 PM MST", "2 Jan 2006, 3:04:05 PM", "02/01/2006, 3:04 PM"},
	},
}
var l_en_AS = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, January 2, 2006", "January 2, 2006", "Jan 2, 2006", "1/2/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, January 2, 2006 at 3:04:05 PM MST -0700", "January 2, 2006 at 3:04:05 PM MST", "Jan 2, 2006, 3:04:05 PM", "1/2/06, 3:04 PM"},
	},
}
var l_en_AU = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sept", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "2/1/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January 2006 at 3:04:05 PM MST -0700", "2 January 2006 at 3:04:05 PM MST", "2 Jan 2006, 3:04:05 PM", "2/1/06, 3:04 PM"},
	},
}
var l_en_BE = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "02 Jan 2006", "02/01/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006 at 15:04:05 MST -0700", "2 January 2006 at 15:04:05 MST", "02 Jan 2006, 15:04:05", "02/01/06, 15:04"},
	},
}
var l_en_BI = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, January 2, 2006", "January 2, 2006", "Jan 2, 2006", "1/2/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, January 2, 2006 at 15:04:05 MST -0700", "January 2, 2006 at 15:04:05 MST", "Jan 2, 2006, 15:04:05", "1/2/06, 15:04"},
	},
}
var l_en_BW = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 02 January 2006", "02 January 2006", "02 Jan 2006", "02/01/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 02 January 2006 at 15:04:05 MST -0700", "02 January 2006 at 15:04:05 MST", "02 Jan 2006, 15:04:05", "02/01/06, 15:04"},
	},
}
var l_en_BZ = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 02 January 2006", "02 January 2006", "02-Jan-2006", "02/01/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 02 January 2006 at 15:04:05 MST -0700", "02 January 2006 at 15:04:05 MST", "02-Jan-2006, 15:04:05", "02/01/06, 15:04"},
	},
}
var l_en_CA = &localize{
	am:              "a.m.",
	pm:              "p.m.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan.", "Feb.", "Mar.", "Apr.", "May", "Jun.", "Jul.", "Aug.", "Sep.", "Oct.", "Nov.", "Dec."},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun.", "Mon.", "Wed.", "Tue.", "Thu.", "Fri.", "Sat."},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, January 2, 2006", "January 2, 2006", "Jan 2, 2006", "2006-01-02"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, January 2, 2006 at 3:04:05 PM MST -0700", "January 2, 2006 at 3:04:05 PM MST", "Jan 2, 2006, 3:04:05 PM", "2006-01-02, 3:04 PM"},
	},
}
var l_en_DK = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"15.04.05 MST -0700", "15.04.05 MST", "15.04.05", "15.04"},
		datetime: [4]string{"Monday, 2 January 2006 at 15.04.05 MST -0700", "2 January 2006 at 15.04.05 MST", "2 Jan 2006, 15.04.05", "02/01/2006, 15.04"},
	},
}
var l_en_GB = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sept", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006 at 15:04:05 MST -0700", "2 January 2006 at 15:04:05 MST", "2 Jan 2006, 15:04:05", "02/01/2006, 15:04"},
	},
}
var l_en_HK = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "2/1/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January 2006 at 3:04:05 PM MST -0700", "2 January 2006 at 3:04:05 PM MST", "2 Jan 2006, 3:04:05 PM", "2/1/2006, 3:04 PM"},
	},
}
var l_en_IE = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 2 January 2006 at 15:04:05 MST -0700", "2 January 2006 at 15:04:05 MST", "2 Jan 2006, 15:04:05", "02/01/2006, 15:04"},
	},
}
var l_en_IL = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006 at 15:04:05 MST -0700", "2 January 2006 at 15:04:05 MST", "2 Jan 2006, 15:04:05", "02/01/2006, 15:04"},
	},
}
var l_en_IN = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January, 2006", "2 January 2006", "02-Jan-2006", "02/01/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January, 2006 at 3:04:05 PM MST -0700", "2 January 2006 at 3:04:05 PM MST", "02-Jan-2006, 3:04:05 PM", "02/01/06, 3:04 PM"},
	},
}
var l_en_MT = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "02 January 2006", "02 Jan 2006", "02/01/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006 at 15:04:05 MST -0700", "02 January 2006 at 15:04:05 MST", "02 Jan 2006, 15:04:05", "02/01/2006, 15:04"},
	},
}
var l_en_NZ = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2/01/2006", "2/01/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January 2006 at 3:04:05 PM MST -0700", "2 January 2006 at 3:04:05 PM MST", "2/01/2006, 3:04:05 PM", "2/01/06, 3:04 PM"},
	},
}
var l_en_PK = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "02-Jan-2006", "02/01/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January 2006 at 3:04:05 PM MST -0700", "2 January 2006 at 3:04:05 PM MST", "02-Jan-2006, 3:04:05 PM", "02/01/2006, 3:04 PM"},
	},
}
var l_en_SD = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        6,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January 2006 at 3:04:05 PM MST -0700", "2 January 2006 at 3:04:05 PM MST", "2 Jan 2006, 3:04:05 PM", "02/01/2006, 3:04 PM"},
	},
}
var l_en_SE = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "2006-01-02"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006 at 15:04:05 MST -0700", "2 January 2006 at 15:04:05 MST", "2 Jan 2006, 15:04:05", "2006-01-02, 15:04"},
	},
}
var l_en_SG = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "2/1/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January 2006 at 3:04:05 PM MST -0700", "2 January 2006 at 3:04:05 PM MST", "2 Jan 2006, 3:04:05 PM", "2/1/06, 3:04 PM"},
	},
}
var l_en_ZA = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 02 January 2006", "02 January 2006", "02 Jan 2006", "2006/01/02"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 02 January 2006 at 15:04:05 MST -0700", "02 January 2006 at 15:04:05 MST", "02 Jan 2006, 15:04:05", "2006/01/02, 15:04"},
	},
}
var l_en_ZW = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		long:  [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sunday", "Monday", "Wednesday", "Tuesday", "Thursday", "Friday", "Saturday"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 02 January 2006", "02 January 2006", "02 Jan,2006", "2/1/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 02 January 2006 at 15:04:05 MST -0700", "02 January 2006 at 15:04:05 MST", "02 Jan,2006, 15:04:05", "2/1/2006, 15:04"},
	},
}
var l_es = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ene", "feb", "mar", "abr", "may", "jun", "jul", "ago", "sept", "oct", "nov", "dic"},
		long:  [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mié", "mar", "jue", "vie", "sáb"},
		long:  [7]string{"domingo", "lunes", "miércoles", "martes", "jueves", "viernes", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "2 Jan 2006", "2/1/06"},
		time:     [4]string{"15:04:05 (MST -0700)", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 de January de 2006, 15:04:05 (MST -0700)", "2 de January de 2006, 15:04:05 MST", "2 Jan 2006 15:04:05", "2/1/06 15:04"},
	},
}
var l_es_419 = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ene.", "feb.", "mar.", "abr.", "may.", "jun.", "jul.", "ago.", "sep.", "oct.", "nov.", "dic."},
		long:  [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mié", "mar", "jue", "vie", "sáb"},
		long:  [7]string{"domingo", "lunes", "miércoles", "martes", "jueves", "viernes", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "2 Jan 2006", "2/1/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 de January de 2006, 15:04:05 MST -0700", "2 de January de 2006, 15:04:05 MST", "2 Jan 2006 15:04:05", "2/1/06 15:04"},
	},
}
var l_es_BO = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ene.", "feb.", "mar.", "abr.", "may.", "jun.", "jul.", "ago.", "sep.", "oct.", "nov.", "dic."},
		long:  [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mié", "mar", "jue", "vie", "sáb"},
		long:  [7]string{"domingo", "lunes", "miércoles", "martes", "jueves", "viernes", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "2 Jan de 2006", "2/1/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 de January de 2006, 15:04:05 MST -0700", "2 de January de 2006, 15:04:05 MST", "2 Jan de 2006 15:04:05", "2/1/06 15:04"},
	},
}
var l_es_BR = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ene.", "feb.", "mar.", "abr.", "may.", "jun.", "jul.", "ago.", "sep.", "oct.", "nov.", "dic."},
		long:  [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mié", "mar", "jue", "vie", "sáb"},
		long:  [7]string{"domingo", "lunes", "miércoles", "martes", "jueves", "viernes", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "2 Jan 2006", "2/1/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 de January de 2006, 15:04:05 MST -0700", "2 de January de 2006, 15:04:05 MST", "2 Jan 2006 15:04:05", "2/1/06 15:04"},
	},
}
var l_es_CL = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ene.", "feb.", "mar.", "abr.", "may.", "jun.", "jul.", "ago.", "sep.", "oct.", "nov.", "dic."},
		long:  [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mié", "mar", "jue", "vie", "sáb"},
		long:  [7]string{"domingo", "lunes", "miércoles", "martes", "jueves", "viernes", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "02-01-2006", "02-01-06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 de January de 2006, 15:04:05 MST -0700", "2 de January de 2006, 15:04:05 MST", "02-01-2006 15:04:05", "02-01-06 15:04"},
	},
}
var l_es_CO = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ene.", "feb.", "mar.", "abr.", "may.", "jun.", "jul.", "ago.", "sep.", "oct.", "nov.", "dic."},
		long:  [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mié", "mar", "jue", "vie", "sáb"},
		long:  [7]string{"domingo", "lunes", "miércoles", "martes", "jueves", "viernes", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "2/01/2006", "2/01/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 de January de 2006, 3:04:05 PM MST -0700", "2 de January de 2006, 3:04:05 PM MST", "2/01/2006, 3:04:05 PM", "2/01/06, 3:04 PM"},
	},
}
var l_es_DO = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ene.", "feb.", "mar.", "abr.", "may.", "jun.", "jul.", "ago.", "sep.", "oct.", "nov.", "dic."},
		long:  [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mié", "mar", "jue", "vie", "sáb"},
		long:  [7]string{"domingo", "lunes", "miércoles", "martes", "jueves", "viernes", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "2 Jan 2006", "2/1/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 de January de 2006, 3:04:05 PM MST -0700", "2 de January de 2006, 3:04:05 PM MST", "2 Jan 2006 3:04:05 PM", "2/1/06 3:04 PM"},
	},
}
var l_es_GT = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ene.", "feb.", "mar.", "abr.", "may.", "jun.", "jul.", "ago.", "sep.", "oct.", "nov.", "dic."},
		long:  [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mié", "mar", "jue", "vie", "sáb"},
		long:  [7]string{"domingo", "lunes", "miércoles", "martes", "jueves", "viernes", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "2/01/2006", "2/01/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 de January de 2006, 15:04:05 MST -0700", "2 de January de 2006, 15:04:05 MST", "2/01/2006 15:04:05", "2/01/06 15:04"},
	},
}
var l_es_HN = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ene.", "feb.", "mar.", "abr.", "may.", "jun.", "jul.", "ago.", "sep.", "oct.", "nov.", "dic."},
		long:  [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mié", "mar", "jue", "vie", "sáb"},
		long:  [7]string{"domingo", "lunes", "miércoles", "martes", "jueves", "viernes", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 02 de January de 2006", "02 de January de 2006", "2 Jan 2006", "2/1/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 02 de January de 2006, 15:04:05 MST -0700", "02 de January de 2006, 15:04:05 MST", "2 Jan 2006 15:04:05", "2/1/06 15:04"},
	},
}
var l_es_MX = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ene.", "feb.", "mar.", "abr.", "may.", "jun.", "jul.", "ago.", "sep.", "oct.", "nov.", "dic."},
		long:  [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mié", "mar", "jue", "vie", "sáb"},
		long:  [7]string{"domingo", "lunes", "miércoles", "martes", "jueves", "viernes", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "2 Jan 2006", "02/01/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 de January de 2006, 15:04:05 MST -0700", "2 de January de 2006, 15:04:05 MST", "2 Jan 2006 15:04:05", "02/01/06 15:04"},
	},
}
var l_es_PA = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ene.", "feb.", "mar.", "abr.", "may.", "jun.", "jul.", "ago.", "sep.", "oct.", "nov.", "dic."},
		long:  [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mié", "mar", "jue", "vie", "sáb"},
		long:  [7]string{"domingo", "lunes", "miércoles", "martes", "jueves", "viernes", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "01/02/2006", "01/02/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 de January de 2006, 3:04:05 PM MST -0700", "2 de January de 2006, 3:04:05 PM MST", "01/02/2006 3:04:05 PM", "01/02/06 3:04 PM"},
	},
}
var l_es_PE = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ene.", "feb.", "mar.", "abr.", "may.", "jun.", "jul.", "ago.", "set.", "oct.", "nov.", "dic."},
		long:  [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "setiembre", "octubre", "noviembre", "diciembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mié", "mar", "jue", "vie", "sáb"},
		long:  [7]string{"domingo", "lunes", "miércoles", "martes", "jueves", "viernes", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "2 Jan 2006", "2/01/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 de January de 2006, 15:04:05 MST -0700", "2 de January de 2006, 15:04:05 MST", "2 Jan 2006 15:04:05", "2/01/06 15:04"},
	},
}
var l_es_PH = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ene", "feb", "mar", "abr", "may", "jun", "jul", "ago", "sept", "oct", "nov", "dic"},
		long:  [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mié", "mar", "jue", "vie", "sáb"},
		long:  [7]string{"domingo", "lunes", "miércoles", "martes", "jueves", "viernes", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "2 Jan 2006", "2/1/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 de January de 2006, 3:04:05 PM MST -0700", "2 de January de 2006, 3:04:05 PM MST", "2 Jan 2006 3:04:05 PM", "2/1/06 3:04 PM"},
	},
}
var l_es_PY = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ene.", "feb.", "mar.", "abr.", "may.", "jun.", "jul.", "ago.", "sept.", "oct.", "nov.", "dic."},
		long:  [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mié", "mar", "jue", "vie", "sáb"},
		long:  [7]string{"domingo", "lunes", "miércoles", "martes", "jueves", "viernes", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "2 Jan 2006", "2/1/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 de January de 2006, 15:04:05 MST -0700", "2 de January de 2006, 15:04:05 MST", "2 Jan 2006 15:04:05", "2/1/06 15:04"},
	},
}
var l_es_US = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ene.", "feb.", "mar.", "abr.", "may.", "jun.", "jul.", "ago.", "sep.", "oct.", "nov.", "dic."},
		long:  [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mié", "mar", "jue", "vie", "sáb"},
		long:  [7]string{"domingo", "lunes", "miércoles", "martes", "jueves", "viernes", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "2 Jan 2006", "2/1/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 de January de 2006, 3:04:05 PM MST -0700", "2 de January de 2006, 3:04:05 PM MST", "2 Jan 2006 3:04:05 PM", "2/1/2006 3:04 PM"},
	},
}
var l_es_UY = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ene.", "feb.", "mar.", "abr.", "may.", "jun.", "jul.", "ago.", "set.", "oct.", "nov.", "dic."},
		long:  [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "setiembre", "octubre", "noviembre", "diciembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mié", "mar", "jue", "vie", "sáb"},
		long:  [7]string{"domingo", "lunes", "miércoles", "martes", "jueves", "viernes", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "2 Jan 2006", "2/1/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 de January de 2006, 15:04:05 MST -0700", "2 de January de 2006, 15:04:05 MST", "2 Jan 2006 15:04:05", "2/1/06 15:04"},
	},
}
var l_es_VE = &localize{
	am:              "a.\u00a0m.",
	pm:              "p.\u00a0m.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ene.", "feb.", "mar.", "abr.", "may.", "jun.", "jul.", "ago.", "sept.", "oct.", "nov.", "dic."},
		long:  [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mié", "mar", "jue", "vie", "sáb"},
		long:  [7]string{"domingo", "lunes", "miércoles", "martes", "jueves", "viernes", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "2 Jan 2006", "2/1/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 de January de 2006, 3:04:05 PM MST -0700", "2 de January de 2006, 3:04:05 PM MST", "2 Jan 2006 3:04:05 PM", "2/1/06 3:04 PM"},
	},
}
var l_et = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jaan", "veebr", "märts", "apr", "mai", "juuni", "juuli", "aug", "sept", "okt", "nov", "dets"},
		long:  [12]string{"jaanuar", "veebruar", "märts", "aprill", "mai", "juuni", "juuli", "august", "september", "oktoober", "november", "detsember"},
	},
	days: days{
		short: [7]string{"P", "E", "K", "T", "N", "R", "L"},
		long:  [7]string{"pühapäev", "esmaspäev", "kolmapäev", "teisipäev", "neljapäev", "reede", "laupäev"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2. January 2006", "2. January 2006", "2. Jan 2006", "02.01.06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2. January 2006 15:04:05 MST -0700", "2. January 2006 15:04:05 MST", "2. Jan 2006 15:04:05", "02.01.06 15:04"},
	},
}
var l_eu = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"urt.", "ots.", "mar.", "api.", "mai.", "eka.", "uzt.", "abu.", "ira.", "urr.", "aza.", "abe."},
		long:  [12]string{"urtarrilak", "otsailak", "martxoak", "apirilak", "maiatzak", "ekainak", "uztailak", "abuztuak", "irailak", "urriak", "azaroak", "abenduak"},
	},
	days: days{
		short: [7]string{"ig.", "al.", "az.", "ar.", "og.", "or.", "lr."},
		long:  [7]string{"igandea", "astelehena", "asteazkena", "asteartea", "osteguna", "ostirala", "larunbata"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006(e)ko renJanuary 2(a), Monday", "2006(e)ko renJanuary 2(a)", "2006(e)ko Jan 2(a)", "06/1/2"},
		time:     [4]string{"15:04:05 (MST -0700)", "15:04:05 (MST)", "15:04:05", "15:04"},
		datetime: [4]string{"2006(e)ko renJanuary 2(a), Monday 15:04:05 (MST -0700)", "2006(e)ko renJanuary 2(a) 15:04:05 (MST)", "2006(e)ko Jan 2(a) 15:04:05", "06/1/2 15:04"},
	},
}
var l_fa = &localize{
	am:              "ق.ظ.",
	pm:              "ب.ظ.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ژانویهٔ", "فوریهٔ", "مارس", "آوریل", "مهٔ", "ژوئن", "ژوئیهٔ", "اوت", "سپتامبر", "اکتبر", "نوامبر", "دسامبر"},
		long:  [12]string{"ژانویهٔ", "فوریهٔ", "مارس", "آوریل", "مهٔ", "ژوئن", "ژوئیهٔ", "اوت", "سپتامبر", "اکتبر", "نوامبر", "دسامبر"},
	},
	days: days{
		short: [7]string{"یکشنبه", "دوشنبه", "چهارشنبه", "سه\u200cشنبه", "پنجشنبه", "جمعه", "شنبه"},
		long:  [7]string{"یکشنبه", "دوشنبه", "چهارشنبه", "سه\u200cشنبه", "پنجشنبه", "جمعه", "شنبه"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2 January 2006", "2 January 2006", "2 Jan 2006", "2006/1/2"},
		time:     [4]string{"15:04:05 (MST -0700)", "15:04:05 (MST)", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 2 January 2006، ساعت 15:04:05 (MST -0700)", "2 January 2006، ساعت 15:04:05 (MST)", "2 Jan 2006،\u200f 15:04:05", "2006/1/2،\u200f 15:04"},
	},
}
var l_fa_AF = &localize{
	am:              "ق.ظ.",
	pm:              "ب.ظ.",
	firstDay:        6,
	twentyFourHours: false,
	months: months{
		short: [12]string{"جنو", "فبروری", "مارچ", "اپریل", "می", "جون", "جول", "اگست", "سپتمبر", "اکتوبر", "نومبر", "دسم"},
		long:  [12]string{"جنوری", "فبروری", "مارچ", "اپریل", "می", "جون", "جولای", "اگست", "سپتمبر", "اکتوبر", "نومبر", "دسمبر"},
	},
	days: days{
		short: [7]string{"یکشنبه", "دوشنبه", "چهارشنبه", "سه\u200cشنبه", "پنجشنبه", "جمعه", "شنبه"},
		long:  [7]string{"یکشنبه", "دوشنبه", "چهارشنبه", "سه\u200cشنبه", "پنجشنبه", "جمعه", "شنبه"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2 January 2006", "2 January 2006", "2 Jan 2006", "2006/1/2"},
		time:     [4]string{"15:04:05 (MST -0700)", "15:04:05 (MST)", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 2 January 2006، ساعت 15:04:05 (MST -0700)", "2 January 2006، ساعت 15:04:05 (MST)", "2 Jan 2006،\u200f 15:04:05", "2006/1/2،\u200f 15:04"},
	},
}
var l_fi = &localize{
	am:              "ap.",
	pm:              "ip.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"tammik.", "helmik.", "maalisk.", "huhtik.", "toukok.", "kesäk.", "heinäk.", "elok.", "syysk.", "lokak.", "marrask.", "jouluk."},
		long:  [12]string{"tammikuuta", "helmikuuta", "maaliskuuta", "huhtikuuta", "toukokuuta", "kesäkuuta", "heinäkuuta", "elokuuta", "syyskuuta", "lokakuuta", "marraskuuta", "joulukuuta"},
	},
	days: days{
		short: [7]string{"su", "ma", "ke", "ti", "to", "pe", "la"},
		long:  [7]string{"sunnuntaina", "maanantaina", "keskiviikkona", "tiistaina", "torstaina", "perjantaina", "lauantaina"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"cccc 2. January 2006", "2. January 2006", "2.1.2006", "2.1.2006"},
		time:     [4]string{"15.04.05 MST -0700", "15.04.05 MST", "15.04.05", "15.04"},
		datetime: [4]string{"cccc 2. January 2006 klo 15.04.05 MST -0700", "2. January 2006 klo 15.04.05 MST", "2.1.2006 klo 15.04.05", "2.1.2006 15.04"},
	},
}
var l_fil = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Ene", "Peb", "Mar", "Abr", "May", "Hun", "Hul", "Ago", "Set", "Okt", "Nob", "Dis"},
		long:  [12]string{"Enero", "Pebrero", "Marso", "Abril", "Mayo", "Hunyo", "Hulyo", "Agosto", "Setyembre", "Oktubre", "Nobyembre", "Disyembre"},
	},
	days: days{
		short: [7]string{"Lin", "Lun", "Miy", "Mar", "Huw", "Biy", "Sab"},
		long:  [7]string{"Linggo", "Lunes", "Miyerkules", "Martes", "Huwebes", "Biyernes", "Sabado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, January 2, 2006", "January 2, 2006", "Jan 2, 2006", "1/2/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, January 2, 2006 nang 3:04:05 PM MST -0700", "January 2, 2006 nang 3:04:05 PM MST", "Jan 2, 2006, 3:04:05 PM", "1/2/06, 3:04 PM"},
	},
}
var l_fr = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"janv.", "févr.", "mars", "avr.", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."},
		long:  [12]string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"},
	},
	days: days{
		short: [7]string{"dim.", "lun.", "mer.", "mar.", "jeu.", "ven.", "sam."},
		long:  [7]string{"dimanche", "lundi", "mercredi", "mardi", "jeudi", "vendredi", "samedi"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 2 January 2006 à 15:04:05 MST -0700", "2 January 2006 à 15:04:05 MST", "2 Jan 2006, 15:04:05", "02/01/2006 15:04"},
	},
}
var l_fr_BE = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"janv.", "févr.", "mars", "avr.", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."},
		long:  [12]string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"},
	},
	days: days{
		short: [7]string{"dim.", "lun.", "mer.", "mar.", "jeu.", "ven.", "sam."},
		long:  [7]string{"dimanche", "lundi", "mercredi", "mardi", "jeudi", "vendredi", "samedi"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2 January 2006", "2 January 2006", "2 Jan 2006", "2/01/06"},
		time:     [4]string{"15 h 04 min 05 s MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 2 January 2006 à 15 h 04 min 05 s MST -0700", "2 January 2006 à 15:04:05 MST", "2 Jan 2006, 15:04:05", "2/01/06 15:04"},
	},
}
var l_fr_CA = &localize{
	am:              "a.m.",
	pm:              "p.m.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"janv.", "févr.", "mars", "avr.", "mai", "juin", "juill.", "août", "sept.", "oct.", "nov.", "déc."},
		long:  [12]string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"},
	},
	days: days{
		short: [7]string{"dim.", "lun.", "mer.", "mar.", "jeu.", "ven.", "sam."},
		long:  [7]string{"dimanche", "lundi", "mercredi", "mardi", "jeudi", "vendredi", "samedi"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2 January 2006", "2 January 2006", "2 Jan 2006", "2006-01-02"},
		time:     [4]string{"15 h 04 min 05 s MST -0700", "15 h 04 min 05 s MST", "15 h 04 min 05 s", "15 h 04"},
		datetime: [4]string{"Monday 2 January 2006 à 15 h 04 min 05 s MST -0700", "2 January 2006 à 15 h 04 min 05 s MST", "2 Jan 2006, 15 h 04 min 05 s", "2006-01-02 15 h 04"},
	},
}
var l_fr_CH = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"janv.", "févr.", "mars", "avr.", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."},
		long:  [12]string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"},
	},
	days: days{
		short: [7]string{"dim.", "lun.", "mer.", "mar.", "jeu.", "ven.", "sam."},
		long:  [7]string{"dimanche", "lundi", "mercredi", "mardi", "jeudi", "vendredi", "samedi"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "02.01.06"},
		time:     [4]string{"15.04:05 h MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006 à 15.04:05 h MST -0700", "2 January 2006 à 15:04:05 MST", "2 Jan 2006, 15:04:05", "02.01.06 15:04"},
	},
}
var l_fr_CM = &localize{
	am:              "mat.",
	pm:              "soir",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"janv.", "févr.", "mars", "avr.", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."},
		long:  [12]string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"},
	},
	days: days{
		short: [7]string{"dim.", "lun.", "mer.", "mar.", "jeu.", "ven.", "sam."},
		long:  [7]string{"dimanche", "lundi", "mercredi", "mardi", "jeudi", "vendredi", "samedi"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 2 January 2006 à 15:04:05 MST -0700", "2 January 2006 à 15:04:05 MST", "2 Jan 2006, 15:04:05", "02/01/2006 15:04"},
	},
}
var l_fr_DJ = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        6,
	twentyFourHours: false,
	months: months{
		short: [12]string{"janv.", "févr.", "mars", "avr.", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."},
		long:  [12]string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"},
	},
	days: days{
		short: [7]string{"dim.", "lun.", "mer.", "mar.", "jeu.", "ven.", "sam."},
		long:  [7]string{"dimanche", "lundi", "mercredi", "mardi", "jeudi", "vendredi", "samedi"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday 2 January 2006 à 3:04:05 PM MST -0700", "2 January 2006 à 3:04:05 PM MST", "2 Jan 2006, 3:04:05 PM", "02/01/2006 3:04 PM"},
	},
}
var l_fr_MA = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan.", "fév.", "mar.", "avr.", "mai", "jui.", "juil.", "août", "sept.", "oct.", "nov.", "déc."},
		long:  [12]string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"},
	},
	days: days{
		short: [7]string{"dim.", "lun.", "mer.", "mar.", "jeu.", "ven.", "sam."},
		long:  [7]string{"dimanche", "lundi", "mercredi", "mardi", "jeudi", "vendredi", "samedi"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 2 January 2006 à 15:04:05 MST -0700", "2 January 2006 à 15:04:05 MST", "2 Jan 2006, 15:04:05", "02/01/2006 15:04"},
	},
}
var l_fr_ML = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"janv.", "févr.", "mars", "avr.", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."},
		long:  [12]string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"},
	},
	days: days{
		short: [7]string{"dim.", "lun.", "mer.", "mar.", "jeu.", "ven.", "sam."},
		long:  [7]string{"dimanche", "lundi", "mercredi", "mardi", "jeudi", "vendredi", "samedi"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 2 January 2006 à 15:04:05 MST -0700", "2 January 2006 à 15:04:05 MST", "2 Jan 2006, 15:04:05", "02/01/2006, 15:04"},
	},
}
var l_fr_MR = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"janv.", "févr.", "mars", "avr.", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."},
		long:  [12]string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"},
	},
	days: days{
		short: [7]string{"dim.", "lun.", "mer.", "mar.", "jeu.", "ven.", "sam."},
		long:  [7]string{"dimanche", "lundi", "mercredi", "mardi", "jeudi", "vendredi", "samedi"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday 2 January 2006 à 3:04:05 PM MST -0700", "2 January 2006 à 3:04:05 PM MST", "2 Jan 2006, 3:04:05 PM", "02/01/2006 3:04 PM"},
	},
}
var l_ga = &localize{
	am:              "r.n.",
	pm:              "i.n.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Ean", "Feabh", "Márta", "Aib", "Beal", "Meith", "Iúil", "Lún", "MFómh", "DFómh", "Samh", "Noll"},
		long:  [12]string{"Eanáir", "Feabhra", "Márta", "Aibreán", "Bealtaine", "Meitheamh", "Iúil", "Lúnasa", "Meán Fómhair", "Deireadh Fómhair", "Samhain", "Nollaig"},
	},
	days: days{
		short: [7]string{"Domh", "Luan", "Céad", "Máirt", "Déar", "Aoine", "Sath"},
		long:  [7]string{"Dé Domhnaigh", "Dé Luain", "Dé Céadaoin", "Dé Máirt", "Déardaoin", "Dé hAoine", "Dé Sathairn"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 2 January 2006 15:04:05 MST -0700", "2 January 2006 15:04:05 MST", "2 Jan 2006 15:04:05", "02/01/2006 15:04"},
	},
}
var l_gl = &localize{
	am:              "a.m.",
	pm:              "p.m.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"xan.", "feb.", "mar.", "abr.", "maio", "xuño", "xul.", "ago.", "set.", "out.", "nov.", "dec."},
		long:  [12]string{"xaneiro", "febreiro", "marzo", "abril", "maio", "xuño", "xullo", "agosto", "setembro", "outubro", "novembro", "decembro"},
	},
	days: days{
		short: [7]string{"dom.", "luns", "mér.", "mar.", "xov.", "ven.", "sáb."},
		long:  [7]string{"domingo", "luns", "mércores", "martes", "xoves", "venres", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "2 de Jan de 2006", "02/01/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"15:04:05 MST -0700 do Monday, 2 de January de 2006", "15:04:05 MST do 2 de January de 2006", "15:04:05, 2 de Jan de 2006", "15:04, 02/01/06"},
	},
}
var l_gu = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"જાન્યુ", "ફેબ્રુ", "માર્ચ", "એપ્રિલ", "મે", "જૂન", "જુલાઈ", "ઑગસ્ટ", "સપ્ટે", "ઑક્ટો", "નવે", "ડિસે"},
		long:  [12]string{"જાન્યુઆરી", "ફેબ્રુઆરી", "માર્ચ", "એપ્રિલ", "મે", "જૂન", "જુલાઈ", "ઑગસ્ટ", "સપ્ટેમ્બર", "ઑક્ટોબર", "નવેમ્બર", "ડિસેમ્બર"},
	},
	days: days{
		short: [7]string{"રવિ", "સોમ", "બુધ", "મંગળ", "ગુરુ", "શુક્ર", "શનિ"},
		long:  [7]string{"રવિવાર", "સોમવાર", "બુધવાર", "મંગળવાર", "ગુરુવાર", "શુક્રવાર", "શનિવાર"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January, 2006", "2 January, 2006", "2 Jan, 2006", "2/1/06"},
		time:     [4]string{"03:04:05 PM MST -0700", "03:04:05 PM MST", "03:04:05 PM", "03:04 PM"},
		datetime: [4]string{"Monday, 2 January, 2006 એ 03:04:05 PM MST -0700 વાગ્યે", "2 January, 2006 એ 03:04:05 PM MST વાગ્યે", "2 Jan, 2006 03:04:05 PM", "2/1/06 03:04 PM"},
	},
}
var l_he = &localize{
	am:              "לפנה״צ",
	pm:              "אחה״צ",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ינו׳", "פבר׳", "מרץ", "אפר׳", "מאי", "יוני", "יולי", "אוג׳", "ספט׳", "אוק׳", "נוב׳", "דצמ׳"},
		long:  [12]string{"ינואר", "פברואר", "מרץ", "אפריל", "מאי", "יוני", "יולי", "אוגוסט", "ספטמבר", "אוקטובר", "נובמבר", "דצמבר"},
	},
	days: days{
		short: [7]string{"יום א׳", "יום ב׳", "יום ד׳", "יום ג׳", "יום ה׳", "יום ו׳", "שבת"},
		long:  [7]string{"יום ראשון", "יום שני", "יום רביעי", "יום שלישי", "יום חמישי", "יום שישי", "יום שבת"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 בJanuary 2006", "2 בJanuary 2006", "2 בJan 2006", "2.1.2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 בJanuary 2006 בשעה 15:04:05 MST -0700", "2 בJanuary 2006 בשעה 15:04:05 MST", "2 בJan 2006, 15:04:05", "2.1.2006, 15:04"},
	},
}
var l_hi = &localize{
	am:              "am",
	pm:              "pm",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"जन॰", "फ़र॰", "मार्च", "अप्रैल", "मई", "जून", "जुल॰", "अग॰", "सित॰", "अक्तू॰", "नव॰", "दिस॰"},
		long:  [12]string{"जनवरी", "फ़रवरी", "मार्च", "अप्रैल", "मई", "जून", "जुलाई", "अगस्त", "सितंबर", "अक्तूबर", "नवंबर", "दिसंबर"},
	},
	days: days{
		short: [7]string{"रवि", "सोम", "बुध", "मंगल", "गुरु", "शुक्र", "शनि"},
		long:  [7]string{"रविवार", "सोमवार", "बुधवार", "मंगलवार", "गुरुवार", "शुक्रवार", "शनिवार"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "2/1/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January 2006 को 3:04:05 PM MST -0700", "2 January 2006 को 3:04:05 PM MST", "2 Jan 2006, 3:04:05 PM", "2/1/06, 3:04 PM"},
	},
}
var l_hr = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"sij", "velj", "ožu", "tra", "svi", "lip", "srp", "kol", "ruj", "lis", "stu", "pro"},
		long:  [12]string{"siječnja", "veljače", "ožujka", "travnja", "svibnja", "lipnja", "srpnja", "kolovoza", "rujna", "listopada", "studenoga", "prosinca"},
	},
	days: days{
		short: [7]string{"ned", "pon", "sri", "uto", "čet", "pet", "sub"},
		long:  [7]string{"nedjelja", "ponedjeljak", "srijeda", "utorak", "četvrtak", "petak", "subota"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2. January 2006.", "2. January 2006.", "2. Jan 2006.", "02. 01. 2006."},
		time:     [4]string{"15:04:05 (MST -0700)", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2. January 2006. u 15:04:05 (MST -0700)", "2. January 2006. u 15:04:05 MST", "2. Jan 2006. 15:04:05", "02. 01. 2006. 15:04"},
	},
}
var l_hr_BA = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"sij", "velj", "ožu", "tra", "svi", "lip", "srp", "kol", "ruj", "lis", "stu", "pro"},
		long:  [12]string{"siječnja", "veljače", "ožujka", "travnja", "svibnja", "lipnja", "srpnja", "kolovoza", "rujna", "listopada", "studenoga", "prosinca"},
	},
	days: days{
		short: [7]string{"ned", "pon", "sri", "uto", "čet", "pet", "sub"},
		long:  [7]string{"nedjelja", "ponedjeljak", "srijeda", "utorak", "četvrtak", "petak", "subota"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2. January 2006.", "2. January 2006.", "2. Jan 2006.", "2. 1. 06."},
		time:     [4]string{"15:04:05 (MST -0700)", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2. January 2006. u 15:04:05 (MST -0700)", "2. January 2006. u 15:04:05 MST", "2. Jan 2006. 15:04:05", "2. 1. 06. 15:04"},
	},
}
var l_hu = &localize{
	am:              "de.",
	pm:              "du.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan.", "febr.", "márc.", "ápr.", "máj.", "jún.", "júl.", "aug.", "szept.", "okt.", "nov.", "dec."},
		long:  [12]string{"január", "február", "március", "április", "május", "június", "július", "augusztus", "szeptember", "október", "november", "december"},
	},
	days: days{
		short: [7]string{"V", "H", "Sze", "K", "Cs", "P", "Szo"},
		long:  [7]string{"vasárnap", "hétfő", "szerda", "kedd", "csütörtök", "péntek", "szombat"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006. January 2., Monday", "2006. January 2.", "2006. Jan 2.", "2006. 01. 02."},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"2006. January 2., Monday 15:04:05 MST -0700", "2006. January 2. 15:04:05 MST", "2006. Jan 2. 15:04:05", "2006. 01. 02. 15:04"},
	},
}
var l_hy = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"հնվ", "փտվ", "մրտ", "ապր", "մյս", "հնս", "հլս", "օգս", "սեպ", "հոկ", "նոյ", "դեկ"},
		long:  [12]string{"հունվարի", "փետրվարի", "մարտի", "ապրիլի", "մայիսի", "հունիսի", "հուլիսի", "օգոստոսի", "սեպտեմբերի", "հոկտեմբերի", "նոյեմբերի", "դեկտեմբերի"},
	},
	days: days{
		short: [7]string{"կիր", "երկ", "չրք", "երք", "հնգ", "ուր", "շբթ"},
		long:  [7]string{"կիրակի", "երկուշաբթի", "չորեքշաբթի", "երեքշաբթի", "հինգշաբթի", "ուրբաթ", "շաբաթ"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006 թ. January 2, Monday", "02 January, 2006 թ.", "02 Jan, 2006 թ.", "02.01.06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"2006 թ. January 2, Monday, 15:04:05 MST -0700", "02 January, 2006 թ., 15:04:05 MST", "02 Jan, 2006 թ., 15:04:05", "02.01.06, 15:04"},
	},
}
var l_id = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agu", "Sep", "Okt", "Nov", "Des"},
		long:  [12]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"},
	},
	days: days{
		short: [7]string{"Min", "Sen", "Rab", "Sel", "Kam", "Jum", "Sab"},
		long:  [7]string{"Minggu", "Senin", "Rabu", "Selasa", "Kamis", "Jumat", "Sabtu"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 02 January 2006", "2 January 2006", "2 Jan 2006", "02/01/06"},
		time:     [4]string{"15.04.05 MST -0700", "15.04.05 MST", "15.04.05", "15.04"},
		datetime: [4]string{"Monday, 02 January 2006 15.04.05 MST -0700", "2 January 2006 15.04.05 MST", "2 Jan 2006 15.04.05", "02/01/06 15.04"},
	},
}
var l_is = &localize{
	am:              "f.h.",
	pm:              "e.h.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan.", "feb.", "mar.", "apr.", "maí", "jún.", "júl.", "ágú.", "sep.", "okt.", "nóv.", "des."},
		long:  [12]string{"janúar", "febrúar", "mars", "apríl", "maí", "júní", "júlí", "ágúst", "september", "október", "nóvember", "desember"},
	},
	days: days{
		short: [7]string{"sun.", "mán.", "mið.", "þri.", "fim.", "fös.", "lau."},
		long:  [7]string{"sunnudagur", "mánudagur", "miðvikudagur", "þriðjudagur", "fimmtudagur", "föstudagur", "laugardagur"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2. January 2006", "2. January 2006", "2. Jan 2006", "2.1.2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2. January 2006 kl. 15:04:05 MST -0700", "2. January 2006 kl. 15:04:05 MST", "2. Jan 2006, 15:04:05", "2.1.2006, 15:04"},
	},
}
var l_it = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"gen", "feb", "mar", "apr", "mag", "giu", "lug", "ago", "set", "ott", "nov", "dic"},
		long:  [12]string{"gennaio", "febbraio", "marzo", "aprile", "maggio", "giugno", "luglio", "agosto", "settembre", "ottobre", "novembre", "dicembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mer", "mar", "gio", "ven", "sab"},
		long:  [7]string{"domenica", "lunedì", "mercoledì", "martedì", "giovedì", "venerdì", "sabato"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 2 January 2006 15:04:05 MST -0700", "2 January 2006 15:04:05 MST", "2 Jan 2006, 15:04:05", "02/01/06, 15:04"},
	},
}
var l_it_CH = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"gen", "feb", "mar", "apr", "mag", "giu", "lug", "ago", "set", "ott", "nov", "dic"},
		long:  [12]string{"gennaio", "febbraio", "marzo", "aprile", "maggio", "giugno", "luglio", "agosto", "settembre", "ottobre", "novembre", "dicembre"},
	},
	days: days{
		short: [7]string{"dom", "lun", "mer", "mar", "gio", "ven", "sab"},
		long:  [7]string{"domenica", "lunedì", "mercoledì", "martedì", "giovedì", "venerdì", "sabato"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "02.01.06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006 15:04:05 MST -0700", "2 January 2006 15:04:05 MST", "2 Jan 2006, 15:04:05", "02.01.06, 15:04"},
	},
}
var l_ja = &localize{
	am:              "午前",
	pm:              "午後",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		long:  [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
	},
	days: days{
		short: [7]string{"日", "月", "水", "火", "木", "金", "土"},
		long:  [7]string{"日曜日", "月曜日", "水曜日", "火曜日", "木曜日", "金曜日", "土曜日"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006年1月2日Monday", "2006年1月2日", "2006/01/02", "2006/01/02"},
		time:     [4]string{"15時04分05秒 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"2006年1月2日Monday 15時04分05秒 MST -0700", "2006年1月2日 15:04:05 MST", "2006/01/02 15:04:05", "2006/01/02 15:04"},
	},
}
var l_jv = &localize{
	am:              "Isuk",
	pm:              "Wengi",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agt", "Sep", "Okt", "Nov", "Des"},
		long:  [12]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"},
	},
	days: days{
		short: [7]string{"Ahad", "Sen", "Rab", "Sel", "Kam", "Jum", "Sab"},
		long:  [7]string{"Ahad", "Senin", "Rabu", "Selasa", "Kamis", "Jumat", "Sabtu"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "02-01-2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006 15:04:05 MST -0700", "2 January 2006 15:04:05 MST", "2 Jan 2006, 15:04:05", "02-01-2006, 15:04"},
	},
}
var l_ka = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"იან", "თებ", "მარ", "აპრ", "მაი", "ივნ", "ივლ", "აგვ", "სექ", "ოქტ", "ნოე", "დეკ"},
		long:  [12]string{"იანვარი", "თებერვალი", "მარტი", "აპრილი", "მაისი", "ივნისი", "ივლისი", "აგვისტო", "სექტემბერი", "ოქტომბერი", "ნოემბერი", "დეკემბერი"},
	},
	days: days{
		short: [7]string{"კვი", "ორშ", "ოთხ", "სამ", "ხუთ", "პარ", "შაბ"},
		long:  [7]string{"კვირა", "ორშაბათი", "ოთხშაბათი", "სამშაბათი", "ხუთშაბათი", "პარასკევი", "შაბათი"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 02 January, 2006", "2 January, 2006", "2 Jan. 2006", "02.01.06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 02 January, 2006, 15:04:05 MST -0700", "2 January, 2006, 15:04:05 MST", "2 Jan. 2006, 15:04:05", "02.01.06, 15:04"},
	},
}
var l_kk = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"қаң.", "ақп.", "нау.", "сәу.", "мам.", "мау.", "шіл.", "там.", "қыр.", "қаз.", "қар.", "жел."},
		long:  [12]string{"қаңтар", "ақпан", "наурыз", "сәуір", "мамыр", "маусым", "шілде", "тамыз", "қыркүйек", "қазан", "қараша", "желтоқсан"},
	},
	days: days{
		short: [7]string{"жс", "дс", "ср", "сс", "бс", "жм", "сб"},
		long:  [7]string{"жексенбі", "дүйсенбі", "сәрсенбі", "сейсенбі", "бейсенбі", "жұма", "сенбі"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006 ж. 2 January, Monday", "2006 ж. 2 January", "2006 ж. 02 Jan", "02.01.06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"2006 ж. 2 January, Monday, 15:04:05 MST -0700", "2006 ж. 2 January, 15:04:05 MST", "2006 ж. 02 Jan, 15:04:05", "02.01.06, 15:04"},
	},
}
var l_km = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"មករា", "កុម្ភៈ", "មីនា", "មេសា", "ឧសភា", "មិថុនា", "កក្កដា", "សីហា", "កញ្ញា", "តុលា", "វិច្ឆិកា", "ធ្នូ"},
		long:  [12]string{"មករា", "កុម្ភៈ", "មីនា", "មេសា", "ឧសភា", "មិថុនា", "កក្កដា", "សីហា", "កញ្ញា", "តុលា", "វិច្ឆិកា", "ធ្នូ"},
	},
	days: days{
		short: [7]string{"អាទិត្យ", "ចន្ទ", "ពុធ", "អង្គារ", "ព្រហ", "សុក្រ", "សៅរ៍"},
		long:  [7]string{"អាទិត្យ", "ច័ន្ទ", "ពុធ", "អង្គារ", "ព្រហស្បតិ៍", "សុក្រ", "សៅរ៍"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2 January 2006", "2 January 2006", "2 Jan 2006", "2/1/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday 2 January 2006 នៅ\u200bម៉ោង 3:04:05 PM MST -0700", "2 January 2006 នៅ\u200bម៉ោង 3:04:05 PM MST", "2 Jan 2006, 3:04:05 PM", "2/1/06, 3:04 PM"},
	},
}
var l_kn = &localize{
	am:              "ಪೂರ್ವಾಹ್ನ",
	pm:              "ಅಪರಾಹ್ನ",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ಜನವರಿ", "ಫೆಬ್ರವರಿ", "ಮಾರ್ಚ್", "ಏಪ್ರಿ", "ಮೇ", "ಜೂನ್", "ಜುಲೈ", "ಆಗ", "ಸೆಪ್ಟೆಂ", "ಅಕ್ಟೋ", "ನವೆಂ", "ಡಿಸೆಂ"},
		long:  [12]string{"ಜನವರಿ", "ಫೆಬ್ರವರಿ", "ಮಾರ್ಚ್", "ಏಪ್ರಿಲ್", "ಮೇ", "ಜೂನ್", "ಜುಲೈ", "ಆಗಸ್ಟ್", "ಸೆಪ್ಟೆಂಬರ್", "ಅಕ್ಟೋಬರ್", "ನವೆಂಬರ್", "ಡಿಸೆಂಬರ್"},
	},
	days: days{
		short: [7]string{"ಭಾನು", "ಸೋಮ", "ಬುಧ", "ಮಂಗಳ", "ಗುರು", "ಶುಕ್ರ", "ಶನಿ"},
		long:  [7]string{"ಭಾನುವಾರ", "ಸೋಮವಾರ", "ಬುಧವಾರ", "ಮಂಗಳವಾರ", "ಗುರುವಾರ", "ಶುಕ್ರವಾರ", "ಶನಿವಾರ"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, January 2, 2006", "January 2, 2006", "Jan 2, 2006", "2/1/06"},
		time:     [4]string{"03:04:05 PM MST -0700", "03:04:05 PM MST", "03:04:05 PM", "03:04 PM"},
		datetime: [4]string{"Monday, January 2, 2006 03:04:05 PM MST -0700", "January 2, 2006 03:04:05 PM MST", "Jan 2, 2006 03:04:05 PM", "2/1/06 03:04 PM"},
	},
}
var l_ko = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"1월", "2월", "3월", "4월", "5월", "6월", "7월", "8월", "9월", "10월", "11월", "12월"},
		long:  [12]string{"1월", "2월", "3월", "4월", "5월", "6월", "7월", "8월", "9월", "10월", "11월", "12월"},
	},
	days: days{
		short: [7]string{"일", "월", "수", "화", "목", "금", "토"},
		long:  [7]string{"일요일", "월요일", "수요일", "화요일", "목요일", "금요일", "토요일"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006년 1월 2일 Monday", "2006년 1월 2일", "2006. 1. 2.", "06. 1. 2."},
		time:     [4]string{"PM 3시 4분 5초 MST -0700", "PM 3시 4분 5초 MST", "PM 3:04:05", "PM 3:04"},
		datetime: [4]string{"2006년 1월 2일 Monday PM 3시 4분 5초 MST -0700", "2006년 1월 2일 PM 3시 4분 5초 MST", "2006. 1. 2. PM 3:04:05", "06. 1. 2. PM 3:04"},
	},
}
var l_ky = &localize{
	am:              "тң",
	pm:              "тк",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"янв.", "фев.", "мар.", "апр.", "май", "июн.", "июл.", "авг.", "сен.", "окт.", "ноя.", "дек."},
		long:  [12]string{"январь", "февраль", "март", "апрель", "май", "июнь", "июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь"},
	},
	days: days{
		short: [7]string{"жек.", "дүй.", "шарш.", "шейш.", "бейш.", "жума", "ишм."},
		long:  [7]string{"жекшемби", "дүйшөмбү", "шаршемби", "шейшемби", "бейшемби", "жума", "ишемби"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006-ж., 2-January, Monday", "2006-ж., 2-January", "2006-ж., 2-Jan", "2/1/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"2006-ж., 2-January, Monday 15:04:05 MST -0700", "2006-ж., 2-January 15:04:05 MST", "2006-ж., 2-Jan 15:04:05", "2/1/06 15:04"},
	},
}
var l_lo = &localize{
	am:              "ກ່ອນທ່ຽງ",
	pm:              "ຫຼັງທ່ຽງ",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ມ.ກ.", "ກ.ພ.", "ມ.ນ.", "ມ.ສ.", "ພ.ພ.", "ມິ.ຖ.", "ກ.ລ.", "ສ.ຫ.", "ກ.ຍ.", "ຕ.ລ.", "ພ.ຈ.", "ທ.ວ."},
		long:  [12]string{"ມັງກອນ", "ກຸມພາ", "ມີນາ", "ເມສາ", "ພຶດສະພາ", "ມິຖຸນາ", "ກໍລະກົດ", "ສິງຫາ", "ກັນຍາ", "ຕຸລາ", "ພະຈິກ", "ທັນວາ"},
	},
	days: days{
		short: [7]string{"ອາທິດ", "ຈັນ", "ພຸດ", "ອັງຄານ", "ພະຫັດ", "ສຸກ", "ເສົາ"},
		long:  [7]string{"ວັນອາທິດ", "ວັນຈັນ", "ວັນພຸດ", "ວັນອັງຄານ", "ວັນພະຫັດ", "ວັນສຸກ", "ວັນເສົາ"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday ທີ 2 January G 2006", "2 January 2006", "2 Jan 2006", "2/1/2006"},
		time:     [4]string{"15 ໂມງ 4 ນາທີ 05 ວິນາທີ MST -0700", "15 ໂມງ 4 ນາທີ 05 ວິນາທີ MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday ທີ 2 January G 2006, 15 ໂມງ 4 ນາທີ 05 ວິນາທີ MST -0700", "2 January 2006, 15 ໂມງ 4 ນາທີ 05 ວິນາທີ MST", "2 Jan 2006, 15:04:05", "2/1/2006, 15:04"},
	},
}
var l_lt = &localize{
	am:              "priešpiet",
	pm:              "popiet",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"saus.", "vas.", "kov.", "bal.", "geg.", "birž.", "liep.", "rugp.", "rugs.", "spal.", "lapkr.", "gruod."},
		long:  [12]string{"sausio", "vasario", "kovo", "balandžio", "gegužės", "birželio", "liepos", "rugpjūčio", "rugsėjo", "spalio", "lapkričio", "gruodžio"},
	},
	days: days{
		short: [7]string{"sk", "pr", "tr", "an", "kt", "pn", "št"},
		long:  [7]string{"sekmadienis", "pirmadienis", "trečiadienis", "antradienis", "ketvirtadienis", "penktadienis", "šeštadienis"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006 m. January 2 d., Monday", "2006 m. January 2 d.", "2006-01-02", "2006-01-02"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"2006 m. January 2 d., Monday 15:04:05 MST -0700", "2006 m. January 2 d. 15:04:05 MST", "2006-01-02 15:04:05", "2006-01-02 15:04"},
	},
}
var l_lv = &localize{
	am:              "priekšp.",
	pm:              "pēcp.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"janv.", "febr.", "marts", "apr.", "maijs", "jūn.", "jūl.", "aug.", "sept.", "okt.", "nov.", "dec."},
		long:  [12]string{"janvāris", "februāris", "marts", "aprīlis", "maijs", "jūnijs", "jūlijs", "augusts", "septembris", "oktobris", "novembris", "decembris"},
	},
	days: days{
		short: [7]string{"svētd.", "pirmd.", "trešd.", "otrd.", "ceturtd.", "piektd.", "sestd."},
		long:  [7]string{"svētdiena", "pirmdiena", "trešdiena", "otrdiena", "ceturtdiena", "piektdiena", "sestdiena"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2006. gada 2. January", "2006. gada 2. January", "2006. gada 2. Jan", "02.01.06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2006. gada 2. January 15:04:05 MST -0700", "2006. gada 2. January 15:04:05 MST", "2006. gada 2. Jan 15:04:05", "02.01.06 15:04"},
	},
}
var l_mk = &localize{
	am:              "претпл.",
	pm:              "попл.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"јан.", "фев.", "мар.", "апр.", "мај", "јун.", "јул.", "авг.", "септ.", "окт.", "ноем.", "дек."},
		long:  [12]string{"јануари", "февруари", "март", "април", "мај", "јуни", "јули", "август", "септември", "октомври", "ноември", "декември"},
	},
	days: days{
		short: [7]string{"нед.", "пон.", "сре.", "вт.", "чет.", "пет.", "саб."},
		long:  [7]string{"недела", "понеделник", "среда", "вторник", "четврток", "петок", "сабота"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2.1.2006", "2.1.06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006, во 15:04:05 MST -0700", "2 January 2006, во 15:04:05 MST", "2.1.2006, во 15:04:05", "2.1.06, во 15:04"},
	},
}
var l_ml = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ജനു", "ഫെബ്രു", "മാർ", "ഏപ്രി", "മേയ്", "ജൂൺ", "ജൂലൈ", "ഓഗ", "സെപ്റ്റം", "ഒക്ടോ", "നവം", "ഡിസം"},
		long:  [12]string{"ജനുവരി", "ഫെബ്രുവരി", "മാർച്ച്", "ഏപ്രിൽ", "മേയ്", "ജൂൺ", "ജൂലൈ", "ഓഗസ്റ്റ്", "സെപ്റ്റംബർ", "ഒക്\u200cടോബർ", "നവംബർ", "ഡിസംബർ"},
	},
	days: days{
		short: [7]string{"ഞായർ", "തിങ്കൾ", "ബുധൻ", "ചൊവ്വ", "വ്യാഴം", "വെള്ളി", "ശനി"},
		long:  [7]string{"ഞായറാഴ്\u200cച", "തിങ്കളാഴ്\u200cച", "ബുധനാഴ്\u200cച", "ചൊവ്വാഴ്ച", "വ്യാഴാഴ്\u200cച", "വെള്ളിയാഴ്\u200cച", "ശനിയാഴ്\u200cച"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006, January 2, Monday", "2006, January 2", "2006, Jan 2", "2/1/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"2006, January 2, Monday 3:04:05 PM MST -0700", "2006, January 2 3:04:05 PM MST", "2006, Jan 2 3:04:05 PM", "2/1/06 3:04 PM"},
	},
}
var l_mn = &localize{
	am:              "ү.ө.",
	pm:              "ү.х.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"1-р сар", "2-р сар", "3-р сар", "4-р сар", "5-р сар", "6-р сар", "7-р сар", "8-р сар", "9-р сар", "10-р сар", "11-р сар", "12-р сар"},
		long:  [12]string{"нэгдүгээр сар", "хоёрдугаар сар", "гуравдугаар сар", "дөрөвдүгээр сар", "тавдугаар сар", "зургаадугаар сар", "долоодугаар сар", "наймдугаар сар", "есдүгээр сар", "аравдугаар сар", "арван нэгдүгээр сар", "арван хоёрдугаар сар"},
	},
	days: days{
		short: [7]string{"Ня", "Да", "Лх", "Мя", "Пү", "Ба", "Бя"},
		long:  [7]string{"ням", "даваа", "лхагва", "мягмар", "пүрэв", "баасан", "бямба"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006 оны ынJanuary 2, Monday гараг", "2006 оны ынJanuary 2", "2006 оны ынJan 2", "2006.01.02"},
		time:     [4]string{"15:04:05 (MST -0700)", "15:04:05 (MST)", "15:04:05", "15:04"},
		datetime: [4]string{"2006 оны ынJanuary 2, Monday гараг 15:04:05 (MST -0700)", "2006 оны ынJanuary 2 15:04:05 (MST)", "2006 оны ынJan 2 15:04:05", "2006.01.02 15:04"},
	},
}
var l_mr = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"जाने", "फेब्रु", "मार्च", "एप्रि", "मे", "जून", "जुलै", "ऑग", "सप्टें", "ऑक्टो", "नोव्हें", "डिसें"},
		long:  [12]string{"जानेवारी", "फेब्रुवारी", "मार्च", "एप्रिल", "मे", "जून", "जुलै", "ऑगस्ट", "सप्टेंबर", "ऑक्टोबर", "नोव्हेंबर", "डिसेंबर"},
	},
	days: days{
		short: [7]string{"रवि", "सोम", "बुध", "मंगळ", "गुरु", "शुक्र", "शनि"},
		long:  [7]string{"रविवार", "सोमवार", "बुधवार", "मंगळवार", "गुरुवार", "शुक्रवार", "शनिवार"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January, 2006", "2 January, 2006", "2 Jan, 2006", "2/1/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January, 2006 रोजी 3:04:05 PM MST -0700", "2 January, 2006 रोजी 3:04:05 PM MST", "2 Jan, 2006, 3:04:05 PM", "2/1/06, 3:04 PM"},
	},
}
var l_ms = &localize{
	am:              "PG",
	pm:              "PTG",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mac", "Apr", "Mei", "Jun", "Jul", "Ogo", "Sep", "Okt", "Nov", "Dis"},
		long:  [12]string{"Januari", "Februari", "Mac", "April", "Mei", "Jun", "Julai", "Ogos", "September", "Oktober", "November", "Disember"},
	},
	days: days{
		short: [7]string{"Ahd", "Isn", "Rab", "Sel", "Kha", "Jum", "Sab"},
		long:  [7]string{"Ahad", "Isnin", "Rabu", "Selasa", "Khamis", "Jumaat", "Sabtu"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "2/01/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January 2006 3:04:05 PM MST -0700", "2 January 2006 3:04:05 PM MST", "2 Jan 2006, 3:04:05 PM", "2/01/06, 3:04 PM"},
	},
}
var l_ms_BN = &localize{
	am:              "PG",
	pm:              "PTG",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mac", "Apr", "Mei", "Jun", "Jul", "Ogo", "Sep", "Okt", "Nov", "Dis"},
		long:  [12]string{"Januari", "Februari", "Mac", "April", "Mei", "Jun", "Julai", "Ogos", "September", "Oktober", "November", "Disember"},
	},
	days: days{
		short: [7]string{"Ahd", "Isn", "Rab", "Sel", "Kha", "Jum", "Sab"},
		long:  [7]string{"Ahad", "Isnin", "Rabu", "Selasa", "Khamis", "Jumaat", "Sabtu"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"02 January 2006", "2 January 2006", "2 Jan 2006", "2/01/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"02 January 2006 3:04:05 PM MST -0700", "2 January 2006 3:04:05 PM MST", "2 Jan 2006, 3:04:05 PM", "2/01/06, 3:04 PM"},
	},
}
var l_ms_ID = &localize{
	am:              "PG",
	pm:              "PTG",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mac", "Apr", "Mei", "Jun", "Jul", "Ogo", "Sep", "Okt", "Nov", "Dis"},
		long:  [12]string{"Januari", "Februari", "Mac", "April", "Mei", "Jun", "Julai", "Ogos", "September", "Oktober", "November", "Disember"},
	},
	days: days{
		short: [7]string{"Ahd", "Isn", "Rab", "Sel", "Kha", "Jum", "Sab"},
		long:  [7]string{"Ahad", "Isnin", "Rabu", "Selasa", "Khamis", "Jumaat", "Sabtu"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 02 January 2006", "2 January 2006", "2 Jan 2006", "02/01/06"},
		time:     [4]string{"15.04.05 MST -0700", "15.04.05 MST", "15.04.05", "15.04"},
		datetime: [4]string{"Monday, 02 January 2006 15.04.05 MST -0700", "2 January 2006 15.04.05 MST", "2 Jan 2006, 15.04.05", "02/01/06, 15.04"},
	},
}
var l_ms_SG = &localize{
	am:              "PG",
	pm:              "PTG",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mac", "Apr", "Mei", "Jun", "Jul", "Ogo", "Sep", "Okt", "Nov", "Dis"},
		long:  [12]string{"Januari", "Februari", "Mac", "April", "Mei", "Jun", "Julai", "Ogos", "September", "Oktober", "November", "Disember"},
	},
	days: days{
		short: [7]string{"Ahd", "Isn", "Rab", "Sel", "Kha", "Jum", "Sab"},
		long:  [7]string{"Ahad", "Isnin", "Rabu", "Selasa", "Khamis", "Jumaat", "Sabtu"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "2/01/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January 2006 3:04:05 PM MST -0700", "2 January 2006 3:04:05 PM MST", "2 Jan 2006, 3:04:05 PM", "2/01/06, 3:04 PM"},
	},
}
var l_my = &localize{
	am:              "နံနက်",
	pm:              "ညနေ",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ဇန်", "ဖေ", "မတ်", "ဧ", "မေ", "ဇွန်", "ဇူ", "ဩ", "စက်", "အောက်", "နို", "ဒီ"},
		long:  [12]string{"ဇန်နဝါရီ", "ဖေဖော်ဝါရီ", "မတ်", "ဧပြီ", "မေ", "ဇွန်", "ဇူလိုင်", "ဩဂုတ်", "စက်တင်ဘာ", "အောက်တိုဘာ", "နိုဝင်ဘာ", "ဒီဇင်ဘာ"},
	},
	days: days{
		short: [7]string{"တနင်္ဂနွေ", "တနင်္လာ", "ဗုဒ္ဓဟူး", "အင်္ဂါ", "ကြာသပတေး", "သောကြာ", "စနေ"},
		long:  [7]string{"တနင်္ဂနွေ", "တနင်္လာ", "ဗုဒ္ဓဟူး", "အင်္ဂါ", "ကြာသပတေး", "သောကြာ", "စနေ"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006၊ January 2၊ Monday", "2006၊ 2 January", "2006၊ Jan 2", "02-01-06"},
		time:     [4]string{"MST -0700 15:04:05", "MST 15:04:05", "B 15:04:05", "B 15:04"},
		datetime: [4]string{"2006၊ January 2၊ Monday MST -0700 15:04:05", "2006၊ 2 January MST 15:04:05", "2006၊ Jan 2 B 15:04:05", "02-01-06 B 15:04"},
	},
}
var l_nb = &localize{
	am:              "a.m.",
	pm:              "p.m.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan.", "feb.", "mar.", "apr.", "mai", "jun.", "jul.", "aug.", "sep.", "okt.", "nov.", "des."},
		long:  [12]string{"januar", "februar", "mars", "april", "mai", "juni", "juli", "august", "september", "oktober", "november", "desember"},
	},
	days: days{
		short: [7]string{"søn.", "man.", "ons.", "tir.", "tor.", "fre.", "lør."},
		long:  [7]string{"søndag", "mandag", "onsdag", "tirsdag", "torsdag", "fredag", "lørdag"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2. January 2006", "2. January 2006", "2. Jan 2006", "02.01.2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 2. January 2006 kl. 15:04:05 MST -0700", "2. January 2006 kl. 15:04:05 MST", "2. Jan 2006, 15:04:05", "02.01.2006, 15:04"},
	},
}
var l_ne = &localize{
	am:              "पूर्वाह्न",
	pm:              "अपराह्न",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"जनवरी", "फेब्रुअरी", "मार्च", "अप्रिल", "मे", "जुन", "जुलाई", "अगस्ट", "सेप्टेम्बर", "अक्टोबर", "नोभेम्बर", "डिसेम्बर"},
		long:  [12]string{"जनवरी", "फेब्रुअरी", "मार्च", "अप्रिल", "मे", "जुन", "जुलाई", "अगस्ट", "सेप्टेम्बर", "अक्टोबर", "नोभेम्बर", "डिसेम्बर"},
	},
	days: days{
		short: [7]string{"आइत", "सोम", "बुध", "मङ्गल", "बिहि", "शुक्र", "शनि"},
		long:  [7]string{"आइतबार", "सोमबार", "बुधबार", "मङ्गलबार", "बिहिबार", "शुक्रबार", "शनिबार"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006 January 2, Monday", "2006 January 2", "2006 Jan 2", "06/1/2"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"2006 January 2, Monday 15:04:05 MST -0700", "2006 January 2 15:04:05 MST", "2006 Jan 2, 15:04:05", "06/1/2, 15:04"},
	},
}
var l_ne_IN = &localize{
	am:              "पूर्वाह्न",
	pm:              "अपराह्न",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"जनवरी", "फेब्रुअरी", "मार्च", "अप्रिल", "मे", "जुन", "जुलाई", "अगस्ट", "सेप्टेम्बर", "अक्टोबर", "नोभेम्बर", "डिसेम्बर"},
		long:  [12]string{"जनवरी", "फेब्रुअरी", "मार्च", "अप्रिल", "मे", "जुन", "जुलाई", "अगस्ट", "सेप्टेम्बर", "अक्टोबर", "नोभेम्बर", "डिसेम्बर"},
	},
	days: days{
		short: [7]string{"आइत", "सोम", "बुध", "मङ्गल", "बिहि", "शुक्र", "शनि"},
		long:  [7]string{"आइतबार", "सोमबार", "बुधबार", "मङ्गलबार", "बिहिबार", "शुक्रबार", "शनिबार"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006 January 2, Monday", "2006 January 2", "2006 Jan 2", "06/1/2"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"2006 January 2, Monday 3:04:05 PM MST -0700", "2006 January 2 3:04:05 PM MST", "2006 Jan 2, 3:04:05 PM", "06/1/2, 3:04 PM"},
	},
}
var l_nl = &localize{
	am:              "a.m.",
	pm:              "p.m.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan.", "feb.", "mrt.", "apr.", "mei", "jun.", "jul.", "aug.", "sep.", "okt.", "nov.", "dec."},
		long:  [12]string{"januari", "februari", "maart", "april", "mei", "juni", "juli", "augustus", "september", "oktober", "november", "december"},
	},
	days: days{
		short: [7]string{"zo", "ma", "wo", "di", "do", "vr", "za"},
		long:  [7]string{"zondag", "maandag", "woensdag", "dinsdag", "donderdag", "vrijdag", "zaterdag"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2 January 2006", "2 January 2006", "2 Jan 2006", "02-01-2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 2 January 2006 om 15:04:05 MST -0700", "2 January 2006 om 15:04:05 MST", "2 Jan 2006 15:04:05", "02-01-2006 15:04"},
	},
}
var l_nl_BE = &localize{
	am:              "a.m.",
	pm:              "p.m.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan.", "feb.", "mrt.", "apr.", "mei", "jun.", "jul.", "aug.", "sep.", "okt.", "nov.", "dec."},
		long:  [12]string{"januari", "februari", "maart", "april", "mei", "juni", "juli", "augustus", "september", "oktober", "november", "december"},
	},
	days: days{
		short: [7]string{"zo", "ma", "wo", "di", "do", "vr", "za"},
		long:  [7]string{"zondag", "maandag", "woensdag", "dinsdag", "donderdag", "vrijdag", "zaterdag"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2 January 2006", "2 January 2006", "2 Jan 2006", "2/01/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 2 January 2006 om 15:04:05 MST -0700", "2 January 2006 om 15:04:05 MST", "2 Jan 2006 15:04:05", "2/01/2006 15:04"},
	},
}
var l_nn = &localize{
	am:              "f.m.",
	pm:              "e.m.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan.", "feb.", "mars", "apr.", "mai", "juni", "juli", "aug.", "sep.", "okt.", "nov.", "des."},
		long:  [12]string{"januar", "februar", "mars", "april", "mai", "juni", "juli", "august", "september", "oktober", "november", "desember"},
	},
	days: days{
		short: [7]string{"sø.", "må.", "on.", "ty.", "to.", "fr.", "la."},
		long:  [7]string{"søndag", "måndag", "onsdag", "tysdag", "torsdag", "fredag", "laurdag"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2. January 2006", "2. January 2006", "2. Jan 2006", "02.01.2006"},
		time:     [4]string{"kl. 15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 2. January 2006 kl. 15:04:05 MST -0700", "2. January 2006 kl. 15:04:05 MST", "2. Jan 2006, 15:04:05", "02.01.2006, 15:04"},
	},
}
var l_or = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ଜାନୁଆରୀ", "ଫେବୃଆରୀ", "ମାର୍ଚ୍ଚ", "ଅପ୍ରେଲ", "ମଇ", "ଜୁନ", "ଜୁଲାଇ", "ଅଗଷ୍ଟ", "ସେପ୍ଟେମ୍ବର", "ଅକ୍ଟୋବର", "ନଭେମ୍ବର", "ଡିସେମ୍ବର"},
		long:  [12]string{"ଜାନୁଆରୀ", "ଫେବୃଆରୀ", "ମାର୍ଚ୍ଚ", "ଅପ୍ରେଲ", "ମଇ", "ଜୁନ", "ଜୁଲାଇ", "ଅଗଷ୍ଟ", "ସେପ୍ଟେମ୍ବର", "ଅକ୍ଟୋବର", "ନଭେମ୍ବର", "ଡିସେମ୍ବର"},
	},
	days: days{
		short: [7]string{"ରବି", "ସୋମ", "ବୁଧ", "ମଙ୍ଗଳ", "ଗୁରୁ", "ଶୁକ୍ର", "ଶନି"},
		long:  [7]string{"ରବିବାର", "ସୋମବାର", "ବୁଧବାର", "ମଙ୍ଗଳବାର", "ଗୁରୁବାର", "ଶୁକ୍ରବାର", "ଶନିବାର"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, January 2, 2006", "January 2, 2006", "Jan 2, 2006", "1/2/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"3:04:05 PM MST -0700 ଠାରେ Monday, January 2, 2006", "3:04:05 PM MST ଠାରେ January 2, 2006", "Jan 2, 2006, 3:04:05 PM", "1/2/06, 3:04 PM"},
	},
}
var l_pa = &localize{
	am:              "ਪੂ.ਦੁ.",
	pm:              "ਬਾ.ਦੁ.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ਜਨ", "ਫ਼ਰ", "ਮਾਰਚ", "ਅਪ੍ਰੈ", "ਮਈ", "ਜੂਨ", "ਜੁਲਾ", "ਅਗ", "ਸਤੰ", "ਅਕਤੂ", "ਨਵੰ", "ਦਸੰ"},
		long:  [12]string{"ਜਨਵਰੀ", "ਫ਼ਰਵਰੀ", "ਮਾਰਚ", "ਅਪ੍ਰੈਲ", "ਮਈ", "ਜੂਨ", "ਜੁਲਾਈ", "ਅਗਸਤ", "ਸਤੰਬਰ", "ਅਕਤੂਬਰ", "ਨਵੰਬਰ", "ਦਸੰਬਰ"},
	},
	days: days{
		short: [7]string{"ਐਤ", "ਸੋਮ", "ਬੁੱਧ", "ਮੰਗਲ", "ਵੀਰ", "ਸ਼ੁੱਕਰ", "ਸ਼ਨਿੱਚਰ"},
		long:  [7]string{"ਐਤਵਾਰ", "ਸੋਮਵਾਰ", "ਬੁੱਧਵਾਰ", "ਮੰਗਲਵਾਰ", "ਵੀਰਵਾਰ", "ਸ਼ੁੱਕਰਵਾਰ", "ਸ਼ਨਿੱਚਰਵਾਰ"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "2/1/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January 2006 3:04:05 PM MST -0700", "2 January 2006 3:04:05 PM MST", "2 Jan 2006, 3:04:05 PM", "2/1/06, 3:04 PM"},
	},
}
var l_pl = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"sty", "lut", "mar", "kwi", "maj", "cze", "lip", "sie", "wrz", "paź", "lis", "gru"},
		long:  [12]string{"stycznia", "lutego", "marca", "kwietnia", "maja", "czerwca", "lipca", "sierpnia", "września", "października", "listopada", "grudnia"},
	},
	days: days{
		short: [7]string{"niedz.", "pon.", "śr.", "wt.", "czw.", "pt.", "sob."},
		long:  [7]string{"niedziela", "poniedziałek", "środa", "wtorek", "czwartek", "piątek", "sobota"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "02.01.2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006 15:04:05 MST -0700", "2 January 2006 15:04:05 MST", "2 Jan 2006, 15:04:05", "02.01.2006, 15:04"},
	},
}
var l_ps = &localize{
	am:              "غ.م.",
	pm:              "غ.و.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"جنوري", "فبروري", "مارچ", "اپریل", "مۍ", "جون", "جولای", "اګست", "سېپتمبر", "اکتوبر", "نومبر", "دسمبر"},
		long:  [12]string{"جنوري", "فبروري", "مارچ", "اپریل", "مۍ", "جون", "جولای", "اګست", "سېپتمبر", "اکتوبر", "نومبر", "دسمبر"},
	},
	days: days{
		short: [7]string{"يونۍ", "دونۍ", "څلرنۍ", "درېنۍ", "پينځنۍ", "جمعه", "اونۍ"},
		long:  [7]string{"يونۍ", "دونۍ", "څلرنۍ", "درېنۍ", "پينځنۍ", "جمعه", "اونۍ"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday د 2006 د January 2", "د 2006 د January 2", "2006 Jan 2", "2006/1/2"},
		time:     [4]string{"15:04:05 (MST -0700)", "15:04:05 (MST)", "15:04:05", "15:04"},
		datetime: [4]string{"Monday د 2006 د January 2 15:04:05 (MST -0700)", "د 2006 د January 2 15:04:05 (MST)", "2006 Jan 2 15:04:05", "2006/1/2 15:04"},
	},
}
var l_ps_PK = &localize{
	am:              "غ.م.",
	pm:              "غ.و.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"جنوري", "فبروري", "مارچ", "اپریل", "مۍ", "جون", "جولای", "اګست", "سېپتمبر", "اکتوبر", "نومبر", "دسمبر"},
		long:  [12]string{"جنوري", "فبروري", "مارچ", "اپریل", "مۍ", "جون", "جولای", "اګست", "سېپتمبر", "اکتوبر", "نومبر", "دسمبر"},
	},
	days: days{
		short: [7]string{"يونۍ", "دونۍ", "څلرنۍ", "درېنۍ", "پينځنۍ", "جمعه", "اونۍ"},
		long:  [7]string{"يونۍ", "دونۍ", "څلرنۍ", "درېنۍ", "پينځنۍ", "جمعه", "اونۍ"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday د 2006 د January 2", "د 2006 د January 2", "2006 Jan 2", "2006/1/2"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday د 2006 د January 2 3:04:05 PM MST -0700", "د 2006 د January 2 3:04:05 PM MST", "2006 Jan 2 3:04:05 PM", "2006/1/2 3:04 PM"},
	},
}
var l_pt = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan.", "fev.", "mar.", "abr.", "mai.", "jun.", "jul.", "ago.", "set.", "out.", "nov.", "dez."},
		long:  [12]string{"janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"},
	},
	days: days{
		short: [7]string{"dom.", "seg.", "qua.", "ter.", "qui.", "sex.", "sáb."},
		long:  [7]string{"domingo", "segunda-feira", "quarta-feira", "terça-feira", "quinta-feira", "sexta-feira", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "2 de Jan de 2006", "02/01/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 de January de 2006 15:04:05 MST -0700", "2 de January de 2006 15:04:05 MST", "2 de Jan de 2006 15:04:05", "02/01/2006 15:04"},
	},
}
var l_pt_AO = &localize{
	am:              "a.m.",
	pm:              "p.m.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan.", "fev.", "mar.", "abr.", "mai.", "jun.", "jul.", "ago.", "set.", "out.", "nov.", "dez."},
		long:  [12]string{"janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"},
	},
	days: days{
		short: [7]string{"domingo", "segunda", "quarta", "terça", "quinta", "sexta", "sábado"},
		long:  [7]string{"domingo", "segunda-feira", "quarta-feira", "terça-feira", "quinta-feira", "sexta-feira", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "02/01/2006", "02/01/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 de January de 2006 às 15:04:05 MST -0700", "2 de January de 2006 às 15:04:05 MST", "02/01/2006, 15:04:05", "02/01/06, 15:04"},
	},
}
var l_pt_CH = &localize{
	am:              "a.m.",
	pm:              "p.m.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan.", "fev.", "mar.", "abr.", "mai.", "jun.", "jul.", "ago.", "set.", "out.", "nov.", "dez."},
		long:  [12]string{"janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"},
	},
	days: days{
		short: [7]string{"domingo", "segunda", "quarta", "terça", "quinta", "sexta", "sábado"},
		long:  [7]string{"domingo", "segunda-feira", "quarta-feira", "terça-feira", "quinta-feira", "sexta-feira", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "02/01/2006", "02/01/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 de January de 2006 às 15:04:05 MST -0700", "2 de January de 2006 às 15:04:05 MST", "02/01/2006, 15:04:05", "02/01/06, 15:04"},
	},
}
var l_pt_MO = &localize{
	am:              "a.m.",
	pm:              "p.m.",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan.", "fev.", "mar.", "abr.", "mai.", "jun.", "jul.", "ago.", "set.", "out.", "nov.", "dez."},
		long:  [12]string{"janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"},
	},
	days: days{
		short: [7]string{"domingo", "segunda", "quarta", "terça", "quinta", "sexta", "sábado"},
		long:  [7]string{"domingo", "segunda-feira", "quarta-feira", "terça-feira", "quinta-feira", "sexta-feira", "sábado"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 de January de 2006", "2 de January de 2006", "02/01/2006", "02/01/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 de January de 2006 às 3:04:05 PM MST -0700", "2 de January de 2006 às 3:04:05 PM MST", "02/01/2006, 3:04:05 PM", "02/01/06, 3:04 PM"},
	},
}
var l_ro = &localize{
	am:              "a.m.",
	pm:              "p.m.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ian.", "feb.", "mar.", "apr.", "mai", "iun.", "iul.", "aug.", "sept.", "oct.", "nov.", "dec."},
		long:  [12]string{"ianuarie", "februarie", "martie", "aprilie", "mai", "iunie", "iulie", "august", "septembrie", "octombrie", "noiembrie", "decembrie"},
	},
	days: days{
		short: [7]string{"dum.", "lun.", "mie.", "mar.", "joi", "vin.", "sâm."},
		long:  [7]string{"duminică", "luni", "miercuri", "marți", "joi", "vineri", "sâmbătă"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "02.01.2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006, 15:04:05 MST -0700", "2 January 2006, 15:04:05 MST", "2 Jan 2006, 15:04:05", "02.01.2006, 15:04"},
	},
}
var l_ro_MD = &localize{
	am:              "a.m.",
	pm:              "p.m.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ian.", "feb.", "mar.", "apr.", "mai", "iun.", "iul.", "aug.", "sept.", "oct.", "nov.", "dec."},
		long:  [12]string{"ianuarie", "februarie", "martie", "aprilie", "mai", "iunie", "iulie", "august", "septembrie", "octombrie", "noiembrie", "decembrie"},
	},
	days: days{
		short: [7]string{"Dum", "Lun", "Mie", "Mar", "Joi", "Vin", "Sâm"},
		long:  [7]string{"duminică", "luni", "miercuri", "marți", "joi", "vineri", "sâmbătă"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "02.01.2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006, 15:04:05 MST -0700", "2 January 2006, 15:04:05 MST", "2 Jan 2006, 15:04:05", "02.01.2006, 15:04"},
	},
}
var l_root = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"M01", "M02", "M03", "M04", "M05", "M06", "M07", "M08", "M09", "M10", "M11", "M12"},
		long:  [12]string{"M01", "M02", "M03", "M04", "M05", "M06", "M07", "M08", "M09", "M10", "M11", "M12"},
	},
	days: days{
		short: [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
		long:  [7]string{"Sun", "Mon", "Wed", "Tue", "Thu", "Fri", "Sat"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006 January 2, Monday", "2006 January 2", "2006 Jan 2", "2006-01-02"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"2006 January 2, Monday 15:04:05 MST -0700", "2006 January 2 15:04:05 MST", "2006 Jan 2 15:04:05", "2006-01-02 15:04"},
	},
}
var l_ru = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"янв.", "февр.", "мар.", "апр.", "мая", "июн.", "июл.", "авг.", "сент.", "окт.", "нояб.", "дек."},
		long:  [12]string{"января", "февраля", "марта", "апреля", "мая", "июня", "июля", "августа", "сентября", "октября", "ноября", "декабря"},
	},
	days: days{
		short: [7]string{"вс", "пн", "ср", "вт", "чт", "пт", "сб"},
		long:  [7]string{"воскресенье", "понедельник", "среда", "вторник", "четверг", "пятница", "суббота"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006 г.", "2 January 2006 г.", "2 Jan 2006 г.", "02.01.2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006 г., 15:04:05 MST -0700", "2 January 2006 г., 15:04:05 MST", "2 Jan 2006 г., 15:04:05", "02.01.2006, 15:04"},
	},
}
var l_sd = &localize{
	am:              "صبح، منجهند",
	pm:              "شام، منجهند",
	firstDay:        6,
	twentyFourHours: false,
	months: months{
		short: [12]string{"جنوري", "فيبروري", "مارچ", "اپريل", "مئي", "جون", "جولاءِ", "آگسٽ", "سيپٽمبر", "آڪٽوبر", "نومبر", "ڊسمبر"},
		long:  [12]string{"جنوري", "فيبروري", "مارچ", "اپريل", "مئي", "جون", "جولاءِ", "آگسٽ", "سيپٽمبر", "آڪٽوبر", "نومبر", "ڊسمبر"},
	},
	days: days{
		short: [7]string{"آچر", "سومر", "اربع", "اڱارو", "خميس", "جمعو", "ڇنڇر"},
		long:  [7]string{"آچر", "سومر", "اربع", "اڱارو", "خميس", "جمعو", "ڇنڇر"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006 January 2, Monday", "2006 January 2", "2006 Jan 2", "2006-01-02"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"2006 January 2, Monday 3:04:05 PM MST -0700", "2006 January 2 3:04:05 PM MST", "2006 Jan 2 3:04:05 PM", "2006-01-02 3:04 PM"},
	},
}
var l_si = &localize{
	am:              "පෙ.ව.",
	pm:              "ප.ව.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ජන", "පෙබ", "මාර්තු", "අප්\u200dරේල්", "මැයි", "ජූනි", "ජූලි", "අගෝ", "සැප්", "ඔක්", "නොවැ", "දෙසැ"},
		long:  [12]string{"ජනවාරි", "පෙබරවාරි", "මාර්තු", "අප්\u200dරේල්", "මැයි", "ජූනි", "ජූලි", "අගෝස්තු", "සැප්තැම්බර්", "ඔක්තෝබර්", "නොවැම්බර්", "දෙසැම්බර්"},
	},
	days: days{
		short: [7]string{"ඉරිදා", "සඳුදා", "බදාදා", "අඟහ", "බ්\u200dරහස්", "සිකු", "සෙන"},
		long:  [7]string{"ඉරිදා", "සඳුදා", "බදාදා", "අඟහරුවාදා", "බ්\u200dරහස්පතින්දා", "සිකුරාදා", "සෙනසුරාදා"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006 January 2, Monday", "2006 January 2", "2006 Jan 2", "2006-01-02"},
		time:     [4]string{"15.04.05 MST -0700", "15.04.05 MST", "15.04.05", "15.04"},
		datetime: [4]string{"2006 January 2, Monday 15.04.05 MST -0700", "2006 January 2 15.04.05 MST", "2006 Jan 2 15.04.05", "2006-01-02 15.04"},
	},
}
var l_sk = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan", "feb", "mar", "apr", "máj", "jún", "júl", "aug", "sep", "okt", "nov", "dec"},
		long:  [12]string{"januára", "februára", "marca", "apríla", "mája", "júna", "júla", "augusta", "septembra", "októbra", "novembra", "decembra"},
	},
	days: days{
		short: [7]string{"ne", "po", "st", "ut", "št", "pi", "so"},
		long:  [7]string{"nedeľa", "pondelok", "streda", "utorok", "štvrtok", "piatok", "sobota"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2. January 2006", "2. January 2006", "2. 1. 2006", "2. 1. 2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 2. January 2006, 15:04:05 MST -0700", "2. January 2006, 15:04:05 MST", "2. 1. 2006, 15:04:05", "2. 1. 2006 15:04"},
	},
}
var l_sl = &localize{
	am:              "dop.",
	pm:              "pop.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan.", "feb.", "mar.", "apr.", "maj", "jun.", "jul.", "avg.", "sep.", "okt.", "nov.", "dec."},
		long:  [12]string{"januar", "februar", "marec", "april", "maj", "junij", "julij", "avgust", "september", "oktober", "november", "december"},
	},
	days: days{
		short: [7]string{"ned.", "pon.", "sre.", "tor.", "čet.", "pet.", "sob."},
		long:  [7]string{"nedelja", "ponedeljek", "sreda", "torek", "četrtek", "petek", "sobota"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 02. January 2006", "02. January 2006", "2. Jan 2006", "2. 01. 06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 02. January 2006 15:04:05 MST -0700", "02. January 2006 15:04:05 MST", "2. Jan 2006 15:04:05", "2. 01. 06 15:04"},
	},
}
var l_so = &localize{
	am:              "GH",
	pm:              "GD",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Abr", "May", "Jun", "Lul", "Ogs", "Seb", "Okt", "Nof", "Dis"},
		long:  [12]string{"Bisha Koobaad", "Bisha Labaad", "Bisha Saddexaad", "Bisha Afraad", "Bisha Shanaad", "Bisha Lixaad", "Bisha Todobaad", "Bisha Sideedaad", "Bisha Sagaalaad", "Bisha Tobnaad", "Bisha Kow iyo Tobnaad", "Bisha Laba iyo Tobnaad"},
	},
	days: days{
		short: [7]string{"Axd", "Isn", "Arbc", "Tldo", "Khms", "Jmc", "Sbti"},
		long:  [7]string{"Axad", "Isniin", "Arbaco", "Talaado", "Khamiis", "Jimco", "Sabti"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, January 02, 2006", "02 January 2006", "02-Jan-2006", "02/01/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, January 02, 2006 3:04:05 PM MST -0700", "02 January 2006 3:04:05 PM MST", "02-Jan-2006 3:04:05 PM", "02/01/06 3:04 PM"},
	},
}
var l_so_DJ = &localize{
	am:              "GH",
	pm:              "GD",
	firstDay:        6,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Abr", "May", "Jun", "Lul", "Ogs", "Seb", "Okt", "Nof", "Dis"},
		long:  [12]string{"Bisha Koobaad", "Bisha Labaad", "Bisha Saddexaad", "Bisha Afraad", "Bisha Shanaad", "Bisha Lixaad", "Bisha Todobaad", "Bisha Sideedaad", "Bisha Sagaalaad", "Bisha Tobnaad", "Bisha Kow iyo Tobnaad", "Bisha Laba iyo Tobnaad"},
	},
	days: days{
		short: [7]string{"Axd", "Isn", "Arbc", "Tldo", "Khms", "Jmc", "Sbti"},
		long:  [7]string{"Axad", "Isniin", "Arbaco", "Talaado", "Khamiis", "Jimco", "Sabti"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, January 02, 2006", "02 January 2006", "02-Jan-2006", "02/01/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, January 02, 2006 3:04:05 PM MST -0700", "02 January 2006 3:04:05 PM MST", "02-Jan-2006 3:04:05 PM", "02/01/06 3:04 PM"},
	},
}
var l_so_ET = &localize{
	am:              "GH",
	pm:              "GD",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Abr", "May", "Jun", "Lul", "Ogs", "Seb", "Okt", "Nof", "Dis"},
		long:  [12]string{"Bisha Koobaad", "Bisha Labaad", "Bisha Saddexaad", "Bisha Afraad", "Bisha Shanaad", "Bisha Lixaad", "Bisha Todobaad", "Bisha Sideedaad", "Bisha Sagaalaad", "Bisha Tobnaad", "Bisha Kow iyo Tobnaad", "Bisha Laba iyo Tobnaad"},
	},
	days: days{
		short: [7]string{"Axd", "Isn", "Arbc", "Tldo", "Khms", "Jmc", "Sbti"},
		long:  [7]string{"Axad", "Isniin", "Arbaco", "Talaado", "Khamiis", "Jimco", "Sabti"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, January 02, 2006", "02 January 2006", "02-Jan-2006", "02/01/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, January 02, 2006 3:04:05 PM MST -0700", "02 January 2006 3:04:05 PM MST", "02-Jan-2006 3:04:05 PM", "02/01/06 3:04 PM"},
	},
}
var l_so_KE = &localize{
	am:              "GH",
	pm:              "GD",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mar", "Abr", "May", "Jun", "Lul", "Ogs", "Seb", "Okt", "Nof", "Dis"},
		long:  [12]string{"Bisha Koobaad", "Bisha Labaad", "Bisha Saddexaad", "Bisha Afraad", "Bisha Shanaad", "Bisha Lixaad", "Bisha Todobaad", "Bisha Sideedaad", "Bisha Sagaalaad", "Bisha Tobnaad", "Bisha Kow iyo Tobnaad", "Bisha Laba iyo Tobnaad"},
	},
	days: days{
		short: [7]string{"Axd", "Isn", "Arbc", "Tldo", "Khms", "Jmc", "Sbti"},
		long:  [7]string{"Axad", "Isniin", "Arbaco", "Talaado", "Khamiis", "Jimco", "Sabti"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, January 02, 2006", "02 January 2006", "02-Jan-2006", "02/01/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, January 02, 2006 15:04:05 MST -0700", "02 January 2006 15:04:05 MST", "02-Jan-2006 15:04:05", "02/01/06 15:04"},
	},
}
var l_sq = &localize{
	am:              "p.d.",
	pm:              "m.d.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan", "shk", "mar", "pri", "maj", "qer", "korr", "gush", "sht", "tet", "nën", "dhj"},
		long:  [12]string{"janar", "shkurt", "mars", "prill", "maj", "qershor", "korrik", "gusht", "shtator", "tetor", "nëntor", "dhjetor"},
	},
	days: days{
		short: [7]string{"Die", "Hën", "Mër", "Mar", "Enj", "Pre", "Sht"},
		long:  [7]string{"e diel", "e hënë", "e mërkurë", "e martë", "e enjte", "e premte", "e shtunë"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "2.1.06"},
		time:     [4]string{"3:04:05 PM, MST -0700", "3:04:05 PM, MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday, 2 January 2006 në 3:04:05 PM, MST -0700", "2 January 2006 në 3:04:05 PM, MST", "2 Jan 2006, 3:04:05 PM", "2.1.06, 3:04 PM"},
	},
}
var l_sq_MK = &localize{
	am:              "p.d.",
	pm:              "m.d.",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan", "shk", "mar", "pri", "maj", "qer", "korr", "gush", "sht", "tet", "nën", "dhj"},
		long:  [12]string{"janar", "shkurt", "mars", "prill", "maj", "qershor", "korrik", "gusht", "shtator", "tetor", "nëntor", "dhjetor"},
	},
	days: days{
		short: [7]string{"Die", "Hën", "Mër", "Mar", "Enj", "Pre", "Sht"},
		long:  [7]string{"e diel", "e hënë", "e mërkurë", "e martë", "e enjte", "e premte", "e shtunë"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "2.1.06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006 në 15:04:05 MST -0700", "2 January 2006 në 15:04:05 MST", "2 Jan 2006, 15:04:05", "2.1.06, 15:04"},
	},
}
var l_sr = &localize{
	am:              "пре подне",
	pm:              "по подне",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"јан", "феб", "мар", "апр", "мај", "јун", "јул", "авг", "сеп", "окт", "нов", "дец"},
		long:  [12]string{"јануар", "фебруар", "март", "април", "мај", "јун", "јул", "август", "септембар", "октобар", "новембар", "децембар"},
	},
	days: days{
		short: [7]string{"нед", "пон", "сре", "уто", "чет", "пет", "суб"},
		long:  [7]string{"недеља", "понедељак", "среда", "уторак", "четвртак", "петак", "субота"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 02. January 2006.", "02. January 2006.", "02.01.2006.", "2.1.06."},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 02. January 2006. 15:04:05 MST -0700", "02. January 2006. 15:04:05 MST", "02.01.2006. 15:04:05", "2.1.06. 15:04"},
	},
}
var l_sr_Cyrl_BA = &localize{
	am:              "прије подне",
	pm:              "по подне",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"јан", "феб", "мар", "апр", "мај", "јун", "јул", "авг", "сеп", "окт", "нов", "дец"},
		long:  [12]string{"јануар", "фебруар", "март", "април", "мај", "јун", "јул", "август", "септембар", "октобар", "новембар", "децембар"},
	},
	days: days{
		short: [7]string{"нед", "пон", "сре", "уто", "чет", "пет", "суб"},
		long:  [7]string{"недјеља", "понедјељак", "сриједа", "уторак", "четвртак", "петак", "субота"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 02. January 2006.", "02. January 2006.", "02.01.2006.", "2.1.06."},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 02. January 2006. 15:04:05 MST -0700", "02. January 2006. 15:04:05 MST", "02.01.2006. 15:04:05", "2.1.06. 15:04"},
	},
}
var l_sr_Cyrl_ME = &localize{
	am:              "прије подне",
	pm:              "по подне",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"јан", "феб", "март", "апр", "мај", "јун", "јул", "авг", "септ", "окт", "нов", "дец"},
		long:  [12]string{"јануар", "фебруар", "март", "април", "мај", "јун", "јул", "август", "септембар", "октобар", "новембар", "децембар"},
	},
	days: days{
		short: [7]string{"нед", "пон", "сре", "уто", "чет", "пет", "суб"},
		long:  [7]string{"недјеља", "понедељак", "сриједа", "уторак", "четвртак", "петак", "субота"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 02. January 2006.", "02. January 2006.", "02.01.2006.", "2.1.06."},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 02. January 2006. 15:04:05 MST -0700", "02. January 2006. 15:04:05 MST", "02.01.2006. 15:04:05", "2.1.06. 15:04"},
	},
}
var l_sr_Cyrl_XK = &localize{
	am:              "пре подне",
	pm:              "по подне",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"јан", "феб", "март", "апр", "мај", "јун", "јул", "авг", "септ", "окт", "нов", "дец"},
		long:  [12]string{"јануар", "фебруар", "март", "април", "мај", "јун", "јул", "август", "септембар", "октобар", "новембар", "децембар"},
	},
	days: days{
		short: [7]string{"нед", "пон", "сре", "уто", "чет", "пет", "суб"},
		long:  [7]string{"недеља", "понедељак", "среда", "уторак", "четвртак", "петак", "субота"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 02. January 2006.", "02. January 2006.", "02.01.2006.", "2.1.06."},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 02. January 2006. 15:04:05 MST -0700", "02. January 2006. 15:04:05 MST", "02.01.2006. 15:04:05", "2.1.06. 15:04"},
	},
}
var l_sr_Latn = &localize{
	am:              "pre podne",
	pm:              "po podne",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan", "feb", "mar", "apr", "maj", "jun", "jul", "avg", "sep", "okt", "nov", "dec"},
		long:  [12]string{"januar", "februar", "mart", "april", "maj", "jun", "jul", "avgust", "septembar", "oktobar", "novembar", "decembar"},
	},
	days: days{
		short: [7]string{"ned", "pon", "sre", "uto", "čet", "pet", "sub"},
		long:  [7]string{"nedelja", "ponedeljak", "sreda", "utorak", "četvrtak", "petak", "subota"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 02. January 2006.", "02. January 2006.", "02.01.2006.", "2.1.06."},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 02. January 2006. 15:04:05 MST -0700", "02. January 2006. 15:04:05 MST", "02.01.2006. 15:04:05", "2.1.06. 15:04"},
	},
}
var l_sr_Latn_BA = &localize{
	am:              "prije podne",
	pm:              "po podne",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan", "feb", "mar", "apr", "maj", "jun", "jul", "avg", "sep", "okt", "nov", "dec"},
		long:  [12]string{"januar", "februar", "mart", "april", "maj", "jun", "jul", "avgust", "septembar", "oktobar", "novembar", "decembar"},
	},
	days: days{
		short: [7]string{"ned", "pon", "sre", "uto", "čet", "pet", "sub"},
		long:  [7]string{"nedjelja", "ponedjeljak", "srijeda", "utorak", "četvrtak", "petak", "subota"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 02. January 2006.", "02. January 2006.", "02.01.2006.", "2.1.06."},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 02. January 2006. 15:04:05 MST -0700", "02. January 2006. 15:04:05 MST", "02.01.2006. 15:04:05", "2.1.06. 15:04"},
	},
}
var l_sr_Latn_ME = &localize{
	am:              "prije podne",
	pm:              "po podne",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan", "feb", "mart", "apr", "maj", "jun", "jul", "avg", "sept", "okt", "nov", "dec"},
		long:  [12]string{"januar", "februar", "mart", "april", "maj", "jun", "jul", "avgust", "septembar", "oktobar", "novembar", "decembar"},
	},
	days: days{
		short: [7]string{"ned", "pon", "sre", "uto", "čet", "pet", "sub"},
		long:  [7]string{"nedjelja", "ponedeljak", "srijeda", "utorak", "četvrtak", "petak", "subota"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 02. January 2006.", "02. January 2006.", "02.01.2006.", "2.1.06."},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 02. January 2006. 15:04:05 MST -0700", "02. January 2006. 15:04:05 MST", "02.01.2006. 15:04:05", "2.1.06. 15:04"},
	},
}
var l_sr_Latn_XK = &localize{
	am:              "pre podne",
	pm:              "po podne",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan", "feb", "mart", "apr", "maj", "jun", "jul", "avg", "sept", "okt", "nov", "dec"},
		long:  [12]string{"januar", "februar", "mart", "april", "maj", "jun", "jul", "avgust", "septembar", "oktobar", "novembar", "decembar"},
	},
	days: days{
		short: [7]string{"ned", "pon", "sre", "uto", "čet", "pet", "sub"},
		long:  [7]string{"nedelja", "ponedeljak", "sreda", "utorak", "četvrtak", "petak", "subota"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 02. January 2006.", "02. January 2006.", "02.01.2006.", "2.1.06."},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 02. January 2006. 15:04:05 MST -0700", "02. January 2006. 15:04:05 MST", "02.01.2006. 15:04:05", "2.1.06. 15:04"},
	},
}
var l_sv = &localize{
	am:              "fm",
	pm:              "em",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan.", "feb.", "mars", "apr.", "maj", "juni", "juli", "aug.", "sep.", "okt.", "nov.", "dec."},
		long:  [12]string{"januari", "februari", "mars", "april", "maj", "juni", "juli", "augusti", "september", "oktober", "november", "december"},
	},
	days: days{
		short: [7]string{"sön", "mån", "ons", "tis", "tors", "fre", "lör"},
		long:  [7]string{"söndag", "måndag", "onsdag", "tisdag", "torsdag", "fredag", "lördag"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2 January 2006", "2 January 2006", "2 Jan 2006", "2006-01-02"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 2 January 2006 15:04:05 MST -0700", "2 January 2006 15:04:05 MST", "2 Jan 2006 15:04:05", "2006-01-02 15:04"},
	},
}
var l_sv_AX = &localize{
	am:              "fm",
	pm:              "em",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"jan.", "feb.", "mars", "apr.", "maj", "juni", "juli", "aug.", "sep.", "okt.", "nov.", "dec."},
		long:  [12]string{"januari", "februari", "mars", "april", "maj", "juni", "juli", "augusti", "september", "oktober", "november", "december"},
	},
	days: days{
		short: [7]string{"sön", "mån", "ons", "tis", "tors", "fre", "lör"},
		long:  [7]string{"söndag", "måndag", "onsdag", "tisdag", "torsdag", "fredag", "lördag"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday 2 January 2006", "2 January 2006", "2 Jan 2006", "2006-01-02"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday 2 January 2006 15:04:05 MST -0700", "2 January 2006 15:04:05 MST", "2 Jan 2006 15:04:05", "2006-01-02 15:04"},
	},
}
var l_sw = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mac", "Apr", "Mei", "Jun", "Jul", "Ago", "Sep", "Okt", "Nov", "Des"},
		long:  [12]string{"Januari", "Februari", "Machi", "Aprili", "Mei", "Juni", "Julai", "Agosti", "Septemba", "Oktoba", "Novemba", "Desemba"},
	},
	days: days{
		short: [7]string{"Jumapili", "Jumatatu", "Jumatano", "Jumanne", "Alhamisi", "Ijumaa", "Jumamosi"},
		long:  [7]string{"Jumapili", "Jumatatu", "Jumatano", "Jumanne", "Alhamisi", "Ijumaa", "Jumamosi"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006 15:04:05 MST -0700", "2 January 2006 15:04:05 MST", "2 Jan 2006 15:04:05", "02/01/2006 15:04"},
	},
}
var l_sw_KE = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mac", "Apr", "Mei", "Jun", "Jul", "Ago", "Sep", "Okt", "Nov", "Des"},
		long:  [12]string{"Januari", "Februari", "Machi", "Aprili", "Mei", "Juni", "Julai", "Agosti", "Septemba", "Oktoba", "Novemba", "Desemba"},
	},
	days: days{
		short: [7]string{"Jumapili", "Jumatatu", "Jumatano", "Jumanne", "Alhamisi", "Ijumaa", "Jumamosi"},
		long:  [7]string{"Jumapili", "Jumatatu", "Jumatano", "Jumanne", "Alhamisi", "Ijumaa", "Jumamosi"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006", "2 January 2006", "2 Jan 2006", "02/01/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006 15:04:05 MST -0700", "2 January 2006 15:04:05 MST", "2 Jan 2006 15:04:05", "02/01/2006 15:04"},
	},
}
var l_ta = &localize{
	am:              "முற்பகல்",
	pm:              "பிற்பகல்",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ஜன.", "பிப்.", "மார்.", "ஏப்.", "மே", "ஜூன்", "ஜூலை", "ஆக.", "செப்.", "அக்.", "நவ.", "டிச."},
		long:  [12]string{"ஜனவரி", "பிப்ரவரி", "மார்ச்", "ஏப்ரல்", "மே", "ஜூன்", "ஜூலை", "ஆகஸ்ட்", "செப்டம்பர்", "அக்டோபர்", "நவம்பர்", "டிசம்பர்"},
	},
	days: days{
		short: [7]string{"ஞாயி.", "திங்.", "புத.", "செவ்.", "வியா.", "வெள்.", "சனி"},
		long:  [7]string{"ஞாயிறு", "திங்கள்", "புதன்", "செவ்வாய்", "வியாழன்", "வெள்ளி", "சனி"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January, 2006", "2 January, 2006", "2 Jan, 2006", "2/1/06"},
		time:     [4]string{"PM 3:04:05 MST -0700", "PM 3:04:05 MST", "PM 3:04:05", "PM 3:04"},
		datetime: [4]string{"Monday, 2 January, 2006 ’அன்று’ PM 3:04:05 MST -0700", "2 January, 2006 ’அன்று’ PM 3:04:05 MST", "2 Jan, 2006, PM 3:04:05", "2/1/06, PM 3:04"},
	},
}
var l_ta_LK = &localize{
	am:              "முற்பகல்",
	pm:              "பிற்பகல்",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ஜன.", "பிப்.", "மார்.", "ஏப்.", "மே", "ஜூன்", "ஜூலை", "ஆக.", "செப்.", "அக்.", "நவ.", "டிச."},
		long:  [12]string{"ஜனவரி", "பிப்ரவரி", "மார்ச்", "ஏப்ரல்", "மே", "ஜூன்", "ஜூலை", "ஆகஸ்ட்", "செப்டம்பர்", "அக்டோபர்", "நவம்பர்", "டிசம்பர்"},
	},
	days: days{
		short: [7]string{"ஞாயி.", "திங்.", "புத.", "செவ்.", "வியா.", "வெள்.", "சனி"},
		long:  [7]string{"ஞாயிறு", "திங்கள்", "புதன்", "செவ்வாய்", "வியாழன்", "வெள்ளி", "சனி"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January, 2006", "2 January, 2006", "2 Jan, 2006", "2/1/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January, 2006 ’அன்று’ 15:04:05 MST -0700", "2 January, 2006 ’அன்று’ 15:04:05 MST", "2 Jan, 2006, 15:04:05", "2/1/06, 15:04"},
	},
}
var l_ta_SG = &localize{
	am:              "முற்பகல்",
	pm:              "பிற்பகல்",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ஜன.", "பிப்.", "மார்.", "ஏப்.", "மே", "ஜூன்", "ஜூலை", "ஆக.", "செப்.", "அக்.", "நவ.", "டிச."},
		long:  [12]string{"ஜனவரி", "பிப்ரவரி", "மார்ச்", "ஏப்ரல்", "மே", "ஜூன்", "ஜூலை", "ஆகஸ்ட்", "செப்டம்பர்", "அக்டோபர்", "நவம்பர்", "டிசம்பர்"},
	},
	days: days{
		short: [7]string{"ஞாயி.", "திங்.", "புத.", "செவ்.", "வியா.", "வெள்.", "சனி"},
		long:  [7]string{"ஞாயிறு", "திங்கள்", "புதன்", "செவ்வாய்", "வியாழன்", "வெள்ளி", "சனி"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January, 2006", "2 January, 2006", "2 Jan, 2006", "2/1/06"},
		time:     [4]string{"PM 3:04:05 MST -0700", "PM 3:04:05 MST", "PM 3:04:05", "PM 3:04"},
		datetime: [4]string{"Monday, 2 January, 2006 ’அன்று’ PM 3:04:05 MST -0700", "2 January, 2006 ’அன்று’ PM 3:04:05 MST", "2 Jan, 2006, PM 3:04:05", "2/1/06, PM 3:04"},
	},
}
var l_te = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"జన", "ఫిబ్ర", "మార్చి", "ఏప్రి", "మే", "జూన్", "జులై", "ఆగ", "సెప్టెం", "అక్టో", "నవం", "డిసెం"},
		long:  [12]string{"జనవరి", "ఫిబ్రవరి", "మార్చి", "ఏప్రిల్", "మే", "జూన్", "జులై", "ఆగస్టు", "సెప్టెంబర్", "అక్టోబర్", "నవంబర్", "డిసెంబర్"},
	},
	days: days{
		short: [7]string{"ఆది", "సోమ", "బుధ", "మంగళ", "గురు", "శుక్ర", "శని"},
		long:  [7]string{"ఆదివారం", "సోమవారం", "బుధవారం", "మంగళవారం", "గురువారం", "శుక్రవారం", "శనివారం"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2, January 2006, Monday", "2 January, 2006", "2 Jan, 2006", "02-01-06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"2, January 2006, Monday 3:04:05 PM MST -0700కి", "2 January, 2006 3:04:05 PM MSTకి", "2 Jan, 2006 3:04:05 PM", "02-01-06 3:04 PM"},
	},
}
var l_th = &localize{
	am:              "ก่อนเที่ยง",
	pm:              "หลังเที่ยง",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.", "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค."},
		long:  [12]string{"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน", "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"},
	},
	days: days{
		short: [7]string{"อา.", "จ.", "พ.", "อ.", "พฤ.", "ศ.", "ส."},
		long:  [7]string{"วันอาทิตย์", "วันจันทร์", "วันพุธ", "วันอังคาร", "วันพฤหัสบดี", "วันศุกร์", "วันเสาร์"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Mondayที่ 2 January G 2006", "2 January G 2006", "2 Jan 2006", "2/1/06"},
		time:     [4]string{"15 นาฬิกา 04 นาที 05 วินาที MST -0700", "15 นาฬิกา 04 นาที 05 วินาที MST", "15:04:05", "15:04"},
		datetime: [4]string{"Mondayที่ 2 January G 2006 15 นาฬิกา 04 นาที 05 วินาที MST -0700", "2 January G 2006 15 นาฬิกา 04 นาที 05 วินาที MST", "2 Jan 2006 15:04:05", "2/1/06 15:04"},
	},
}
var l_tk = &localize{
	am:              "go.öň",
	pm:              "go.soň",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"ýan", "few", "mart", "apr", "maý", "iýun", "iýul", "awg", "sen", "okt", "noý", "dek"},
		long:  [12]string{"ýanwar", "fewral", "mart", "aprel", "maý", "iýun", "iýul", "awgust", "sentýabr", "oktýabr", "noýabr", "dekabr"},
	},
	days: days{
		short: [7]string{"ýek", "duş", "çar", "siş", "pen", "ann", "şen"},
		long:  [7]string{"ýekşenbe", "duşenbe", "çarşenbe", "sişenbe", "penşenbe", "anna", "şenbe"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2 January 2006 Monday", "2 January 2006", "2 Jan 2006", "02.01.2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"2 January 2006 Monday 15:04:05 MST -0700", "2 January 2006 15:04:05 MST", "2 Jan 2006 15:04:05", "02.01.2006 15:04"},
	},
}
var l_tr = &localize{
	am:              "ÖÖ",
	pm:              "ÖS",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Oca", "Şub", "Mar", "Nis", "May", "Haz", "Tem", "Ağu", "Eyl", "Eki", "Kas", "Ara"},
		long:  [12]string{"Ocak", "Şubat", "Mart", "Nisan", "Mayıs", "Haziran", "Temmuz", "Ağustos", "Eylül", "Ekim", "Kasım", "Aralık"},
	},
	days: days{
		short: [7]string{"Paz", "Pzt", "Çar", "Sal", "Per", "Cum", "Cmt"},
		long:  [7]string{"Pazar", "Pazartesi", "Çarşamba", "Salı", "Perşembe", "Cuma", "Cumartesi"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2 January 2006 Monday", "2 January 2006", "2 Jan 2006", "2.01.2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"2 January 2006 Monday 15:04:05 MST -0700", "2 January 2006 15:04:05 MST", "2 Jan 2006 15:04:05", "2.01.2006 15:04"},
	},
}
var l_tr_CY = &localize{
	am:              "ÖÖ",
	pm:              "ÖS",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Oca", "Şub", "Mar", "Nis", "May", "Haz", "Tem", "Ağu", "Eyl", "Eki", "Kas", "Ara"},
		long:  [12]string{"Ocak", "Şubat", "Mart", "Nisan", "Mayıs", "Haziran", "Temmuz", "Ağustos", "Eylül", "Ekim", "Kasım", "Aralık"},
	},
	days: days{
		short: [7]string{"Paz", "Pzt", "Çar", "Sal", "Per", "Cum", "Cmt"},
		long:  [7]string{"Pazar", "Pazartesi", "Çarşamba", "Salı", "Perşembe", "Cuma", "Cumartesi"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2 January 2006 Monday", "2 January 2006", "2 Jan 2006", "2.01.2006"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"2 January 2006 Monday 3:04:05 PM MST -0700", "2 January 2006 3:04:05 PM MST", "2 Jan 2006 3:04:05 PM", "2.01.2006 3:04 PM"},
	},
}
var l_uk = &localize{
	am:              "дп",
	pm:              "пп",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"січ.", "лют.", "бер.", "квіт.", "трав.", "черв.", "лип.", "серп.", "вер.", "жовт.", "лист.", "груд."},
		long:  [12]string{"січня", "лютого", "березня", "квітня", "травня", "червня", "липня", "серпня", "вересня", "жовтня", "листопада", "грудня"},
	},
	days: days{
		short: [7]string{"нд", "пн", "ср", "вт", "чт", "пт", "сб"},
		long:  [7]string{"неділя", "понеділок", "середа", "вівторок", "четвер", "пʼятниця", "субота"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January 2006 р.", "2 January 2006 р.", "2 Jan 2006 р.", "02.01.06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2 January 2006 р. о 15:04:05 MST -0700", "2 January 2006 р. о 15:04:05 MST", "2 Jan 2006 р., 15:04:05", "02.01.06, 15:04"},
	},
}
var l_ur = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"جنوری", "فروری", "مارچ", "اپریل", "مئی", "جون", "جولائی", "اگست", "ستمبر", "اکتوبر", "نومبر", "دسمبر"},
		long:  [12]string{"جنوری", "فروری", "مارچ", "اپریل", "مئی", "جون", "جولائی", "اگست", "ستمبر", "اکتوبر", "نومبر", "دسمبر"},
	},
	days: days{
		short: [7]string{"اتوار", "پیر", "بدھ", "منگل", "جمعرات", "جمعہ", "ہفتہ"},
		long:  [7]string{"اتوار", "پیر", "بدھ", "منگل", "جمعرات", "جمعہ", "ہفتہ"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday، 2 January، 2006", "2 January، 2006", "2 Jan، 2006", "2/1/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday، 2 January، 2006 3:04:05 PM MST -0700", "2 January، 2006 3:04:05 PM MST", "2 Jan، 2006 3:04:05 PM", "2/1/06 3:04 PM"},
	},
}
var l_ur_IN = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"جنوری", "فروری", "مارچ", "اپریل", "مئی", "جون", "جولائی", "اگست", "ستمبر", "اکتوبر", "نومبر", "دسمبر"},
		long:  [12]string{"جنوری", "فروری", "مارچ", "اپریل", "مئی", "جون", "جولائی", "اگست", "ستمبر", "اکتوبر", "نومبر", "دسمبر"},
	},
	days: days{
		short: [7]string{"اتوار", "پیر", "بدھ", "منگل", "جمعرات", "جمعہ", "ہفتہ"},
		long:  [7]string{"اتوار", "پیر", "بدھ", "منگل", "جمعرات", "جمعہ", "ہفتہ"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday، 2 January، 2006", "2 January، 2006", "2 Jan، 2006", "2/1/06"},
		time:     [4]string{"3:04:05 PM MST -0700", "3:04:05 PM MST", "3:04:05 PM", "3:04 PM"},
		datetime: [4]string{"Monday، 2 January، 2006 3:04:05 PM MST -0700", "2 January، 2006 3:04:05 PM MST", "2 Jan، 2006 3:04:05 PM", "2/1/06 3:04 PM"},
	},
}
var l_uz = &localize{
	am:              "TO",
	pm:              "TK",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"yan", "fev", "mar", "apr", "may", "iyn", "iyl", "avg", "sen", "okt", "noy", "dek"},
		long:  [12]string{"yanvar", "fevral", "mart", "aprel", "may", "iyun", "iyul", "avgust", "sentabr", "oktabr", "noyabr", "dekabr"},
	},
	days: days{
		short: [7]string{"Yak", "Dush", "Chor", "Sesh", "Pay", "Jum", "Shan"},
		long:  [7]string{"yakshanba", "dushanba", "chorshanba", "seshanba", "payshanba", "juma", "shanba"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2-January, 2006", "2-January, 2006", "2-Jan, 2006", "02/01/06"},
		time:     [4]string{"15:04:05 (MST -0700)", "15:04:05 (MST)", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, 2-January, 2006, 15:04:05 (MST -0700)", "2-January, 2006, 15:04:05 (MST)", "2-Jan, 2006, 15:04:05", "02/01/06, 15:04"},
	},
}
var l_vi = &localize{
	am:              "SA",
	pm:              "CH",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"thg 1", "thg 2", "thg 3", "thg 4", "thg 5", "thg 6", "thg 7", "thg 8", "thg 9", "thg 10", "thg 11", "thg 12"},
		long:  [12]string{"tháng 1", "tháng 2", "tháng 3", "tháng 4", "tháng 5", "tháng 6", "tháng 7", "tháng 8", "tháng 9", "tháng 10", "tháng 11", "tháng 12"},
	},
	days: days{
		short: [7]string{"CN", "Th 2", "Th 4", "Th 3", "Th 5", "Th 6", "Th 7"},
		long:  [7]string{"Chủ Nhật", "Thứ Hai", "Thứ Tư", "Thứ Ba", "Thứ Năm", "Thứ Sáu", "Thứ Bảy"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, 2 January, 2006", "2 January, 2006", "2 Jan, 2006", "02/01/2006"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"15:04:05 MST -0700 Monday, 2 January, 2006", "15:04:05 MST 2 January, 2006", "15:04:05, 2 Jan, 2006", "15:04, 02/01/2006"},
	},
}
var l_yue = &localize{
	am:              "上午",
	pm:              "下午",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		long:  [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
	},
	days: days{
		short: [7]string{"星期日", "星期一", "星期三", "星期二", "星期四", "星期五", "星期六"},
		long:  [7]string{"星期日", "星期一", "星期三", "星期二", "星期四", "星期五", "星期六"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006年1月2日 Monday", "2006年1月2日", "2006年1月2日", "2006/1/2"},
		time:     [4]string{"PM3:04:05 [MST -0700]", "PM3:04:05 [MST]", "PM3:04:05", "PM3:04"},
		datetime: [4]string{"2006年1月2日 Monday PM3:04:05 [MST -0700]", "2006年1月2日 PM3:04:05 [MST]", "2006年1月2日 PM3:04:05", "2006/1/2 PM3:04"},
	},
}
var l_zh = &localize{
	am:              "上午",
	pm:              "下午",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		long:  [12]string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"},
	},
	days: days{
		short: [7]string{"周日", "周一", "周三", "周二", "周四", "周五", "周六"},
		long:  [7]string{"星期日", "星期一", "星期三", "星期二", "星期四", "星期五", "星期六"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006年1月2日Monday", "2006年1月2日", "2006年1月2日", "2006/1/2"},
		time:     [4]string{"MST -0700 PM3:04:05", "MST PM3:04:05", "PM3:04:05", "PM3:04"},
		datetime: [4]string{"2006年1月2日Monday MST -0700 PM3:04:05", "2006年1月2日 MST PM3:04:05", "2006年1月2日 PM3:04:05", "2006/1/2 PM3:04"},
	},
}
var l_zh_Hans_HK = &localize{
	am:              "上午",
	pm:              "下午",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		long:  [12]string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"},
	},
	days: days{
		short: [7]string{"周日", "周一", "周三", "周二", "周四", "周五", "周六"},
		long:  [7]string{"星期日", "星期一", "星期三", "星期二", "星期四", "星期五", "星期六"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006年1月2日Monday", "2006年1月2日", "2006年1月2日", "2/1/06"},
		time:     [4]string{"MST -0700 PM3:04:05", "MST PM3:04:05", "PM3:04:05", "PM3:04"},
		datetime: [4]string{"2006年1月2日Monday MST -0700 PM3:04:05", "2006年1月2日 MST PM3:04:05", "2006年1月2日 PM3:04:05", "2/1/06 PM3:04"},
	},
}
var l_zh_Hans_SG = &localize{
	am:              "上午",
	pm:              "下午",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		long:  [12]string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"},
	},
	days: days{
		short: [7]string{"周日", "周一", "周三", "周二", "周四", "周五", "周六"},
		long:  [7]string{"星期日", "星期一", "星期三", "星期二", "星期四", "星期五", "星期六"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006年1月2日Monday", "2006年1月2日", "2006年1月2日", "02/01/06"},
		time:     [4]string{"MST -0700 PM3:04:05", "MST PM3:04:05", "PM3:04:05", "PM3:04"},
		datetime: [4]string{"2006年1月2日Monday MST -0700 PM3:04:05", "2006年1月2日 MST PM3:04:05", "2006年1月2日 PM3:04:05", "02/01/06 PM3:04"},
	},
}
var l_zh_Hant = &localize{
	am:              "上午",
	pm:              "下午",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		long:  [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
	},
	days: days{
		short: [7]string{"週日", "週一", "週三", "週二", "週四", "週五", "週六"},
		long:  [7]string{"星期日", "星期一", "星期三", "星期二", "星期四", "星期五", "星期六"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006年1月2日 Monday", "2006年1月2日", "2006年1月2日", "2006/1/2"},
		time:     [4]string{"PM3:04:05 [MST -0700]", "PM3:04:05 [MST]", "PM3:04:05", "PM3:04"},
		datetime: [4]string{"2006年1月2日 Monday PM3:04:05 [MST -0700]", "2006年1月2日 PM3:04:05 [MST]", "2006年1月2日 PM3:04:05", "2006/1/2 PM3:04"},
	},
}
var l_zh_Hant_HK = &localize{
	am:              "上午",
	pm:              "下午",
	firstDay:        0,
	twentyFourHours: false,
	months: months{
		short: [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		long:  [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
	},
	days: days{
		short: [7]string{"週日", "週一", "週三", "週二", "週四", "週五", "週六"},
		long:  [7]string{"星期日", "星期一", "星期三", "星期二", "星期四", "星期五", "星期六"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"2006年1月2日Monday", "2006年1月2日", "2006年1月2日", "2/1/2006"},
		time:     [4]string{"PM3:04:05 [MST -0700]", "PM3:04:05 [MST]", "PM3:04:05", "PM3:04"},
		datetime: [4]string{"2006年1月2日Monday PM3:04:05 [MST -0700]", "2006年1月2日 PM3:04:05 [MST]", "2006年1月2日 PM3:04:05", "2/1/2006 PM3:04"},
	},
}
var l_zu = &localize{
	am:              "AM",
	pm:              "PM",
	firstDay:        1,
	twentyFourHours: false,
	months: months{
		short: [12]string{"Jan", "Feb", "Mas", "Eph", "Mey", "Jun", "Jul", "Aga", "Sep", "Okt", "Nov", "Dis"},
		long:  [12]string{"Januwari", "Februwari", "Mashi", "Ephreli", "Meyi", "Juni", "Julayi", "Agasti", "Septhemba", "Okthoba", "Novemba", "Disemba"},
	},
	days: days{
		short: [7]string{"Son", "Mso", "Tha", "Bil", "Sin", "Hla", "Mgq"},
		long:  [7]string{"ISonto", "UMsombuluko", "ULwesithathu", "ULwesibili", "ULwesine", "ULwesihlanu", "UMgqibelo"},
	},
	timeFormat: timeFormat{
		date:     [4]string{"Monday, January 2, 2006", "January 2, 2006", "Jan 2, 2006", "1/2/06"},
		time:     [4]string{"15:04:05 MST -0700", "15:04:05 MST", "15:04:05", "15:04"},
		datetime: [4]string{"Monday, January 2, 2006 15:04:05 MST -0700", "January 2, 2006 15:04:05 MST", "Jan 2, 2006 15:04:05", "1/2/06 15:04"},
	},
}
