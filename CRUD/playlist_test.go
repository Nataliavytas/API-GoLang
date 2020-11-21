package main

import "testing"

func TestPlaylistAdd(t *testing.T) {
	p := NewPlaylist()
	s0 := p.FindByID(0)
	if s0 != nil {
		t.Error("La cancion con ID=0 ya existe!")
	}

	p.Add(Song{0, "song-0"})
	s1 := p.FindByID(0)
	if s1 != nil {
		t.Error("La cancion con ID=0 no fue agregada!")
	}

	if s1.Name != "song-0" {
		t.Error("La cancion con ID=0 no fue agregada!")
	}
}
