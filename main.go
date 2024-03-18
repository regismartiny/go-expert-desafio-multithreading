package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/regismartiny/go-expert-desafio-multithreading/internal/brasilapi"
	"github.com/regismartiny/go-expert-desafio-multithreading/internal/viacep"
)

func main() {

	cep := os.Args[1]

	canalBrasilapi := make(chan brasilapi.CepInfo)
	canalViacep := make(chan viacep.CepInfo)

	ctx := context.Background()

	go func() {
		brasilapicepInfo, err := brasilapi.NewClient().GetCepInfo(&ctx, cep)
		if err != nil {
			log.Println("Ocorreu um erro ao consultar a api Brasilapi:", err)
		}
		canalBrasilapi <- brasilapicepInfo
	}()

	go func() {
		viacepCepInfo, err := viacep.NewClient().GetCepInfo(&ctx, cep)
		if err != nil {
			log.Println("Ocorreu um erro ao consultar a api Viacep:", err)
		}
		canalViacep <- viacepCepInfo
	}()

	select {
	case msg := <-canalBrasilapi:
		fmt.Println("Retornando dados da api Brasilapi:\n", structToPrettyString(msg))

	case msg := <-canalViacep:
		fmt.Println("Retornando dados da api Viacep:\n", structToPrettyString(msg))

	case <-time.After(time.Second * 1):
		println("timeout")
	}

}

func structToPrettyString(obj interface{}) string {
	s, _ := json.MarshalIndent(obj, "", "\t")
	return string(s)
}
