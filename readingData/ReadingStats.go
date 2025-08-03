package readingdata

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"honnef.co/go/tools/config"
)

var STAT_KEYS = []string{"MemTotal", "MemFree", "MemAvailable", "SwapTotal", "SwapCache", "SwapFree"}

func parseMemInfoLine(line string) (string, int64, error) {

	parts := strings.SplitN(line, ":", 2)

	typeInfo := strings.TrimSpace(parts[0])
	valueParts := strings.Fields(parts[1])
	value, err := strconv.ParseInt(valueParts[0], 10, 64)

	if err != nil {
		fmt.Errorf("Error during conversion to int64")
		panic(err)

	}

	return typeInfo, value, nil

}

func ReadMemoryStatus() (map[string]int64, error) {
	file, err := os.Open("/proc/meminfo")

	if err != nil {
		fmt.Errorf("Error In Accessing OS data")
		panic(err)

	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	stats := make(map[string]int64)

	for scanner.Scan() {
        
		key, value, err := parseMemInfoLine(scanner.Text())
		if slices.Contains(STAT_KEYS,key){
			stats[key] = value


			
		}
		
		fmt.Println("Read the stats")
		if err != nil {

		}

		

		

	}
	return stats, nil
}


 func sendStatsToProducer(stats map[string ]int64) (map[string]int64,error){
 	producerUrl:=config.Config.Prod
	
	

 }