package erratum

// Use opens a resource, calls Frob(input) on the result resource and then
// closes that resource (in all cases).
func Use(o ResourceOpener, input string) (result error) {
	// Attemp to open the resource and return if an error was found
	resource, result := open(o)
	if result != nil {
		return
	}
	defer resource.Close()

	// Handle panic condition and recover from it
	defer func() {
		// check if panic is in progress
		if recover := recover(); recover != nil {
			// if a FrobError, calls resource's defrob function
			if asFrob, ok := recover.(FrobError); ok {
				resource.Defrob(asFrob.defrobTag)
			}

			// if any error was thrown, return it
			if asError, ok := recover.(error); ok {
				result = asError
			}
		}
	}()

	resource.Frob(input)

	return
}

func open(o ResourceOpener) (resource Resource, err error) {
	for {
		resource, err = o()

		if _, ok := err.(TransientError); ok {
			// If a transient error, try again
			continue
		} else {
			// terminate if any other error occurred or
			// the resource was opened successfully
			break
		}
	}

	return
}
