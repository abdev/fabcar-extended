package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/types"
	"github.com/tendermint/tmlibs/pubsub/query"
)

const rpcListenAddress = "tcp://0.0.0.0:46657"
const webSocket = "/websocket"

func getHTTPClient() *client.HTTP {
	return client.NewHTTP(rpcListenAddress, webSocket)
}

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

func startClient() error {
	client := getHTTPClient()
	txs := doSubscribe(client)

	go func() {
		for e := range txs {
			data := e.(types.TMEventData).Unwrap().(types.EventDataTx)
			fmt.Println("got ", data.TxResult.Result.GetInfo())
		}
	}()

	//deliver car record
	data := []byte(`{"operation": "createCar", "data": {"ID": "car21", "Make": "Toyota", "Model": "Prius", "Colour": "black", "Owner": "Jane"}}`)

	tx := types.Tx([]byte(base64.StdEncoding.EncodeToString(data)))

	fmt.Println(base64.StdEncoding.EncodeToString(data))

	fmt.Println("The transaction is", tx)

	resultDeliver, err := client.BroadcastTxCommit(tx)

	if err != nil {
		panic(err)
	}

	fmt.Println("The deliver result is:", resultDeliver)

	resultStatus, err := client.ABCIQuery("allCars", nil)

	fmt.Println("Query all car", string(resultStatus.Response.Value))

	if err != nil {
		fmt.Println("we have an error in query")
		panic(err)
	}

	//log.Println(httpClient)

	return nil
}

func main() {
	err := startClient()

	if err != nil {
		panic(err)
	}
}
