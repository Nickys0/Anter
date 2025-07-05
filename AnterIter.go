package Anter

type AnterIter struct {
	an 		   *Anter
	cur_arg		int
}

// Creates a new simple Argument Iterator
func NewIter(an *Anter) AnterIter {
    return AnterIter {
        an: 		an, 
		cur_arg: 	-1,
    }
}

// It checks if there are any other arguments
func (ai *AnterIter) peakNext( ) bool {
    return ai.an.arg_count > (ai.cur_arg + 1)
}

// Goes to the next argument if exist
// Returns false if there is no other
func (ai *AnterIter) Next() bool {
    if ai.peakNext() {
        ai.cur_arg += 1
        return  true
    }
    return false
}

// It gets the current argument
// If there are no more arguments 
// it returns an EOA Arg
func (ai *AnterIter) Get( ) Arg {
    if ai.cur_arg < 0 {
		return argEOA(-1)
    }
    
	// We don't need to check the ai.cur_arg 
	// becouse it won't go further than the 
	// actual array size
	return ai.an.args[ai.cur_arg]
}

