package Anter

import (
	"fmt"
	"log"
	"os"
	"path"
	"testing"
)

var _def_coms = []string{"test1", "test2", "test3", "test4"}
var _def_flags = []LFlag{
	{ str: "bool", tp: FTYPE_BOOL,  flag: F_DEFAULT_FLAG  },
	{ str: "val", tp: FTYPE_VALUE, flag: F_DEFAULT_FLAG  },
}


/////////////////////////////////////////////////////////////
//                        TESTERS                         //
///////////////////////////////////////////////////////////

func DefInit(){
	os.Args = []string{"binpath", "test1", "--val", "value"}
}

/// @Test [1] Command + flag and value
func TestArg01(t *testing.T){
	defer argTestReset( )

	DefInit( )

	/* This function should not fail */
	if _, err := argTestInit(_def_coms, _def_flags); err != nil {
		t.Fatalf("error: %s", err.Error())
	}
}

/// @Test [2] Testing an invalid flag error
func TestLFlag01(t *testing.T){
	defer argTestReset( )

	DefInit()

	var flags []LFlag 
	copy(flags, _def_flags)

	flags = append(flags, LFlag{ str: "-bool", tp: FTYPE_VALUE, flag: F_DEFAULT_FLAG })
	
	if _, err := argTestInit(_def_coms, flags); err != nil{
		log.Printf("error: %s", err.Error())
	}else{
		t.Fatal("This function should fail")
	}
}

// Getting the value of flag 
// Expected: Everything should go well
// argv: ["binary", "test1, "--bool" "--val", "value"]
//			 0        1         2        3		 4
func TestGetFlags01(t *testing.T){
	defer argTestReset( )

	os.Args = []string{"binary", "test", "--bool", "--val", "value"}

	parser, err := argTestInit(_def_coms, _def_flags)

	if err != nil {
		t.Fatal(err.Error())
	}

	// We excepect to get the flag value
	if name, er := parser.GetFlagString("--val"); er == nil {
		assert(name != "", "Empty string error!")
		fmt.Printf("The provided value was: %q\n", name)
	}else{
		t.Fatal(er.Error())
	}

	// We excepect get true as the return value
	if boolean, er := parser.GetFlagBool("--bool"); er == nil {
		assert(boolean == true, "The flag was not given or an error was occured")
		log.Printf("The flag was succesfully provied! %q", "bool")
	} else{
		t.Fatalf("%s", er.Error())
	}

	printArg(parser)
}

// Getting the value of flag 
// Expected: GetFlagString should fail becouse 
// 			the flag "name" doesn't exist
func TestGetFlags02(t *testing.T){
	defer argTestReset( )
	
	DefInit()

	parser, err := argTestInit(_def_coms, _def_flags)

	if err != nil{
		t.Fatal(err)
	}

	if _, er := parser.GetFlagString("name"); er == nil {
		t.Fatal("This should fail becouse the flag 'name' doesn't exist!")
	}
		
	printArg(parser)
}

func TestArgIterator(t *testing.T){
	defer argTestReset( )
	DefInit( )

	parser, err := argTestInit(_def_coms, _def_flags)
	if err != nil {
		t.Fatalf("error: %s", err.Error())	
	}

	it := NewIter(&parser)
	for it.Next( ) {
		arg := it.Get( )
		if arg.IsEOA() {
			log.Printf("Arg[%02d]: %s\n", arg.AIdx(), "<EOA>")
		}else{
			log.Printf("Arg[%02d]: %s\n", arg.AIdx(), arg.Str())
		}
	}
}

func TestArgLast(t *testing.T){
	defer argTestReset()
	DefInit()

	parser, err := argTestInit(_def_coms, _def_flags)
	if err != nil {
		t.Fatalf("%s: %s", RedTxt("error"), err.Error())
	}

	last := parser.args[parser.arg_count] 
	if last.IsEOA( ) {
		log.Printf("Last Argument[%d]: <EOA>", last.a_indx)
	}else{
		t.Fatalf("error: this argument should be the last but is type: %s", ArgtpToString(last.tp))
	}

}
/////////////////////////////////////////////////////////////
//                                                        //
///////////////////////////////////////////////////////////


/////////////////////////////////////////////////////////////
//                       INTERNALS                        //
///////////////////////////////////////////////////////////
// Basic fucntion that initialize and analize the arguments
func argTestInit(coms []string, flags []LFlag) (Anter, error) {
	if err := InitLib(coms, flags); err != nil {
		return Anter{}, err
	}
	
	An, err := AnalArg( )
	if err != nil {
		return Anter{}, err
	}

	return An, nil
}

func argTestReset() {
	os.Args = 	nil
	argv = 		nil
	argc = 		0
}

func printArg(an Anter){
	for idx, arg := range an.args{
		switch arg.tp {
		case ARGTP_EOA:
			fmt.Printf("Arg[%02d]{str: %10s, type: %10s}\n", idx, "<EOA>", ArgtpToString(arg.tp))
		case ARGTP_BINPATH:
			fmt.Printf("\nArg[%02d]{str: %10s, type: %10s}\n", idx, path.Base(arg.str), ArgtpToString(arg.tp))
		case ARGTP_COMMAND, ARGTP_FLAG, ARGTP_UNKNOWN, ARGTP_VALUE:
			fmt.Printf("Arg[%02d]{str: %10s, type: %10s}\n", idx, arg.str, ArgtpToString(arg.tp))
		default:
			panic("invalid type")
		}
	}

	fmt.Println("")
}
