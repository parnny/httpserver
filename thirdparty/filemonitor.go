package thirdparty

import (
	"github.com/howeyc/fsnotify"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	utils "github.com/parnny/utils4go"
	log "github.com/cihub/seelog"
)

type monitor struct {
	watch 		*fsnotify.Watcher
	Pathmap		*utils.SafeMap
}

func NewFileMonitor() (monitor, error) {
	Mon, err := fsnotify.NewWatcher()
	return monitor{
		Mon,
		utils.NewSafeMap()}, err
}

func (self monitor)Watch(path string)  {
	self.watch.Watch(path)
	self.Pathmap.Set(path,true)

	childs, _ := ioutil.ReadDir(path)
	for _, child := range childs {
		if child.IsDir() {
			self.Watch(filepath.Join(path,child.Name()))
		}
	}
}

func (self monitor)RemoveWatch(path string)  {
	self.Pathmap.Foreach(func(fullpath string, v interface{}) {
		if strings.Contains(fullpath, path) {
			delete(self.Pathmap.Data, fullpath)
			self.watch.RemoveWatch(fullpath)
		}
	})
}

func (self monitor) Start() {
	go func() {
		for {
			select {
			case w := <-self.watch.Event:
				if w.IsModify() {
					continue
				}
				if w.IsDelete() {
					self.RemoveWatch(w.Name)
					continue
				}
				if w.IsRename() {
					self.RemoveWatch(w.Name)
					continue
				}
				if w.IsCreate() {
					root, _ := os.Stat(w.Name)
					if root.IsDir() {
						self.Watch(w.Name)
					}
					continue
				}
			case err := <-self.watch.Error:
				log.Error(err)
			}
		}
	}()
}

