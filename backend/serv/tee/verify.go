package tee

import (
	"crypto"
	"crypto/rsa"
	"fmt"
	"reflect"

	"github.com/ochanoco/proxy/core"
)

func verifyDB(pubkey *rsa.PublicKey, proxy *core.OchanocoProxy) error {
	fmt.Println("logger is verifing DB...")

	// todo: use db relation
	serviceLogs, err := proxy.Database.SelectAllServiceLogs()
	if err != nil {
		return err
	}

	hashChains, err := proxy.Database.SelectAllHashChains()
	if err != nil {
		return err
	}

	if len(serviceLogs) == 0 {
		fmt.Printf("no record!!!!\n")
		return nil
	}

	var previous core.HashChain

	for index, log := range serviceLogs[1:] {
		chain := hashChains[index]

		err = rsa.VerifyPSS(pubkey, crypto.SHA3_256, chain.Hash, chain.Signature, &rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthAuto,
		})

		if err != nil {
			return err
		}
		previous = core.HashChain{
			Hash: core.CalcHashChain(log, &previous),
		}
	}

	last, err := proxy.Database.FindLastHashChain()

	if err != nil {
		return err
	}

	if !reflect.DeepEqual(last.Hash, previous.Hash) {
		return fmt.Errorf("verifyDB: hash chain is invalid (%v, %v)", last.Hash, previous.Hash)
	}

	fmt.Printf("DB is valid (%v, %v)\n", last.Hash, previous.Hash)

	if err != nil {
		return err
	}

	return nil
}
