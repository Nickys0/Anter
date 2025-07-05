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

	/* Mandatory */
	tp 	int			/* Arg type */
	str string 		/* arg str  */

	r_indx int		/* [relative index] base on the arg type */
	a_indx int		/* Index of the arg that was given from the CLI */
	argf int		/* Arg flag: it is used to indicate some addtional info */
}

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
