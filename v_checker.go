package checker

import (
	"fmt"
	"log"
	"time"
	"runtime"
	"github.com/abdfnx/shell"
	"github.com/briandowns/spinner"
	commands "github.com/scmn-dev/secman/tools/constants"
)

func Checker() {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " 🔍 Checking for updates..."
	s.Start()

	err, out, errout := shell.ShellOut("")
	
	if runtime.GOOS == "windows" {
		err, out, errout = shell.PWSLOut(commands.Check_w())
	} else {
		err, out, errout = shell.ShellOut(commands.Check_ml())
	}
		
	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}
		
	s.Stop()
	fmt.Print(out)
}
