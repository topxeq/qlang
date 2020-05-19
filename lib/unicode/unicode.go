package unicode

import (
	"unicode"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "unicode",

	"LowerCase":       unicode.LowerCase,
	"MaxASCII":        unicode.MaxASCII,
	"MaxCase":         unicode.MaxCase,
	"MaxLatin1":       unicode.MaxLatin1,
	"MaxRune":         unicode.MaxRune,
	"ReplacementChar": unicode.ReplacementChar,
	"TitleCase":       unicode.TitleCase,
	"UpperCase":       unicode.UpperCase,
	"UpperLower":      unicode.UpperLower,
	"Version":         unicode.Version,

	"ASCII_Hex_Digit":                    unicode.ASCII_Hex_Digit,
	"Adlam":                              unicode.Adlam,
	"Ahom":                               unicode.Ahom,
	"Anatolian_Hieroglyphs":              unicode.Anatolian_Hieroglyphs,
	"Arabic":                             unicode.Arabic,
	"Armenian":                           unicode.Armenian,
	"Avestan":                            unicode.Avestan,
	"AzeriCase":                          unicode.AzeriCase,
	"Balinese":                           unicode.Balinese,
	"Bamum":                              unicode.Bamum,
	"Bassa_Vah":                          unicode.Bassa_Vah,
	"Batak":                              unicode.Batak,
	"Bengali":                            unicode.Bengali,
	"Bhaiksuki":                          unicode.Bhaiksuki,
	"Bidi_Control":                       unicode.Bidi_Control,
	"Bopomofo":                           unicode.Bopomofo,
	"Brahmi":                             unicode.Brahmi,
	"Braille":                            unicode.Braille,
	"Buginese":                           unicode.Buginese,
	"Buhid":                              unicode.Buhid,
	"C":                                  unicode.C,
	"Canadian_Aboriginal":                unicode.Canadian_Aboriginal,
	"Carian":                             unicode.Carian,
	"CaseRanges":                         unicode.CaseRanges,
	"Categories":                         unicode.Categories,
	"Caucasian_Albanian":                 unicode.Caucasian_Albanian,
	"Cc":                                 unicode.Cc,
	"Cf":                                 unicode.Cf,
	"Chakma":                             unicode.Chakma,
	"Cham":                               unicode.Cham,
	"Cherokee":                           unicode.Cherokee,
	"Co":                                 unicode.Co,
	"Common":                             unicode.Common,
	"Coptic":                             unicode.Coptic,
	"Cs":                                 unicode.Cs,
	"Cuneiform":                          unicode.Cuneiform,
	"Cypriot":                            unicode.Cypriot,
	"Cyrillic":                           unicode.Cyrillic,
	"Dash":                               unicode.Dash,
	"Deprecated":                         unicode.Deprecated,
	"Deseret":                            unicode.Deseret,
	"Devanagari":                         unicode.Devanagari,
	"Diacritic":                          unicode.Diacritic,
	"Digit":                              unicode.Digit,
	"Dogra":                              unicode.Dogra,
	"Duployan":                           unicode.Duployan,
	"Egyptian_Hieroglyphs":               unicode.Egyptian_Hieroglyphs,
	"Elbasan":                            unicode.Elbasan,
	"Elymaic":                            unicode.Elymaic,
	"Ethiopic":                           unicode.Ethiopic,
	"Extender":                           unicode.Extender,
	"FoldCategory":                       unicode.FoldCategory,
	"FoldScript":                         unicode.FoldScript,
	"Georgian":                           unicode.Georgian,
	"Glagolitic":                         unicode.Glagolitic,
	"Gothic":                             unicode.Gothic,
	"Grantha":                            unicode.Grantha,
	"GraphicRanges":                      unicode.GraphicRanges,
	"Greek":                              unicode.Greek,
	"Gujarati":                           unicode.Gujarati,
	"Gunjala_Gondi":                      unicode.Gunjala_Gondi,
	"Gurmukhi":                           unicode.Gurmukhi,
	"Han":                                unicode.Han,
	"Hangul":                             unicode.Hangul,
	"Hanifi_Rohingya":                    unicode.Hanifi_Rohingya,
	"Hanunoo":                            unicode.Hanunoo,
	"Hatran":                             unicode.Hatran,
	"Hebrew":                             unicode.Hebrew,
	"Hex_Digit":                          unicode.Hex_Digit,
	"Hiragana":                           unicode.Hiragana,
	"Hyphen":                             unicode.Hyphen,
	"IDS_Binary_Operator":                unicode.IDS_Binary_Operator,
	"IDS_Trinary_Operator":               unicode.IDS_Trinary_Operator,
	"Ideographic":                        unicode.Ideographic,
	"Imperial_Aramaic":                   unicode.Imperial_Aramaic,
	"Inherited":                          unicode.Inherited,
	"Inscriptional_Pahlavi":              unicode.Inscriptional_Pahlavi,
	"Inscriptional_Parthian":             unicode.Inscriptional_Parthian,
	"Javanese":                           unicode.Javanese,
	"Join_Control":                       unicode.Join_Control,
	"Kaithi":                             unicode.Kaithi,
	"Kannada":                            unicode.Kannada,
	"Katakana":                           unicode.Katakana,
	"Kayah_Li":                           unicode.Kayah_Li,
	"Kharoshthi":                         unicode.Kharoshthi,
	"Khmer":                              unicode.Khmer,
	"Khojki":                             unicode.Khojki,
	"Khudawadi":                          unicode.Khudawadi,
	"L":                                  unicode.L,
	"Lao":                                unicode.Lao,
	"Latin":                              unicode.Latin,
	"Lepcha":                             unicode.Lepcha,
	"Letter":                             unicode.Letter,
	"Limbu":                              unicode.Limbu,
	"Linear_A":                           unicode.Linear_A,
	"Linear_B":                           unicode.Linear_B,
	"Lisu":                               unicode.Lisu,
	"Ll":                                 unicode.Ll,
	"Lm":                                 unicode.Lm,
	"Lo":                                 unicode.Lo,
	"Logical_Order_Exception":            unicode.Logical_Order_Exception,
	"Lower":                              unicode.Lower,
	"Lt":                                 unicode.Lt,
	"Lu":                                 unicode.Lu,
	"Lycian":                             unicode.Lycian,
	"Lydian":                             unicode.Lydian,
	"M":                                  unicode.M,
	"Mahajani":                           unicode.Mahajani,
	"Makasar":                            unicode.Makasar,
	"Malayalam":                          unicode.Malayalam,
	"Mandaic":                            unicode.Mandaic,
	"Manichaean":                         unicode.Manichaean,
	"Marchen":                            unicode.Marchen,
	"Mark":                               unicode.Mark,
	"Masaram_Gondi":                      unicode.Masaram_Gondi,
	"Mc":                                 unicode.Mc,
	"Me":                                 unicode.Me,
	"Medefaidrin":                        unicode.Medefaidrin,
	"Meetei_Mayek":                       unicode.Meetei_Mayek,
	"Mende_Kikakui":                      unicode.Mende_Kikakui,
	"Meroitic_Cursive":                   unicode.Meroitic_Cursive,
	"Meroitic_Hieroglyphs":               unicode.Meroitic_Hieroglyphs,
	"Miao":                               unicode.Miao,
	"Mn":                                 unicode.Mn,
	"Modi":                               unicode.Modi,
	"Mongolian":                          unicode.Mongolian,
	"Mro":                                unicode.Mro,
	"Multani":                            unicode.Multani,
	"Myanmar":                            unicode.Myanmar,
	"N":                                  unicode.N,
	"Nabataean":                          unicode.Nabataean,
	"Nandinagari":                        unicode.Nandinagari,
	"Nd":                                 unicode.Nd,
	"New_Tai_Lue":                        unicode.New_Tai_Lue,
	"Newa":                               unicode.Newa,
	"Nko":                                unicode.Nko,
	"Nl":                                 unicode.Nl,
	"No":                                 unicode.No,
	"Noncharacter_Code_Point":            unicode.Noncharacter_Code_Point,
	"Number":                             unicode.Number,
	"Nushu":                              unicode.Nushu,
	"Nyiakeng_Puachue_Hmong":             unicode.Nyiakeng_Puachue_Hmong,
	"Ogham":                              unicode.Ogham,
	"Ol_Chiki":                           unicode.Ol_Chiki,
	"Old_Hungarian":                      unicode.Old_Hungarian,
	"Old_Italic":                         unicode.Old_Italic,
	"Old_North_Arabian":                  unicode.Old_North_Arabian,
	"Old_Permic":                         unicode.Old_Permic,
	"Old_Persian":                        unicode.Old_Persian,
	"Old_Sogdian":                        unicode.Old_Sogdian,
	"Old_South_Arabian":                  unicode.Old_South_Arabian,
	"Old_Turkic":                         unicode.Old_Turkic,
	"Oriya":                              unicode.Oriya,
	"Osage":                              unicode.Osage,
	"Osmanya":                            unicode.Osmanya,
	"Other":                              unicode.Other,
	"Other_Alphabetic":                   unicode.Other_Alphabetic,
	"Other_Default_Ignorable_Code_Point": unicode.Other_Default_Ignorable_Code_Point,
	"Other_Grapheme_Extend":              unicode.Other_Grapheme_Extend,
	"Other_ID_Continue":                  unicode.Other_ID_Continue,
	"Other_ID_Start":                     unicode.Other_ID_Start,
	"Other_Lowercase":                    unicode.Other_Lowercase,
	"Other_Math":                         unicode.Other_Math,
	"Other_Uppercase":                    unicode.Other_Uppercase,
	"P":                                  unicode.P,
	"Pahawh_Hmong":                       unicode.Pahawh_Hmong,
	"Palmyrene":                          unicode.Palmyrene,
	"Pattern_Syntax":                     unicode.Pattern_Syntax,
	"Pattern_White_Space":                unicode.Pattern_White_Space,
	"Pau_Cin_Hau":                        unicode.Pau_Cin_Hau,
	"Pc":                                 unicode.Pc,
	"Pd":                                 unicode.Pd,
	"Pe":                                 unicode.Pe,
	"Pf":                                 unicode.Pf,
	"Phags_Pa":                           unicode.Phags_Pa,
	"Phoenician":                         unicode.Phoenician,
	"Pi":                                 unicode.Pi,
	"Po":                                 unicode.Po,
	"Prepended_Concatenation_Mark":       unicode.Prepended_Concatenation_Mark,
	"PrintRanges":                        unicode.PrintRanges,
	"Properties":                         unicode.Properties,
	"Ps":                                 unicode.Ps,
	"Psalter_Pahlavi":                    unicode.Psalter_Pahlavi,
	"Punct":                              unicode.Punct,
	"Quotation_Mark":                     unicode.Quotation_Mark,
	"Radical":                            unicode.Radical,
	"Regional_Indicator":                 unicode.Regional_Indicator,
	"Rejang":                             unicode.Rejang,
	"Runic":                              unicode.Runic,
	"S":                                  unicode.S,
	"STerm":                              unicode.STerm,
	"Samaritan":                          unicode.Samaritan,
	"Saurashtra":                         unicode.Saurashtra,
	"Sc":                                 unicode.Sc,
	"Scripts":                            unicode.Scripts,
	"Sentence_Terminal":                  unicode.Sentence_Terminal,
	"Sharada":                            unicode.Sharada,
	"Shavian":                            unicode.Shavian,
	"Siddham":                            unicode.Siddham,
	"SignWriting":                        unicode.SignWriting,
	"Sinhala":                            unicode.Sinhala,
	"Sk":                                 unicode.Sk,
	"Sm":                                 unicode.Sm,
	"So":                                 unicode.So,
	"Soft_Dotted":                        unicode.Soft_Dotted,
	"Sogdian":                            unicode.Sogdian,
	"Sora_Sompeng":                       unicode.Sora_Sompeng,
	"Soyombo":                            unicode.Soyombo,
	"Space":                              unicode.Space,
	"Sundanese":                          unicode.Sundanese,
	"Syloti_Nagri":                       unicode.Syloti_Nagri,
	"Symbol":                             unicode.Symbol,
	"Syriac":                             unicode.Syriac,
	"Tagalog":                            unicode.Tagalog,
	"Tagbanwa":                           unicode.Tagbanwa,
	"Tai_Le":                             unicode.Tai_Le,
	"Tai_Tham":                           unicode.Tai_Tham,
	"Tai_Viet":                           unicode.Tai_Viet,
	"Takri":                              unicode.Takri,
	"Tamil":                              unicode.Tamil,
	"Tangut":                             unicode.Tangut,
	"Telugu":                             unicode.Telugu,
	"Terminal_Punctuation":               unicode.Terminal_Punctuation,
	"Thaana":                             unicode.Thaana,
	"Thai":                               unicode.Thai,
	"Tibetan":                            unicode.Tibetan,
	"Tifinagh":                           unicode.Tifinagh,
	"Tirhuta":                            unicode.Tirhuta,
	"Title":                              unicode.Title,
	"TurkishCase":                        unicode.TurkishCase,
	"Ugaritic":                           unicode.Ugaritic,
	"Unified_Ideograph":                  unicode.Unified_Ideograph,
	"Upper":                              unicode.Upper,
	"Vai":                                unicode.Vai,
	"Variation_Selector":                 unicode.Variation_Selector,
	"Wancho":                             unicode.Wancho,
	"Warang_Citi":                        unicode.Warang_Citi,
	"White_Space":                        unicode.White_Space,
	"Yi":                                 unicode.Yi,
	"Z":                                  unicode.Z,
	"Zanabazar_Square":                   unicode.Zanabazar_Square,
	"Zl":                                 unicode.Zl,
	"Zp":                                 unicode.Zp,
	"Zs":                                 unicode.Zs,

	"In":         unicode.In,
	"Is":         unicode.Is,
	"IsControl":  unicode.IsControl,
	"IsDigit":    unicode.IsDigit,
	"IsGraphic":  unicode.IsGraphic,
	"IsLetter":   unicode.IsLetter,
	"IsLower":    unicode.IsLower,
	"IsMark":     unicode.IsMark,
	"IsNumber":   unicode.IsNumber,
	"IsOneOf":    unicode.IsOneOf,
	"IsPrint":    unicode.IsPrint,
	"IsPunct":    unicode.IsPunct,
	"IsSpace":    unicode.IsSpace,
	"IsSymbol":   unicode.IsSymbol,
	"IsTitle":    unicode.IsTitle,
	"IsUpper":    unicode.IsUpper,
	"SimpleFold": unicode.SimpleFold,
	"To":         unicode.To,
	"ToLower":    unicode.ToLower,
	"ToTitle":    unicode.ToTitle,
	"ToUpper":    unicode.ToUpper,

	"CaseRange":  qlang.StructOf((*unicode.CaseRange)(nil)),
	"Range16":    qlang.StructOf((*unicode.Range16)(nil)),
	"Range32":    qlang.StructOf((*unicode.Range32)(nil)),
	"RangeTable": qlang.StructOf((*unicode.RangeTable)(nil)),
}
