# Description

### This is a video sorter I made for sorting video clips from length, I use this for valorant tracker highlights to sort the clips from most content to least
### Arguable pretty useless, just thought I'd add it to my github! :D


### This exports into an output markdown file so you should probably use a markdown reader such as:
#### - Retext - https://github.com/retext-project/retext ( Linux / Windows )

# Requirements for building

#### - Go - https://go.dev/doc/install
#### - Git - https://git-scm.com/downloads

# Building

### The release files are here but you can build the files yourself with:

## Linux / Windows

### Cloning repo (Can download zip, but this is quicker)
```git clone https://github.com/Toakley683/VideoSorter.git```

### Building to linux on linux with go
```go build video_sort.go```

### Building to windows on linux with go
```GOOS=windows GOARCH=amd64 go build -o video_sort.exe video_sort.go```

## MacOS (Unsupported)
