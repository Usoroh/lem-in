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
	Nxt   []*Room
}

var start *Room
var end *Room

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
	return links
}

func linkRooms(rooms []Room, links [][]string) {
	var linkFrom *Room
	var linkTo *Room
	for i := 0; i < len(links); i++ {
		for j := 0; j < len(links[i]); j++ {
			for k := 0; k < len(rooms); k++ {
				if rooms[k].Name == links[i][0] {
					linkFrom = &rooms[k]
				}
				if rooms[k].Name == links[i][1] {
					linkTo = &rooms[k]
				}
			}
		}
		linkFrom.Nxt = append(linkFrom.Nxt, linkTo)
		fmt.Println("Link from: ", linkFrom.Name)
		fmt.Println("Link to: ", linkTo.Name)
	}
}

func isVisited(rooms []string, name string) bool {
	for i := 0; i < len(rooms); i++ {
		if name == rooms[i] {
			return true
		}
	}
	return false
}

func isParent(r Room, c Room) bool {
	for i := 0; i < len(r.Nxt); i++ {
		fmt.Println("checking", *r.Nxt[i])
		if r.Nxt[i].Name == c.Name {
			return true
		}
	}
	return false
}

func clearPath(p []Room) []Room {
	var finalPath []Room
	for i := len(p)-1; i > 0; i-- {
		if isParent(p[i-1], p[i]) {
			finalPath = append(finalPath, p[i])
		}
	}
	return finalPath
}

func bfs(rooms []Room) {
	var q []Room
	var visited []string
	end := rooms[len(rooms)-1]
	q = append(q, rooms[0])
	// visited = append(visited, crt.Name)
	var path []Room
	//visit parent and its children
	for len(q) > 0 {
	
		crt := q[0]
		fmt.Println(crt.Name)
		visited = append(visited, crt.Name)
		path = append(path, crt)
		q = q[1:]
		for i := 0; i < len(crt.Nxt); i++ {
			if crt.Nxt[i].Name == end.Name {
				path = append(path, end)
				fmt.Println(path)
				finPath := clearPath(path)
				fmt.Println(finPath)
				return
			}
			if !isVisited(visited, crt.Nxt[i].Name) {
				q = append(q, *crt.Nxt[i])
			}
		}
	}
	
}


func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		file, _ := os.Open(args[0])
		rooms := createRooms(file)
		fmt.Println("Rooms: ", rooms)
		l := getLinks(args[0])
		fmt.Println("Links: ", l)
		linkRooms(rooms, l)
		fmt.Println("Rooms: ", rooms)
		fmt.Println("start: ", rooms[0])
		fmt.Println("end: ", rooms[len(rooms)-1])
		bfs(rooms)
	}
}
