package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type Player struct {
	exe  *exec.Cmd
	name string
	in   io.WriteCloser
	out  io.ReadCloser
}

func NewPlayer(name, executable string) *Player {
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

func (p *Player) InitField() *Field {
	io.WriteString(p.in, "Arrange!\n")
	f := ReadField(p.out)
	log.Printf("[%s] InitField:\n%s\n", p.name, f)
	return f
}

func (p *Player) ShootCmd() {
	log.Printf("[%s] ShootCmd\n", p.name)
	io.WriteString(p.in, "Shoot!\n")
}

func (p *Player) GetShot() *Shot {
	log.Printf("[%s] GetShot\n", p.name)
	shot := ReadShot(p.out)
	log.Printf("[%s] GetShot: %s", p.name, *shot)
	return shot
}

func (p *Player) SendResult(c cell) {
	switch c {
	case MISS:
		io.WriteString(p.in, "Miss\n")
	case HIT:
		io.WriteString(p.in, "Hit\n")
	case KILL:
		io.WriteString(p.in, "Kill\n")
	default:
		panic(fmt.Errorf("Unexpected result: %s", c))
	}
}

func (p *Player) EnemyShootedInto(s Shot) {
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

func (p Player) Name() string {
	return p.name
}

func (p *Player) Close() (err error) {
	log.Printf("[%s] Close", p.name)
	if err = p.in.Close(); err != nil {
		return
	}
	err = p.out.Close()
	p.exe.Wait()
	log.Printf("[%s] Ok", p.name)
	return
}
