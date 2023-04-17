package main

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Print("start")
	val := m.Run()

	log.Println("after main")
	os.Exit(val)
}

func TestA(t *testing.T) {
	log.Println("TestA running")
}

func TestB(t *testing.T) {
	log.Println("TestB running")
}
