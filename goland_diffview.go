package diffview

import (
	"os"
	"os/exec"
	"strings"
)

type golandDiffViewer struct {
}

func (g golandDiffViewer) open(a, b string) error {

	golandExecutablePath,err := g.golandExecutablePath()
	if err != nil {
		return err
	}

	return g.openDiffView(golandExecutablePath,a,b)
}

func (g golandDiffViewer) openDiffView (golandExecutablePath,a,b string) error {
	cmd := exec.Command(golandExecutablePath, "diff", a, b)
	cmd.Env = os.Environ()
	return cmd.Run()
}

/*
golandExecutablePath finds the goland executable
the Goland IDE must be running before this command is executed

ps -ef | grep goland | awk 'NR==1{for(i=8;i<=NF;i++)print $i}'
get current processes including goland, first line, print column 8 until end

awk 'NR%2{printf "%s ",$0;next;}1'
concatenate lines
*/
func (g golandDiffViewer) golandExecutablePath () (string ,error){
	getPath := exec.Command("sh","-c",`ps -ef | grep goland | awk 'NR==1{for(i=8;i<=NF;i++)print $i}' | awk 'NR%2{printf "%s ",$0;next;}1'`)

	path,err := getPath.Output()
	if err != nil {
		return "",err
	}

	return strings.Replace(string(path),"\n","",-1),err
}
