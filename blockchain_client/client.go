package blockchain_client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/abdev/fabcar-extended/web_app/models"
	"github.com/gobuffalo/uuid"
	"github.com/tendermint/tendermint/rpc/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
	"github.com/tendermint/tmlibs/pubsub/query"
)

const rpcListenAddress = "tcp://0.0.0.0:46657"
const webSocket = "/websocket"

type CarAsset struct {
	ID     uuid.UUID `json:"ID"`
	Make   string    `json:"Make"`
	Model  string    `json:"Model"`
	Colour string    `json:"Colour"`
	Owner  string    `json:"Owner"`
}

type TransactionCreateCar struct {
	Operation string   `json:"operation"`
	Data      CarAsset `json:"data"`
}

type TransactionResult struct {
	Tx       types.Tx
	Response *ctypes.ResultBroadcastTxCommit
	Data     string
}

func getHTTPClient() *client.HTTP {
	return client.NewHTTP(rpcListenAddress, webSocket)
}

var httpClient = getHTTPClient()

func doSubscribe(client *client.HTTP) chan interface{} {
	//do subscribe
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	query := query.MustParse("tm.event = 'Tx'")

	txs := make(chan interface{})

	client.WSEvents.Start()

	fmt.Println(client.WSEvents.IsRunning())

	err := client.Subscribe(ctx, "test-client", query, txs)

	if err != nil {
		fmt.Println("we have an error in client subscribe")
		panic(err)
	}

	return txs
}

func StartClient() error {
	txs := doSubscribe(httpClient)

	go func() {
		for e := range txs {
			data := e.(types.TMEventData).Unwrap().(types.EventDataTx)
			tags, _ := json.Marshal(data.Result.GetTags())
			fmt.Println("Response from tendermint ", string(tags))
		}
	}()

	resultStatus, err := httpClient.ABCIQuery("allCars", nil)

	if err != nil {
		fmt.Println("we have an error in query")
		panic(err)
	}

	fmt.Println("Query all car", string(resultStatus.Response.Value))

	//log.Println(httpClient)

	return nil
}

func CreateCar(c models.Car) (*TransactionResult, error) {
	transactionResult := &TransactionResult{}

	carAsset := &CarAsset{ID: c.ID, Make: c.Make, Model: c.Model, Colour: c.Colour, Owner: c.Owner.Name}
	data := TransactionCreateCar{Operation: "createCar", Data: *carAsset}

	payload, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	log.Println("The payload is", string(payload))

	//deliver car record
	tx := types.Tx([]byte(base64.StdEncoding.EncodeToString(payload)))

	fmt.Println("The transaction is", tx)

	resultDeliver, err := httpClient.BroadcastTxCommit(tx)

	if err != nil {
		panic(err)
	}

	fmt.Println("The deliver result is:", resultDeliver.DeliverTx.Info)

	transactionResult.Tx = tx
	transactionResult.Response = resultDeliver
	transactionResult.Data = string(payload)

	return transactionResult, nil
}
