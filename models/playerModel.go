package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	UserId        uint
	User          User
	CharacterName string
	GameId        uint
	Game          Game
}

type PlayerResoureces struct {
	gorm.Model
	PlayerId uint
	Player   Player
	BookId   uint
	Book     Book
	NoteId   uint
	Note     Note
}

type Book struct {
	gorm.Model
	Name        string
	Language    string
	TimeToRead  uint8
	IsCompleat  bool
	WhereFind   string
	Description string
}

type BookContent struct {
	gorm.Model
	BookId  uint
	Book    Book
	Content string
	SpellId uint
	Spell   Spell
}

type Spell struct {
	gorm.Model
	Name        string
	Description string
}

type Note struct {
	gorm.Model
	Name    string
	Content string
}
