package checker

import (
	"fmt"
	"log"
	"time"
	"runtime"
	"github.com/secman-team/shell"
	"github.com/briandowns/spinner"
)


func Checker() {
	check_w := `
		$releases = "https://api.github.com/repos/secman-team/secman/releases"
	
		$l = (Invoke-WebRequest -Uri $releases -UseBasicParsing | ConvertFrom-Json)[0].tag_name
	
		$c = secman verx
	
		if ($l -ne $c) {
			$nr = "there's a new release of secman is avalaible: "
			$up = "to upgrade run "
			$smu = "sm-upg start"
	
			Write-Host ""
			Write-Host -NoNewline $nr -ForegroundColor DarkYellow
			Write-Host "$c -> $l" -ForegroundColor DarkCyan
			Write-Host -NoNewline $up -ForegroundColor DarkYellow
			Write-Host $smu -ForegroundColor DarkCyan
		}
	`
	
	check_ml := `
		l=$(curl --silent "https://api.github.com/repos/secman-team/secman/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
		c=$(secman verx | tr -d '\n')
	
		if [ $l != $c ]; then
			nr="there's a new release of secman is avalaible:"
			up="to upgrade run"
			smu="secman upgrade"
	
			echo ""
			echo "$nr $c -> $l"
			echo "$up $smu"
		fi
	`

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " 🔍 Checking for updates..."
	s.Start()

	err, out, errout := shell.ShellOut("")
	
	if runtime.GOOS == "windows" {
		err, out, errout = shell.PWSLOut(check_w)
	} else {
		err, out, errout = shell.ShellOut(check_ml)
	}
		
	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}
		
	s.Stop()
	fmt.Print(out)
}
