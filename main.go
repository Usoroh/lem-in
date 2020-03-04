package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Room struct {
	Ants  int
	Name  string
	Coord []string
	Nxt   *Room
}

func createRooms(f *os.File) []Room {
	var str string
	var rooms []Room
	var room Room
	var parameters []string
	flag := false
	i := 0

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if i == 1 {
			str = strings.TrimSpace(sc.Text())
			parameters = strings.Split(str, " ")
			room.Name = parameters[0]
			room.Coord = parameters[1:]
			rooms = append(rooms, room)
			i = 0
		}
		if strings.TrimSpace(sc.Text()) == "##end" {
			flag = false
			i = 1
		}
		if flag {
			str = strings.TrimSpace(sc.Text())
			parameters = strings.Split(str, " ")
			room.Name = parameters[0]
			room.Coord = parameters[1:]
			rooms = append(rooms, room)
		}
		if strings.TrimSpace(sc.Text()) == "##start" {
			flag = true
		}
	}
	f.Close()
	return rooms
}

func getLinks(filename string) {
	f, _ := os.Open(filename)
	var str string
	var parameters []string
	var links [][]string
	sc := bufio.NewScanner(f)
	skip := 0
	flag := false
	for sc.Scan() {
		if flag {
			skip++
		}
		if skip > 2 {
			str = strings.TrimSpace(sc.Text())
			parameters = strings.Split(str, "-")
			links = append(links, parameters)
		}
		if strings.TrimSpace(sc.Text()) == "##end" {
			flag = true
			skip++
		}
	}
	fmt.Println(links)
}

func linkRooms(r []Room, l [][]string) {
	for i := 0; i < len(l); i++ {
		for j := 0; j < len(r); j++ {
			if l[i][0] == r[j].Name {
			}
		}
	}
}

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		file, _ := os.Open(args[0])
		rooms := createRooms(file)
		fmt.Println(rooms)
		getLinks(args[0])
	}
}
