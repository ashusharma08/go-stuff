package adaptiveworkerpool

// this should manage the current state. its job is to add/remove workers
func (w *workerPool) controller() {
	for {
		select {
		case <-w.ticker.C:
			//scale
			w.scale()
		case <-w.controllerShutdown:
			return
		case <-w.ctx.Done():
			return
		}
	}
}

func (w *workerPool) scale() {

}
