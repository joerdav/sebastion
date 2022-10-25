package webui

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"sync"

	"github.com/nxadm/tail"
)

type LogStore interface {
	CreateLogger(string, string) (*log.Logger, func(), error)
	TailLog(string, string) (io.Reader, func(), error)
	GetAllLogs(string, string) (string, error)
}

type Store struct {
	logFolder       string
	runningJobs     map[string]chan bool
	runningJobsLock sync.RWMutex
}

var _ LogStore = (*Store)(nil)

func NewLogStore() (LogStore, error) {
	hd, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("no home dir: %w", err)
	}
	folder, err := url.JoinPath(hd, ".sebastion")
	if err != nil {
		return nil, fmt.Errorf("generating log path: %w", err)
	}
	return &Store{
		logFolder:   folder,
		runningJobs: make(map[string]chan bool),
	}, nil
}

func (s *Store) CreateLogger(taskName, id string) (*log.Logger, func(), error) {
	s.runningJobsLock.Lock()
	defer s.runningJobsLock.Unlock()
	fp, err := url.JoinPath(s.logFolder, taskName, url.PathEscape(id)+".log")
	if err != nil {
		return nil, nil, fmt.Errorf("construct path: %w", err)
	}
	dir := filepath.Dir(fp)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, nil, fmt.Errorf("create log directory %s: %w", dir, err)
	}
	f, err := os.OpenFile(fp, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return nil, nil, fmt.Errorf("open file: %w", err)
	}
	w := io.MultiWriter(f, os.Stdout)
	l := log.New(w, "", log.Flags())
	s.runningJobs[id] = make(chan bool)
	return l, func() {
		f.Close()
		s.runningJobsLock.Lock()
		defer s.runningJobsLock.Unlock()
		close(s.runningJobs[id])
		delete(s.runningJobs, id)
	}, nil
}

type readerFunc func(b []byte) (n int, err error)

func (f readerFunc) Read(b []byte) (n int, err error) {
	return f(b)
}

func (s *Store) GetAllLogs(taskName, id string) (string, error) {
	fp, err := url.JoinPath(s.logFolder, taskName, url.PathEscape(id)+".log")
	if err != nil {
		return "", fmt.Errorf("construct path: %w", err)
	}
	f, err := os.Open(fp)
	if err != nil {
		return "", fmt.Errorf("open file: %w", err)
	}
	fs, err := io.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("read file: %w", err)
	}
	return string(fs), nil
}

func (s *Store) TailLog(taskName, id string) (io.Reader, func(), error) {
	fp, err := url.JoinPath(s.logFolder, taskName, url.PathEscape(id)+".log")
	if err != nil {
		return nil, nil, fmt.Errorf("construct path: %w", err)
	}
	dir := filepath.Dir(fp)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, nil, fmt.Errorf("create log directory %s: %w", dir, err)
	}
	t, err := tail.TailFile(fp, tail.Config{Follow: true})
	if err != nil {
		return nil, nil, fmt.Errorf("open log file: %w", err)
	}
	jobChan := s.runningJobs[id]
	return readerFunc(func(b []byte) (n int, err error) {
			select {
			case l, ok := <-t.Lines:
				if !ok {
					return 0, io.EOF
				}
				n = copy(b, []byte(l.Text+"\n"))
			case <-jobChan:
				return 0, io.EOF
			}
			return
		}),
		func() {
			_ = t.Stop()
		}, nil
}
