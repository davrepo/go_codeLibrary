package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

// given N sorted slices of K songs, merge them into a single sorted slice of songs

const path = "songs.json"

// Song stores all the song related information
type Song struct {
	Name      string `json:"name"`
	Album     string `json:"album"`
	PlayCount int64  `json:"play_count"`
}

// makePlaylist makes the merged sorted list of songs
func makePlaylist(albums [][]Song) []Song {
	// flatten the albums into a single slice
	var songs []Song
	for _, album := range albums {
		songs = append(songs, album...)
	}

	// sort the songs by PlayCount
	quickSort(songs, 0, len(songs)-1)

	return songs
}

// quickSort sorts the songs by PlayCount from high to low
func quickSort(songs []Song, low, high int) {
	if low < high {
		pivot := partition(songs, low, high)
		quickSort(songs, low, pivot-1)
		quickSort(songs, pivot+1, high)
	}
}

// partition partitions the songs slice into two parts
func partition(songs []Song, low, high int) int {
	pivot := songs[high].PlayCount
	i := low - 1
	for j := low; j < high; j++ {
		if songs[j].PlayCount < pivot {
			i++
			songs[i], songs[j] = songs[j], songs[i]
		}
	}
	songs[i+1], songs[high] = songs[high], songs[i+1]
	return i + 1
}

func main() {
	albums := importData()
	printTable(makePlaylist(albums))
}

// printTable prints merged playlist as a table
func printTable(songs []Song) {
	w := tabwriter.NewWriter(os.Stdout, 3, 3, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "####\tSong\tAlbum\tPlay count")
	for i, s := range songs {
		fmt.Fprintf(w, "[%d]:\t%s\t%s\t%d\n", i+1, s.Name, s.Album, s.PlayCount)
	}
	w.Flush()

}

// importData reads the input data from file and creates the friends map
func importData() [][]Song {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data [][]Song
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
