package Anter

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

var argv  []string
var argc 	int

var _commands []string
var _flags 	  []LFlag

var _initialize bool = false

type Anter struct {
	args 	[]Arg		/* Argument array */
	arg_count int		/* Argument count ::! Not always is equal to argc */

	command []*Arg 		/* Arg refrence to commands if present */
	flags 	[]*Arg		/* Arg reference to flags if present */
}

// ** Mandatory **
// This is the first function to be called to
//	Init the library
func InitLib(com []string, fls []LFlag) error{
	__temp := os.Args

	// We split the value declared as: "--flag=value"
	for _, a := range __temp{
		argv = append(argv, strings.Split(a, "=")...)
	}

	argc = len(argv)
	_commands = com

	// TODO #[1]: we shall validate the coms and the flags:
	// Checking that the flag only contains the name and not the dashes
	_flags = fls
	for _, f := range fls {
		if f.tp == FTYPE_UNKNOWN {
			flag := ""
			if f.str == "" {
				flag = "NULL"
			}
			return ErrF("Invalid flag passed to %s: <%s>", 
							GreenTxt("InitLib"), flag)
		}
		
		if strings.Contains(f.str, "-"){
			return ErrF("invalid flag passed in func %s: <%s>", 
							GreenTxt("InitLib"), RedTxt(f.str))
		}
	}

	_initialize = true
	return nil
}

// This function analise the argv from os.Argv
// and stores the info in the Anter structure
func AnalArg() (Anter, error){
	var out Anter

	if !_initialize {
		return out, ErrF("The Library wasn't initizalized! Call 'InitLib' first")
	}

	out.args = append(out.args, Arg{str: argv[0], tp: ARGTP_BINPATH, r_indx: -1})
	out.arg_count++

	for idx, a := range argv[1:]{

		// Is a command?
		if i := slices.Index(_commands, a); i >= 0 {
			out.args = append(out.args, Arg{ARGTP_COMMAND, a, i, (idx + 1), ARGF_NONE})
			
			/* Setting a reference to a command */
			out.command = append(out.command, &out.args[out.arg_count - 1]) 	

			// is a flag?
		}else if fIdx := itsFlag(a); fIdx >= 0 {

			ftype := _flags[fIdx]
			
			//? Should we check if the flag is a valuable flag and so
			//? we need to peak the next arg to know if exist?

			// Setting the argf flag of the Arg structure based on the flag type 
			switch ftype.tp {
			case FTYPE_BOOL:
				out.args = append(out.args, argFlag_Bool(a, fIdx, (idx + 1)))
			case FTYPE_VALUE:
				out.args = append(out.args, argFlag_Value(a, fIdx, (idx + 1)))
			default:
				panic("Unreachable: Invalid flag type")	/* This should be Unreachable */
			}
			
			// Adding a reference of the flag
			out.flags = append(out.flags, &out.args[out.arg_count])

		}else{

			last_arg := out.args[out.arg_count - 1]

			/* If the prec arg was a valueable flag the 
				current one is the value */
			if (last_arg.tp == ARGTP_FLAG) && (last_arg.argf == FTYPE_VALUE){
				out.args = append(out.args, argBasic(ARGTP_VALUE, a, -1, (idx + 1)))
			}else{
				out.args = append(out.args, argBasic(ARGTP_UNKNOWN, a, -1, (idx + 1)))
			}
		}

		out.arg_count++
	}

	out.args = append(out.args, argEOA(out.arg_count))

	return out, nil
}

// TODO: make this function public: 
// 		 func FlagIdx(...) int
// Returns < 0 if the flag was not found
func ifFlagExist_Idx(flag string) int {
	for idx, f := range _flags {
		if f.str == flag {
			return  idx
		}
	}

	return -1
}

// If present it return the flag indx from the arg list
func (an *Anter) isFlagPrs(SFlag string) int {
	for _, f := range an.flags{
		if SFlag == f.str{
			return f.a_indx
		}
	}
	return -1
}

// TODO: add support for single dashed flag:
//	"-h"	->	"h" | "help"
// This function remove the flag header ("--" | "-")
// from the provided string and returns the result
func UnwrapStrFlag(flag string) string{
	out := flag
	if strings.Contains(out, "-"){
		n := strings.Count(out, "-")
		out = flag[n:]
	}
	return out
}	

// For boolean flag it return TRUE | FALSE based on
//   if the flag was provied or not
// For valuable flag it returns the value if
//   given or an error
// If the flag was not provided during initialization it return an error
// If the flag was not given by the user it returns an error
func (an *Anter) GetFlagValue(SFlag string) (string, error){
	out := "" 
	
	if !_initialize {
		return out, ErrF("The Library wasn't initizalized! Call 'InitLib' first")
	}

	// 1) Flag exist?
	fIdx := ifFlagExist_Idx(UnwrapStrFlag(SFlag))
	if fIdx < 0 {
		return out, ErrF("The provided flag was not provided during initialization: %s", GreenTxt(SFlag))
	}

	// 2) was flag provided?
	idx := an.isFlagPrs(SFlag)

	// 2.5) Is a boolean flag ?
	if _flags[fIdx].tp == FTYPE_BOOL {
		if idx < 0 {
			return "FALSE", nil
		}
		return "TRUE", nil
	}

	if idx < 0 {
		return out, ErrF("The flag %s was not given from the user", GreenTxt(SFlag))
	}

	// 3) Flag value was given?
	if (idx + 1) >= an.arg_count {
		return out, ErrF("The flag %s expected value but found: <EOA>", GreenTxt(SFlag))
	}

	out = an.args[idx + 1].str

	// 4) Return the flag value
	return out, nil
}
 
// The flag can be provided like "--flag" | "flag" 
func (an *Anter) GetFlagInt(flag string, bitSize int) (int64, error){
	str_val, err := an.GetFlagValue(flag)
	if err != nil {
		return  0, err
	}
	
	return strconv.ParseInt(str_val, 10, bitSize)
}

// The flag can be provided like "--flag" | "flag" 
func (an *Anter) GetFlagUInt(flag string, base int) (uint64, error){
	str_val, err := an.GetFlagValue(flag)
	if err != nil {
		return  0, err
	}

	return strconv.ParseUint(str_val, 10, base)
}

// The flag can be provided like "--flag" | "flag" 
func (an *Anter) GetFlagFloat(flag string, base int) (float64, error){
	str_val, err := an.GetFlagValue(flag)
	if err != nil {
		return  0, err
	}

	return strconv.ParseFloat(str_val, base)
}

// The flag can be provided like "--flag" | "flag" 
func (an *Anter) GetFlagInt64(flag string) (int64, error){
	return an.GetFlagInt(flag, 64)
}

// The flag can be provided like "--flag" | "flag" 
func (an *Anter) GetFlagInt32(flag string) (int32, error){
	out, err := an.GetFlagInt(flag, 32)
	return int32(out), err
}

// The flag can be provided like "--flag" | "flag" 
func (an *Anter) GetFlagInt16(flag string) (int16, error){
	out, err := an.GetFlagInt(flag, 16)
	return int16(out), err	
}

// The flag can be provided like "--flag" | "flag" 
func (an *Anter) GetFlagInt8(flag string) (int8, error){
	out, err := an.GetFlagInt(flag, 8)
	return int8(out), err	
}

// The flag can be provided like "--flag" | "flag" 
func (an *Anter) GetFlagUInt64(flag string) (uint64, error){
	return an.GetFlagUInt(flag, 64)
}

// The flag can be provided like "--flag" | "flag" 
func (an *Anter) GetFlagUInt32(flag string) (uint32, error){
	out, err := an.GetFlagUInt(flag, 32)
	return uint32(out), err
}

// The flag can be provided like "--flag" | "flag" 
func (an *Anter) GetFlagUInt16(flag string) (uint16, error){
	out, err := an.GetFlagUInt(flag, 16)
	return uint16(out), err
}

// The flag can be provided like "--flag" | "flag"
func (an *Anter) GetFlagUInt8(flag string) (uint8, error){
	out, err := an.GetFlagUInt(flag, 8)
	return uint8(out), err
}

// The flag can be provided like "--flag" | "flag" 
func (an *Anter) GetFlagFloat32(flag string) (float32, error){
	out, err := an.GetFlagFloat(flag, 32)
	return float32(out), err
}

// The flag can be provided like "--flag" | "flag"
func (an *Anter) GetFlagFloat64(flag string) (float64, error){
	return an.GetFlagFloat(flag, 64)
}

// The 'flag' can be passed like "--flag" | "flag" 
// Return true if the flag was provided false if it wasn't 
// provided or due to an error
func (an *Anter) GetFlagBool(flag string) (bool, error){
	out, err := an.GetFlagValue(flag)
	if err == nil && out == "TRUE" {
		return  true, nil
	}

	return false, err
}

// This function gets the flag value if it was a valuable flag and
// returns it
// The flag can be provided like "--flag" | "flag"
func (an *Anter) GetFlagString(flag string) (string, error){
	return  an.GetFlagValue(flag)
}

// It return the first command that occured
// An EOA Arg is return if it wasn't present
func (an *Anter) GetCom( ) Arg {
	if len(an.command) > 0 {
		return *an.command[0]
	}

	return argEOA(-1)
}

// TODO? shall we call this function without the _Str suffix?
// It checks if the provided command (as a string) was given by the user
func (an *Anter) IsComPresent_Str(com string) bool {
	for _, cm := range an.command {
		if cm.str == com {
			return true
		}
	}

	return false
}


/////////////////////////////////////////////////////////////
//                       INTERNALS                        //
///////////////////////////////////////////////////////////

// [INTERNAL]
// Returns the idx of the flag if present. < 0 indicates that it wasn't found
func itsFlag(a string) int {
	
	if !_initialize {
		return -1
	}

	var __temp string

	for idx, f := range _flags {
		if (f.flag & F_DOUBLE_DASHED) == F_DOUBLE_DASHED {
			__temp = "--" + f.str
			if a == __temp {
				return idx
			}
		}
		
		if (f.flag & F_SINGLE_DASHED) == F_SINGLE_DASHED {
			__temp = "-" + f.str
			if a == __temp{
				return idx
			}
		}
	}

	return -1
}

// TODOS:
//	Add support for shorthand flags 
//	Move all the private function in the appropriate section
//  Add func: ArgCount( )   // Should return Anter.arg_count - 1 (becouse of binpath)
// 			  NoArg( )		// Should return true if no arguments (aside from the binpath) were priveded 
// 			  GetFlagValueIfGiven( ) // Should return the value only if given and if it doesn't it just return an empty string
//			  WasFlagGiven( )		// returns false if an error occured or becouse it wasn't given 
