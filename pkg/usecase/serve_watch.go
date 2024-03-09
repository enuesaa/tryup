package usecase

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/watch"
)

func ServeWatch(repos repository.Repos, watchpath string, basepath string, port int) error {
	go func() {
		watchctl := watch.New(repos)
		defer watchctl.Close()

		watchctl.WatchPath = watchpath

		if err := watchctl.Watch(); err != nil {
			repos.Log.Fatal(err)
		}

		<-make(chan struct{})
	}()

	return Serve(repos, basepath, port)
}
