// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package container

// Injectors from wire.go:

func initEsim() *Esim {
	config := provideConf()
	logger := provideLogger(config)
	prometheus := providePrometheus(config, logger)
	tracer := provideTracer(config, logger)
	esim := &Esim{
		prometheus: prometheus,
		Logger:     logger,
		Conf:       config,
		Tracer:     tracer,
	}
	return esim
}

func NewMockEsim() *Esim {
	config := provideMockConf()
	prometheus := provideMockProme(config)
	logger := provideLogger(config)
	tracer := provideNoopTracer()
	esim := &Esim{
		prometheus: prometheus,
		Logger:     logger,
		Conf:       config,
		Tracer:     tracer,
	}
	return esim
}
