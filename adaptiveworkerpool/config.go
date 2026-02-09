package adaptiveworkerpool

type PoolConfig struct {
	MaxWorkers int
	MinWorkers int
}

type ConfigOptions func(*PoolConfig)
