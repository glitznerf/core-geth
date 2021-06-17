package ethash

import (
  "os"
  "log"
	"strconv"
  "encoding/csv"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

// Filename is kept throughout application
var fileName = "D:/Glitznerf/Documents/uzh_bc/etc_parsed/internal_tx.csv"

// Write Transaction to file
func write__Transaction(blockNo *big.Int, txType string, from common.Address, to common.Address, value *big.Int, gas uint64, gasPrice *big.Int) {
  blockNoString := blockNo.String()
  senderString := from.Hex()
	receiverString := to.Hex()
	valueString := value.String()
	gasString := strconv.FormatUint(gas, 10)
	gasPriceString := gasPrice.String()

	line := []string{blockNoString, txType, senderString, receiverString, valueString, gasString, gasPriceString}

  // Ignore burn address reward
  if senderString != "0x0000000000000000000000000000000000000000" {
    writeToFile(line, fileName)
  }
}

// String-Array to CSV Writer
func writeToFile(line []string, fileName string) {

	// Open or Create respective file
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
  defer file.Close()

	if err != nil {
    log.Fatal(err)
  }

	// Start writer and write to file
	writer2 := csv.NewWriter(file)
  defer writer2.Flush()

	writeErr := writer2.Write(line)
	if writeErr != nil {
		log.Fatal(writeErr)
	}

}
