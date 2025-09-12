package main

func mustNew(r interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return r
}

func mustCall(err error) {
	if err != nil {
		panic(err)
	}
}

func mustRoutine(fn func() error) {
	go func() {
		err := fn()
		if err != nil {
			panic(err)
		}
	}()
}
