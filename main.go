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

func inSlice(slice []*Room, val *Room) bool {
    for i := 0; i < len(slice); i++ {
        if slice[i] == val {
            return true
        }
    }
    return false
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

func getLinks(filename string) [][]string {
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
	return links
}

func getAdresses(rooms []Room, links [][]string) []*Room { 
	var adresses []*Room
	for i := 0; i < len(links); i++ {
		for j := 0; j < len(rooms); j++ {
			if links[i][0] == rooms[j].Name {
				adresses = append(adresses, &rooms[j])
			}
			if links[i][1] == rooms[j].Name  {
				adresses = append(adresses, &rooms[j])
			}
		}
	}

	for i := 0; i < len(adresses)-1; i++{
		adresses[i].Nxt = adresses[i+1]
	}
	return adresses
}


func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		file, _ := os.Open(args[0])
		rooms := createRooms(file)
		fmt.Println(rooms)
		l := getLinks(args[0])
		a := getAdresses(rooms, l)
		fmt.Println(a)
		fmt.Println(rooms)
		r := rooms[0]
	}
}
