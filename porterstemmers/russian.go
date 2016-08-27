package porterstemmers

import (
    "regexp"
    "strings"
)

// RussianPorterStemmer - для россии
type RussianPorterStemmer struct {
}

func (s *RussianPorterStemmer) attemptReplacePatterns(token string, patterns []Pattern) string {
    replacement := ""      
    for i:=0; i<len(patterns);i++ {
        if patterns[i].Rx != nil {
            if patterns[i].Rx.MatchString(token) {
                replacement = patterns[i].Rx.ReplaceAllString(token, patterns[i].To)
                break
            }
        }
    }
    return replacement
}

func (s *RussianPorterStemmer) perfectiveGerund(token string) string {
    return s.attemptReplacePatterns(token,
    []Pattern{Pattern{perfectiveRx1,""}, Pattern{perfectiveRx2,""}})
}

func (s *RussianPorterStemmer) adjective(token string) string {
    return s.attemptReplacePatterns(token,
    []Pattern{Pattern{adjectiveRx,""}})
}

func (s *RussianPorterStemmer) participle(token string) string {
    return s.attemptReplacePatterns(token,
    []Pattern{Pattern{participleRx1,"$1"}, Pattern{participleRx2,""}})
}

func (s *RussianPorterStemmer) adjectival(token string) string {
    result := s.adjective(token)
    if len(result) > 0 {
        pariticipleResult := s.participle(result);
        if len(pariticipleResult) > 0 {
            result = pariticipleResult
        }
    }
    return result
}

func (s *RussianPorterStemmer) reflexive(token string) string {
    return s.attemptReplacePatterns(token,
    []Pattern{Pattern{reflexiveRx,""}})
}

func (s *RussianPorterStemmer) verb(token string) string {
    return s.attemptReplacePatterns(token,
    []Pattern{Pattern{verbRx1,"$1"}, Pattern{verbRx2,""}})
}

func (s *RussianPorterStemmer) noun(token string) string {
    return s.attemptReplacePatterns(token,
    []Pattern{Pattern{nounRx,""}})
}

func (s *RussianPorterStemmer) superlative(token string) string {
    return s.attemptReplacePatterns(token,
    []Pattern{Pattern{superlativeRx,""}})
}

func (s *RussianPorterStemmer) derivational(token string) string {
    return s.attemptReplacePatterns(token,
    []Pattern{Pattern{derivationalRx,""}})
}

// StemString - stem string
func (s *RussianPorterStemmer) StemString(token string) string {
    token = strings.TrimSpace(strings.ToLower(token))
    token = eeRx.ReplaceAllString(token, "е")
    rv := volwesRx.FindAllString(token,-1)
    if rv == nil || len(rv) < 3 {
        return token
    } 
    head := rv[1]
    r2 := volwesRx.FindAllString(rv[2],0)
    result := s.perfectiveGerund(rv[2])
    if len(result) == 0 {
        resultReflexive := s.reflexive(rv[2])
        if len(resultReflexive) == 0 {
            resultReflexive = rv[2]
        }
        result = s.adjectival(resultReflexive)
        if len(result) == 0 {
            result = s.verb(resultReflexive)
            if len(result) == 0 {
                result = s.noun(resultReflexive);
                if len(result) == 0 {
                    result = resultReflexive
                }
            }
        }
    }
    result = andRx.ReplaceAllString(result, "")
    derivationalResult := result
    if r2 != nil && len(r2) > 2 && len(r2[2]) > 0 {
        derivationalResult = s.derivational(r2[2]);
        if len(derivationalResult) == 0 {
            derivationalResult = s.derivational(result)
        } else {
            derivationalResult = result
        }
    }
    superlativeResult := s.superlative(derivationalResult)
    if len(superlativeResult) == 0 {
        superlativeResult = derivationalResult
    }
    superlativeResult = wnRx.ReplaceAllString(superlativeResult, "$1")
    superlativeResult = mzRx.ReplaceAllString(superlativeResult, "")
    return head + superlativeResult
}

// RX List
var (
    perfectiveRx1  = regexp.MustCompile("[ая]в(ши|шись)$")
    perfectiveRx2  = regexp.MustCompile("(ив|ивши|ившись|ывши|ывшись|ыв)$")
    adjectiveRx    = regexp.MustCompile("(ее|ие|ые|ое|ими|ыми|ей|ий|ый|ой|ем|им|ым|ом|его|ого|ему|ому|их|ых|ую|юю|ая|яя|ою|ею)$")
    participleRx1  = regexp.MustCompile("([ая])(ем|нн|вш|ющ|щ)$")
    participleRx2  = regexp.MustCompile("(ивш|ывш|ующ)$")
    reflexiveRx    = regexp.MustCompile("(ся|сь)$")
    verbRx1        = regexp.MustCompile("([ая])(ла|на|ете|йте|ли|й|л|ем|н|ло|но|ет|ют|ны|ть|ешь|нно)$")
    verbRx2        = regexp.MustCompile("(ила|ыла|ена|ейте|уйте|ите|или|ыли|ей|уй|ил|ыл|им|ым|ен|ило|ыло|ено|ят|ует|ит|ыт|ены|ить|ыть|ишь|ую|ю)$")
    nounRx         = regexp.MustCompile("(а|ев|ов|ие|ье|е|иями|ями|ами|еи|ии|и|ией|ей|ой|ий|й|иям|ям|ием|ем|ам|ом|о|у|ах|иях|ях|ы|ь|ию|ью|ю|ия|ья|я)$")
    superlativeRx  = regexp.MustCompile("(ейш|ейше)$")
    derivationalRx = regexp.MustCompile("(ост|ость)$")
    eeRx           = regexp.MustCompile("ё")
    volwesRx       = regexp.MustCompile("^(.*?[аеиоюяуыиэ])(.*)$")
    andRx          = regexp.MustCompile("и$")
    wnRx           = regexp.MustCompile("(н)н")
    mzRx           = regexp.MustCompile("ь$")
)