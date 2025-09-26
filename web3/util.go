package main

func mustNew(r interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return r
}
