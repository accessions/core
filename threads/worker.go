package threads

type Worker struct {
	job     func()
	workers int
}

func NewWorker(job func(), workers int) Worker {
	return Worker{job: job, workers: workers}
}

func (w Worker) GRun() {

	g := NewGroup()
	for i := 0; i < w.workers; i++ {
		g.SafeRun(w.job)
	}
	g.Wait()

}
