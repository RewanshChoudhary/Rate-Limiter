package readingdata

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/RewanshChoudhary/System-Monitor/configuration"
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
		if slices.Contains(STAT_KEYS, key) {
			stats[key] = value

		}

		fmt.Println("Read the stats")
		if err != nil {

		}

	}
	return stats, nil
}

func SendStatsToProducer(stats map[string]int64) (map[string]int64, error) {
	producerUrl := configuration.AppConfig.Prod

	requestBodyBytes, err := json.Marshal(stats)
	if err != nil {
		panic(fmt.Errorf("Json was not read %w", err))

	}
	requestBody := bytes.NewReader(requestBodyBytes)

	request, err := http.NewRequest("POST", producerUrl, requestBody)
	if err != nil {
		panic(fmt.Errorf("error in request creation %w", err))
	}
	request.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(fmt.Errorf("While sending the request %w", err))

	}
	if resp.StatusCode != 200 {
		panic(fmt.Errorf("The request was not sent properly"))

	}

	return nil, nil
}
