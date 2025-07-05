package Anter

// Arg type:
// "--flag" 						-> flag
// "new"							-> Command
// "--flag=value" | "--flag value" 	-> value
const(
	ARGTP_EOA = iota
	ARGTP_COMMAND
	ARGTP_FLAG
	ARGTP_VALUE
	ARGTP_BINPATH

	ARGTP_UNKNOWN = -1
	ARGF_NONE = -1
)

func ArgtpToString(tp int) string{
	switch tp{
		case ARGTP_EOA: 		return "EOA"	
		case ARGTP_COMMAND: 	return "COMMAND"	
		case ARGTP_FLAG: 		return "FLAG"	
		case ARGTP_VALUE: 		return "VALUE"	
		case ARGTP_UNKNOWN:		return "UNKNOWN"
		case ARGTP_BINPATH:		return "BINPATH"
	default:
			return "UNKNOWN"
	}
}

type Arg struct{

	tp 	int			/* Arg type */
	str string 		/* arg str  */

	r_indx int		/* [relative index] based on the arg type */
	a_indx int		/* [argument index] Indicates the arg pos */
	
	// TODO: we should try remove this
	argf int		/* Arg flag: it is used to indicate some addtional info */
}
 
func (arg *Arg) Str( ) string {
	return arg.str
}

func (arg *Arg) Type( ) int {
	return arg.tp
}

// Warning the returned idx doesn't always
// mean that the specified argument is in os.Arg[a_indx]
func (arg *Arg) AIdx( ) int {
	return arg.a_indx
}

// If the arg is a flag or a command
// this return the index of the flag | command
// in the array that was provided with the function 'InitLib'
func (arg *Arg) RIdx( ) int {
	return arg.r_indx
}

// An invalid type is any argument that has his
// type == ARGTP_UNKNOWN
func (arg *Arg) IsValidType( ) bool {
	return arg.tp != ARGTP_UNKNOWN
}

func (arg *Arg) IsCom( ) bool {
	return arg.tp == ARGTP_COMMAND
}

func (arg *Arg) IsFlag( ) bool {
	return arg.tp == ARGTP_FLAG
}

func (arg *Arg) IsValue( ) bool {
	return arg.tp == ARGTP_VALUE
}

func (arg *Arg) IsBinPath( ) bool {
	return arg.tp == ARGTP_BINPATH
}

func (arg *Arg) IsEOA() bool {
	return arg.tp == ARGTP_EOA
}




/////////////////////////////////////////////////////////////
//                       INTERNALS                        //
///////////////////////////////////////////////////////////

func argFlag_Bool(_str string, r_idx, a_idx int) Arg{
	return Arg{ tp: ARGTP_FLAG, str: _str, 		// basics
				r_indx: r_idx, a_indx: a_idx, // Idicies
				argf: FTYPE_BOOL }			// Flags
}

func argFlag_Value(_str string, r_idx, a_idx int) Arg{
	return Arg{ tp: ARGTP_FLAG, str: _str, 
				r_indx: r_idx, a_indx: a_idx, 
				argf: FTYPE_VALUE}
}

func argEOA(idx int) Arg{
	return Arg{ tp: ARGTP_EOA, str: "", 
				r_indx: -1, a_indx: idx, 
				argf:  ARGF_NONE}
}
 
func argBasic( _tp int, _str string, r_idx, a_idx int) Arg{
	return Arg{ tp: _tp, str: _str, 		// basics
				r_indx: r_idx, a_indx: a_idx, // Idicies
				argf: ARGF_NONE }			// Flags
}
