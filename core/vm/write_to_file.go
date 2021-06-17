package vm

import (
  "os"
  "log"
	"strconv"
  "encoding/csv"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
  "strings"
)

// Store past lines to filter for duplicates
var container = make(map[uint64]string)

// file name can be made dependent on txType or similar in writeTransaction
var fileName = "D:/Glitznerf/Documents/uzh_bc/etc_parsed/internal_tx.csv"

// Write transacion to file
func writeTransaction(blockNo *big.Int, txType string, from common.Address, to common.Address, value *big.Int, gas uint64, gasPrice *big.Int) {
  blockNoString := blockNo.String()
  blockNoUint := blockNo.Uint64()
  senderString := from.Hex()
	receiverString := to.Hex()
	valueString := value.String()
	gasString := strconv.FormatUint(gas, 10)
	gasPriceString := gasPrice.String()

	line := []string{blockNoString, txType, senderString, receiverString, valueString, gasString, gasPriceString}
  lineString := strings.Join(line[1:],"")

  if containerEntry, ok := container[blockNoUint]; ok {
    if !(strings.Contains(containerEntry, lineString)) {
  	   writeFile(line, fileName)
       container[blockNoUint] = containerEntry + lineString
    }
  } else {
    writeFile(line, fileName)
    container[blockNoUint] = lineString
  }

  // Clear cache
  if container[blockNoUint-uint64(2)] != "" {
    container[blockNoUint-uint64(2)] = ""
  }
  if container[blockNoUint-uint64(3)] != "" {
    container[blockNoUint-uint64(3)] = ""
  }
}

// String-Array to CSV Writer
func writeFile(line []string, fileName string) {

	// Open or Create respective file
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
  defer file.Close()
	if err != nil {
    log.Fatal(err)
  }

	// Start writer and write to file
	writer := csv.NewWriter(file)
  defer writer.Flush()

	writeErr := writer.Write(line)
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}
