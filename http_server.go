package main

import (
	"errors"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/mux"
)

// GalleryServer
type GalleryServer struct {
	dispatcher *mux.Router
	watcher    *fsnotify.Watcher
	dirs       []string
}

func NewGalleryServer() (*GalleryServer, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	return &GalleryServer{
		dispatcher: mux.NewRouter(),
		watcher:    watcher,
		dirs:       make([]string, 0, 4),
	}, nil
}

// ServeHTTP implements an http.Handler that answers HTTP requests
func (s *GalleryServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.dispatcher.ServeHTTP(w, r)
}

// TODO(arkbriar@gmail.com) Provide a function to register JS file, like hugo.
// Or we can just use hugo as our engine.

func (s *GalleryServer) WatchGalleryDir(dir string) error {
	s.dirs = append(s.dirs, dir)
	return s.watcher.Add(dir)
}

func (s *GalleryServer) handleFSEvents() {
	defer s.watcher.Close()
	for {
		select {
		case event := <-s.watcher.Events:
			logrus.Debugln("fs event: ", event)
			if err := s.scanDirs(); err != nil {
				logrus.Errorln(err)
			}
		case err := <-s.watcher.Errors:
			logrus.Errorln("error: ", err)
		}
	}
}

func (s *GalleryServer) scanDirs() error {
	if len(s.dirs) == 0 {
		return errors.New("no dirs watched")
	}
	// TODO(arkbriar@gmail.com) Register js file and render the photo list.
}

func (s *GalleryServer) Start() error {
	if err := s.scanDirs(); err != nil {
		return err
	}
	done := make(chan bool)
	go s.handleFSEvents()
	<-done
}
