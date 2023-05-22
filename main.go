package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

const dnsCacheRefreshInterval = 5 * time.Minute

type Transaction struct {
	TransactionHash       string `json:"transaction_hash"`
	BundleType            string `json:"bundle_type"`
	TxIndex               int    `json:"tx_index"`
	BundleIndex           int    `json:"bundle_index"`
	BlockNumber           int    `json:"block_number"`
	EoaAddress            string `json:"eoa_address"`
	ToAddress             string `json:"to_address"`
	GasUsed               int    `json:"gas_used"`
	GasPrice              string `json:"gas_price"`
	EthSentToFeeRecipient string `json:"eth_sent_to_fee_recipient"`
	FeeRecipientEthDiff   string `json:"fee_recipient_eth_diff"`
}

type Block struct {
	BlockNumber           int           `json:"block_number"`
	FeeRecipientEthDiff   string        `json:"fee_recipient_eth_diff"`
	FeeRecipient          string        `json:"fee_recipient"`
	EthSentToFeeRecipient string        `json:"eth_sent_to_fee_recipient"`
	GasUsed               int           `json:"gas_used"`
	GasPrice              string        `json:"gas_price"`
	Transactions          []Transaction `json:"transactions"`
}

type Blocks struct {
	Blocks            []Block `json:"blocks"`
	LatestBlockNumber int     `json:"latest_block_number"`
}

type customTransport struct {
	http.RoundTripper
	//AuthorizationBearerToken string
}

var (
	fromaddress                                  string
	blockNumber, lowblockNumber, highblockNumber int
	bundleindex                                  int
)

func main() {
	flag.StringVar(&fromaddress, "address", "", "Parse the blocks transactions by the sender address")
	flag.IntVar(&blockNumber, "blockNumber", -1, "Mev block number")
	flag.IntVar(&lowblockNumber, "lowblockNumber", -1, "The lower bound block number")
	flag.IntVar(&highblockNumber, "highblockNumber", -1, "The higher bound block number")
	flag.IntVar(&bundleindex, "bundleindex", -1, "Get the indexed bundled transactions")
	flag.Parse()
	var blocks []Blocks
	if highblockNumber >= 0 {
		if lowblockNumber < 0 {
			lowblockNumber = 0
		}
		for i := lowblockNumber; i <= highblockNumber; i++ {
			blocksInter := getMevInfofromBlock(i)
			if len(blocksInter.Blocks) > 0 {
				blocks = append(blocks, blocksInter)
			}
		}
	} else {
		if blockNumber < 0 {
			fmt.Println("Give a vaild blockNumber")
			return
		} else {
			blocksInfo := getMevInfofromBlock(blockNumber)
			blocks = append(blocks, blocksInfo)
		}
	}
	for _, blocksInfo := range blocks {
		var targetTx []Transaction
		if bundleindex >= 0 && fromaddress == "" {
			targetTx = getTxByBundleIndex(blocksInfo, bundleindex)
		}
		if fromaddress != "" && bundleindex < 0 {
			targetTx, _ = getTxByBundleFrom(blocksInfo, fromaddress)
		}
		if fromaddress != "" && bundleindex >= 0 {
			_, index := getTxByBundleFrom(blocksInfo, fromaddress)
			for _, _index := range index {
				if _index == bundleindex {
					targetTx = getTxByBundleIndex(blocksInfo, _index)
				}
			}
		}
		resTx := parseTransactionArray(targetTx)
		fmt.Println(resTx)
	}
}

func getTxByBundleIndex(blocks Blocks, index int) []Transaction {
	resTx := make([]Transaction, 0)
	for _, Tx := range blocks.Blocks[0].Transactions {
		if Tx.BundleIndex == index {
			resTx = append(resTx, Tx)
		}
	}
	return resTx
}

func getTxByBundleFrom(blocks Blocks, from string) ([]Transaction, []int) {
	resTx := make([]Transaction, 0)
	indexArray := make([]int, 0)
	mapIndex := make(map[int]bool)
	for _, Tx := range blocks.Blocks[0].Transactions {
		if Tx.EoaAddress == from {
			if !mapIndex[Tx.BundleIndex] {
				indexArray = append(indexArray, Tx.BundleIndex)
			}
			mapIndex[Tx.BundleIndex] = true
			//resTx[from] = append(resTx[Tx.EoaAddress], Tx)
		}
	}
	for _, index := range indexArray {
		resTx = append(resTx, getTxByBundleIndex(blocks, index)...)
	}
	return resTx, indexArray
}

func parseTransactionArray(tx []Transaction) []string {
	resTx := make([]string, 0)
	for _, restx := range tx {
		gasPrice, _ := strconv.Atoi(restx.GasPrice)
		gasFee := restx.GasUsed * gasPrice
		fbalance := new(big.Float)
		fbalance.SetString(strconv.Itoa(gasFee))
		fbalance = fbalance.Quo(fbalance, big.NewFloat(math.Pow10(18)))
		transJSON, _ := json.Marshal(restx)
		bytegasFee := `,"gas_fee_to_Miner":` + fbalance.String()
		transJSON = append(append(transJSON[:len(transJSON)-1], []byte(bytegasFee)...), transJSON[len(transJSON)-1:]...)
		//resTx = append(resTx, string(transJSON))
		ntransJSON, _ := prettyPrint(string(transJSON))
		resTx = append(resTx, ntransJSON)
	}
	return resTx
}

func prettyPrint(str string) (string, error) {
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(str), &obj)
	if err != nil {
		return "", err
	}
	data, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func getMevInfofromBlock(blockNumber int) Blocks {
	url := "https://blocks.flashbots.net/v1/blocks?block_number={blockNumber}"
	url = strings.ReplaceAll(url, "{blockNumber}", strconv.Itoa(blockNumber))
	req, _ := http.NewRequest("GET", url, nil)
	fcdns := NewCachedDNS(dnsCacheRefreshInterval)
	client := &http.Client{
		Transport: otelhttp.NewTransport(&customTransport{
			RoundTripper: &http.Transport{
				MaxIdleConns:        100, //最大空闲连接数
				MaxConnsPerHost:     10,
				MaxIdleConnsPerHost: 10,               //每个主机最大空闲连接数
				IdleConnTimeout:     30 * time.Second, //空闲连接超时时间
				DialContext:         fcdns.dialWithCachedDNS,
			},
		}),
		Timeout: 10 * time.Second, // 设置超时时间
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("get request failed:", err)
		//return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error get:%d\n", resp.StatusCode)
		//return false
	}

	blocks := Blocks{}
	err = json.NewDecoder(resp.Body).Decode(&blocks)
	if err != nil {
		fmt.Println(err)
	}
	return blocks
}
