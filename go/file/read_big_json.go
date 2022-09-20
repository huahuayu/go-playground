package file

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type FlashbotBlock struct {
	BlockNumber       int    `json:"block_number"`
	MinerReward       string `json:"miner_reward"`
	Miner             string `json:"miner"`
	CoinbaseTransfers string `json:"coinbase_transfers"`
	GasUsed           int    `json:"gas_used"`
	GasPrice          string `json:"gas_price"`
	Transactions      []struct {
		TransactionHash  string `json:"transaction_hash"`
		TxIndex          int    `json:"tx_index"`
		BundleType       string `json:"bundle_type"`
		BundleIndex      int    `json:"bundle_index"`
		BlockNumber      int    `json:"block_number"`
		EoaAddress       string `json:"eoa_address"`
		ToAddress        string `json:"to_address"`
		GasUsed          int    `json:"gas_used"`
		GasPrice         string `json:"gas_price"`
		CoinbaseTransfer string `json:"coinbase_transfer"`
		TotalMinerReward string `json:"total_miner_reward"`
	} `json:"transactions"`
}

func (fb *FlashbotBlock) Unmarshal(b []byte) error {
	return json.Unmarshal(b, fb)
}

func unmarsha() {
	start := time.Now()

	fileName := "/Users/shiming/flashbot_blocks.json"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", fileName, err.Error())
	}

	fi, err := f.Stat()
	if err != nil {
		log.Fatalf("Could not obtain stat, handle error: %v", err.Error())
	}

	r := bufio.NewReader(f)
	d := json.NewDecoder(r)

	i := 0

	d.Token()
	for d.More() {
		fb := &FlashbotBlock{}
		d.Decode(fb)
		if i == 0 {
			fmt.Printf("%v \n", fb)
		}
		i++
	}
	//d.Token()
	elapsed := time.Since(start)

	fmt.Printf("Total of [%v] object created.\n", i)
	fmt.Printf("The [%s] is %s long\n", fileName, fileSize(fi.Size()))
	fmt.Printf("To parse the file took [%v]\n", elapsed)
}
