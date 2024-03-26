package main

import "fmt"

type Player interface {
	Play()
}

type MP3Player struct{}

func (p *MP3Player) Play() {
	fmt.Println("Playing MP3...")
}

type OldRecordPlayer struct{}

func (p *OldRecordPlayer) PlayRecord() {
	fmt.Println("Playing record...")
}

type RecordAdapter struct {
	recordPlayer *OldRecordPlayer
}

func NewRecordAdapter(player *OldRecordPlayer) *RecordAdapter {
	return &RecordAdapter{
		recordPlayer: player,
	}
}

func (ra *RecordAdapter) Play() {
	ra.recordPlayer.PlayRecord()
}
func main() {
	mp3Player := &MP3Player{}
	recordPlayer := &OldRecordPlayer{}

	mp3Player.Play()

	recordAdapter := NewRecordAdapter(recordPlayer)
	recordAdapter.Play()
}
