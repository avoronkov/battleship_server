package main

type StatusMessage int

const (
	Unknown StatusMessage = iota
	WaitP1
	WaitP2
	Finish
)
