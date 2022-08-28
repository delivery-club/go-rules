package errCheckInIf

type foo string

func (foo) boo() error {
	return nil
}

func (foo) boo2() (string, error) {
	return "", nil
}

func boo() error {
	return nil
}

func boo2() (string, error) {
	return "", nil
}

func warning1() {
	var err2 error
	if err := boo(); err2 != nil { // want `errCheckInIf: returned error 'err' must be checked`
		print(err)
	}

	var err error
	if err = boo(); err2 != nil { // want `errCheckInIf: returned error 'err' must be checked`
	}

	print(err)
}

func warning2() {
	var (
		d    foo
		err2 error
	)

	if err := d.boo(); err2 != nil { // want `errCheckInIf: returned error 'err' must be checked`
		print(err)
	}

	if _, err := d.boo2(); err2 != nil { // want `errCheckInIf: returned error 'err' must be checked`
		print(err)
	}

	var err error

	if err = d.boo(); err2 != nil { // want `errCheckInIf: returned error 'err' must be checked`
	}

	if _, err = d.boo2(); err2 != nil { // want `errCheckInIf: returned error 'err' must be checked`
	}

	print(err)
}

func warning3() {
	var (
		err2 error
	)
	if _, err := boo2(); err2 != nil { // want `errCheckInIf: returned error 'err' must be checked`
		print(err)
	}
	if _, err := boo2(); err2 != nil { // want `errCheckInIf: returned error 'err' must be checked`
		print(err)
	}
}
