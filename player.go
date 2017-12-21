package main

import (
	"battleship_server/cell"
	"battleship_server/field"
	bio "battleship_server/io"
	"battleship_server/shot"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type PlayerName struct {
	Name string
	N    int
}

func (pn PlayerName) String() string {
	return fmt.Sprintf("%s (%d)", pn.Name, pn.N)
}

type Player struct {
	exe  *exec.Cmd
	name PlayerName
	in   io.WriteCloser
	out  io.ReadCloser
}

func NewPlayer(name PlayerName, executable string) *Player {
	p := new(Player)
	p.name = name
	p.exe = exec.Command(executable)
	p.exe.Stderr = os.Stderr
	var err error
	if p.in, err = p.exe.StdinPipe(); err != nil {
		panic(err)
	}
	if p.out, err = p.exe.StdoutPipe(); err != nil {
		panic(err)
	}
	if err = p.exe.Start(); err != nil {
		panic(err)
	}
	return p
}

func (p *Player) InitField() (*field.Field, error) {
	io.WriteString(p.in, "Arrange!\n")
	f, err := bio.ReadField(p.out)
	log.Printf("[%s] InitField:\n%s\n(%v)", p.name, f, err)
	return f, err
}

func (p *Player) ShootCmd() {
	log.Printf("[%s] ShootCmd\n", p.name)
	io.WriteString(p.in, "Shoot!\n")
}

func (p *Player) GetShot() (*shot.Shot, error) {
	log.Printf("[%s] GetShot\n", p.name)
	shot, err := bio.ReadShot(p.out)
	log.Printf("[%s] GetShot: %s (%s)", p.name, shot, err)
	return shot, err
}

func (p *Player) SendResult(c cell.Cell) {
	switch c {
	case cell.MISS:
		io.WriteString(p.in, "Miss\n")
	case cell.HIT:
		io.WriteString(p.in, "Hit\n")
	case cell.KILL:
		io.WriteString(p.in, "Kill\n")
	default:
		panic(fmt.Errorf("Unexpected result: %s", c))
	}
}

func (p *Player) EnemyShootedInto(s shot.Shot) {
	fmt.Fprintf(p.in, "Enemy shooted into %s", s)
}

func (p *Player) Win() {
	log.Printf("[%s] Win!", p.name)
	io.WriteString(p.in, "Win!\n")
}
func (p *Player) Lose() {
	log.Printf("[%s] Lose!", p.name)
	io.WriteString(p.in, "Lose\n")
}

func (p Player) Name() PlayerName {
	return p.name
}

func (p *Player) Close() (err error) {
	log.Printf("[%s] Close", p.name)
	if err = p.in.Close(); err != nil {
		return
	}
	p.out.Close()
	e := p.exe.Wait()
	log.Printf("[%s] Closed (%s)", p.name, e)
	return
}
