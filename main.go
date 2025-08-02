package systemmonitor

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


var STAT_KEYS=[]string{"MemTotal","MemFree","MemAvailable","SwapTotal","SwapCache","SwapFree"}

func parseMemInfoLine(line string) (string, int, error) {
	fields := strings.Fields(line)
	if len(fields) < 2 {
		return "", 0, fmt.Errorf("invalid line: %s", line)
	}

	value, err := strconv.Atoi(fields[1])
	if err != nil {
		return "", 0, err
	}

	return strings.TrimSuffix(fields[0], ":"), value, nil
}

func ReadMemoryStatus() (map[string]int,error) {
	file,err:=os.Open("/proc/meminfo")

	if (err !=nil){
		fmt.Errorf("Error In Accessing OS data")

		

		
	}
	defer file.Close()
	scanner:=bufio.NewScanner(file)
	stats:=make(map[string ]int )



   for scanner.Scan(){
	key,value:=parseMemInfoLine(scanner.Text())


   }
}
