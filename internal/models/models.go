package models

type Character struct {
	Kanji 			string
	Kaeriten 		string
	Okurigana 		string
	SecondOkurigana string
	IsSaidokumoji 	bool
	IsJukugo		bool
} 

type Sentence struct {
    Characters []Character
} 

