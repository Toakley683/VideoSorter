package main

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/alfg/mp4"
)

type Video struct {
	filePath string
	fileName string
	duration float64
}

func videoToString(index int, video Video) string {
	var str = "#### ○	"

	//str = str + "F: " + video.filePath
	str = str + "[" + video.fileName + "](" + video.filePath + ")"

	if video.duration != -1 {
		str = str + " - [**" + strconv.FormatFloat(video.duration, 'f', -1, 64) + "**]"
	}

	wd, _ := os.Getwd()
	fileDirectory := strings.Split((wd + "/" + video.filePath), "/")
	fileDirectory = fileDirectory[:len(fileDirectory)-1]

	fileDirectoryString := strings.Join(fileDirectory, "/")

	str = str + "  - [**Directory**](" + fileDirectoryString + "/)"

	return str

}

func main() {

	var count int32 = 0

	files := []Video{}
	brokenFiles := []Video{}

	filepath.WalkDir("./", func(path string, info os.DirEntry, err error) error {

		if err != nil || info.IsDir() {
			return err
		}

		var split = strings.Split(info.Name(), ".")
		var extension = split[len(split)-1]

		if extension != "mp4" {
			return nil
		}

		duration, err := Mp4Duration(path)

		if duration == -1 {

			// File has invalid data!
			brokenFiles = append(brokenFiles, Video{
				filePath: path,
				fileName: info.Name(),
				duration: -1,
			})
			return nil

		}

		if err != nil {
			return err
		}

		count++

		var Video = Video{
			filePath: path,
			fileName: info.Name(),
			duration: duration,
		}

		files = append(files, Video)

		return nil
	})

	var sorted = mergeSort(files)

	var textData = ""
	textData = textData + "# List of videos that have unreadable data: [" + strconv.Itoa(len(brokenFiles)) + "] \n\n"

	for index := 0; index < len(brokenFiles); index++ {

		var video = brokenFiles[index]
		textData = textData + videoToString(index, video) + " <br>"

	}

	if len(brokenFiles) != 0 {
		textData = textData + "\n"
	}

	textData = textData + "# Sorted list of videos in order of duration: [" + strconv.Itoa(len(files)) + "] \n\n"

	for index := 0; index < len(sorted); index++ {

		var video = sorted[index]
		textData = textData + videoToString(index, video) + "\n"

	}

	err := os.WriteFile("./output.md", []byte(textData), 0644)

	if err != nil {
		panic(err)
	}

}

func Mp4Duration(path string) (float64, error) {

	file, err := os.Open(path)

	if err != nil {
		return 0, err
	}
	defer file.Close()

	info, err := file.Stat()

	if err != nil {
		return 0, err
	}

	mp4, err := mp4.OpenFromReader(file, info.Size())

	if err != nil {
		return 0, err
	}

	if mp4 == nil ||
		mp4.Moov == nil ||
		mp4.Moov.Mvhd == nil {

		// Invalid data error!

		return -1, err

	}
	file.Close()

	var length = float64(mp4.Moov.Mvhd.Duration) / float64(mp4.Moov.Mvhd.Timescale)

	return length, err

}

func mergeSort(items []Video) []Video {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []Video, b []Video) []Video {

	final := []Video{}

	i := 0

	j := 0

	for i < len(a) && j < len(b) {

		// > Reverse sorted
		// < Normal sorted

		if a[i].duration > b[j].duration {

			final = append(final, a[i])

			i++

		} else {

			final = append(final, b[j])

			j++

		}

	}

	for ; i < len(a); i++ {

		final = append(final, a[i])

	}

	for ; j < len(b); j++ {

		final = append(final, b[j])

	}

	return final

}
