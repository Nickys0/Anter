package Anter

// LFlag.tp
const(
	FTYPE_UNKNOWN = iota	/* It is used to indicate an error */
	FTYPE_BOOL				/* It means that the flag didn't expect any value */
	FTYPE_VALUE				/* It specifies that the flag expect a value */
)

// LFlag.flag
const(
	F_SINGLE_DASHED = 1 << iota //TODO: Add support for shorthand flags
	F_DOUBLE_DASHED	
	F_CONCAT

	F_DEFAULT_FLAG = 3
)

type LFlag struct {
	str 	string	/* flag string */
	tp 		int		/* Indicates the flag value type: FTYPE_... */
	flag 	int		/* It is used to indicate some info about the flag: F_...*/
}

// -name whould be provided without any dashes
// -tp indicates the flag type:
//		[-] FTYPE_VALUE	-> Expects a value after the flag
//		[-] FTYPE_BOOL		-> It is used to activate a flag
// - flag indicates some other flag for the flag	[You can use the F_DEFAULT_FLAG to activate default flags]
func NewFlag(name string, tp, flag int) LFlag{
	return LFlag{ str: name, tp: tp, flag: flag }
}

// like NewFlag but flag = F_DEFAULT_FLAG
func NewDefFlag(name string, tp int) LFlag{
	return LFlag{ str: name, tp: tp, flag: F_DEFAULT_FLAG }
}
