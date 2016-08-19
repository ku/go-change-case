package changecase

import (
	"bytes"
	"strings"
	"unicode"
)

const (
	INITIAL_STATE                 = iota
	EXPECT_FOLLOWING_SMALL_LETTER = iota
	IN_CONSECUTIVE_CAPITALS       = iota
	IN_WORD                       = iota
	SEEK_FOR_NEXT_WORD            = iota
)

type CaseTranslator struct {
	FirstLetter       func(rune) rune
	LetterInWord      func(rune) rune
	FirstLetterOfWord func(rune) rune
	Separator         rune
}

type processor struct {
	state        int
	buffer       *bytes.Buffer
	tr           *CaseTranslator
	bufferedRune rune
}

func passThrough(r rune) rune {
	return r
}

func NewProcessor(t *CaseTranslator) *processor {
	p := new(processor)
	p.state = INITIAL_STATE
	p.buffer = bytes.NewBuffer(nil)
	p.tr = t
	return p
}

func (p *processor) flushRuneBuffer() {
	if p.bufferedRune != 0 {
		p.writeRune(p.bufferedRune)
	}
}
func (p *processor) putCharInRuneBuffer(r rune) {
	if p.bufferedRune != 0 {
		p.charInWord(p.bufferedRune)
	}
	p.bufferedRune = r
}
func (p *processor) writeRune(r rune) {
	p.buffer.WriteRune(r)
}
func (p *processor) firstLetter(r rune) {
	p.writeRune(p.tr.FirstLetter(r))
	if unicode.IsUpper(r) {
		p.state = EXPECT_FOLLOWING_SMALL_LETTER
	} else {
		p.state = IN_WORD
	}
}
func (p *processor) foundNewWord(r rune) {
	if p.tr.Separator != 0 {
		p.writeRune(p.tr.Separator)
	}
	p.writeRune(p.tr.FirstLetterOfWord(r))

	if unicode.IsUpper(r) {
		p.state = EXPECT_FOLLOWING_SMALL_LETTER
	} else {
		p.state = IN_WORD
	}
}

func (p *processor) charInWord(r rune) {
	r = p.tr.LetterInWord(r)
	p.writeRune(r)
}
func (p *processor) firstLetterOfWord(r rune) {
	r = p.tr.FirstLetterOfWord(r)
	p.writeRune(r)
}

func (p *processor) convert(s string) string {
	p.buffer.Grow(len(s))
	for _, r := range s {
		isNumber := unicode.Is(unicode.Number, r)
		isWord := unicode.Is(unicode.Letter, r) || isNumber

		switch p.state {
		case INITIAL_STATE:
			if isWord {
				p.firstLetter(r)
			}
		case EXPECT_FOLLOWING_SMALL_LETTER:
			if isWord {
				if unicode.IsUpper(r) {
					p.putCharInRuneBuffer(r)
					p.state = IN_CONSECUTIVE_CAPITALS
				} else {
					p.flushRuneBuffer()
					p.charInWord(r)
					p.state = IN_WORD
				}
			} else {
				p.putCharInRuneBuffer(0)
				p.state = SEEK_FOR_NEXT_WORD
			}
		case IN_CONSECUTIVE_CAPITALS:
			if isWord {
				if unicode.IsUpper(r) || isNumber {
					p.putCharInRuneBuffer(r)
				} else {
					p.foundNewWord(p.bufferedRune)
					p.bufferedRune = 0
					p.charInWord(r)
					p.state = IN_WORD
				}
			} else {
				p.putCharInRuneBuffer(0)
				p.state = SEEK_FOR_NEXT_WORD
			}
		case IN_WORD:
			if isWord {
				if unicode.IsUpper(r) {
					p.foundNewWord(r)
				} else {
					p.charInWord(r)
				}
			} else {
				p.state = SEEK_FOR_NEXT_WORD
			}
		case SEEK_FOR_NEXT_WORD:
			if isWord {
				p.foundNewWord(r)
			}
		}
	}
	if p.bufferedRune != 0 {
		p.charInWord(p.bufferedRune)
	}
	return p.buffer.String()
}

func (p *processor) Convert(s string) string {
	return p.convert(s)
}

func NewLowerProcessor(separator rune) *processor {
	return NewProcessor(&CaseTranslator{
		FirstLetter:       unicode.ToLower,
		LetterInWord:      unicode.ToLower,
		FirstLetterOfWord: unicode.ToLower,
		Separator:         separator,
	})
}

func Camel(s string) string {
	return NewProcessor(&CaseTranslator{
		FirstLetter:       unicode.ToLower,
		LetterInWord:      unicode.ToLower,
		FirstLetterOfWord: unicode.ToUpper,
	}).Convert(s)
}

func Constant(s string) string {
	return NewProcessor(&CaseTranslator{
		FirstLetter:       unicode.ToUpper,
		LetterInWord:      unicode.ToUpper,
		FirstLetterOfWord: unicode.ToUpper,
		Separator:         '_',
	}).Convert(s)
}
func Dot(s string) string {
	return NewLowerProcessor('.').Convert(s)
}

func Header(s string) string {
	return NewProcessor(&CaseTranslator{
		FirstLetter:       unicode.ToUpper,
		LetterInWord:      unicode.ToLower,
		FirstLetterOfWord: unicode.ToUpper,
		Separator:         '-',
	}).Convert(s)
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}
func Lower(s string) string {
	return strings.ToLower(s)
}

func LcFirst(s string) string {
	return strings.ToLower(s[0:1]) + s[1:len(s)]
}

func No(s string) string {
	return NewProcessor(&CaseTranslator{
		FirstLetter:       unicode.ToLower,
		LetterInWord:      unicode.ToLower,
		FirstLetterOfWord: unicode.ToLower,
		Separator:         ' ',
	}).Convert(s)
}

func Param(s string) string {
	return NewLowerProcessor('-').Convert(s)
}
func Pascal(s string) string {
	return NewProcessor(&CaseTranslator{
		FirstLetter:       unicode.ToUpper,
		LetterInWord:      unicode.ToLower,
		FirstLetterOfWord: unicode.ToUpper,
	}).Convert(s)
}

func Path(s string) string {
	return NewLowerProcessor('/').Convert(s)
}

func Sentence(s string) string {
	return NewProcessor(&CaseTranslator{
		FirstLetter:       unicode.ToUpper,
		LetterInWord:      unicode.ToLower,
		FirstLetterOfWord: unicode.ToLower,
		Separator:         ' ',
	}).Convert(s)
}
func Snake(s string) string {
	return NewLowerProcessor('_').Convert(s)
}

func Swap(s string) string {
	b := bytes.NewBuffer(nil)
	b.Grow(len(s))
	for _, r := range s {
		if unicode.IsUpper(r) {
			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteRune(unicode.ToUpper(r))
		}
	}
	return b.String()
}

func Title(s string) string {
	return NewProcessor(&CaseTranslator{
		FirstLetter:       unicode.ToUpper,
		LetterInWord:      unicode.ToLower,
		FirstLetterOfWord: unicode.ToUpper,
		Separator:         ' ',
	}).Convert(s)
}

func Upper(s string) string {
	return strings.ToUpper(s)
}

func UcFirst(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:len(s)]
}
