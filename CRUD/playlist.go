package main

import "fmt"

type Playlist struct {
	songs map[int]*Song
}

type Song struct {
	ID   int
	Name string
}

func NewPlaylist() Playlist {
	songs := make(map[int]*Song)
	return Playlist{
		songs,
	}
}

func (p Playlist) Add(s Song) {
	p.songs[s.ID] = &s
}

// Print...
func (p Playlist) Print() {
	for _, v := range p.songs {
		fmt.Printf("[%v]\t%v\n", v.ID, v.Name)
	}
}

func (p Playlist) FindByID(ID int) *Song {
	return p.songs[ID]
}

func (p Playlist) Delete(ID int) {
	delete(p.songs, ID)
}

func (p Playlist) Update(s Song) {
	p.songs[s.ID] = &s
}
