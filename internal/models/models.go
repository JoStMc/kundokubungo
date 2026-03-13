package models

type Character struct {
	Kanji 			string
	Kaeriten 		string
	Okurigana 		string
	SecondOkurigana string
	IsSaidokumoji 	bool
	IsJukugoHead	bool
	IsJukugoTail	bool
} 

type Sentence struct {
    Characters []Character
} 

