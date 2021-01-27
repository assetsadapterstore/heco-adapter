/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package openwtester

import (
	"github.com/astaxie/beego/config"
	"github.com/blocktree/openwallet/v2/openw"
	"path/filepath"
	"testing"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openwallet"
)


func TestWalletManager_GetAssetsAccountBalance(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "W6pCLPenzu1XuHx9gazjitoyUXNoJnkapj"
	accountID := "5ry4GxRf3QosSFZ9e1ouj4duCrTpMstoSKJ2MPmyfFsk"
	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance)
}

func TestWalletManager_GetAssetsAccountTokenBalance(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "W4Fyh8Voz6R1iGf5bugscqUyYSsNMVoHh7"
	//accountID := "3dv8NH8xMUXAo2zMwZ6gAJqs74dYLg5zFSf3FVEDLdVv"
	accountID := "CxRfMAt2gwUiCxZy823H4apWKaHR3UfgtiDcXg2DDJPV"

	contract := openwallet.SmartContract{
		Address:  "0x35def49e4c26aadd2a0734f38f07dc6c7993f764",
		Symbol:   "HT",
		Name:     "HT",
		Token:    "HT",
		Decimals: 18,
	}

	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance.Balance)
}

func TestWalletManager_GetEstimateFeeRate(t *testing.T) {
	tm := testInitWalletManager()
	coin := openwallet.Coin{
		Symbol: "HT",
	}
	feeRate, unit, err := tm.GetEstimateFeeRate(coin)
	if err != nil {
		log.Error("GetEstimateFeeRate failed, unexpected error:", err)
		return
	}
	log.Std.Info("feeRate: %s %s/%s", feeRate, coin.Symbol, unit)
}


func TestGetAddressVerify(t *testing.T) {
	symbol := "HT"
	assetsMgr, err := openw.GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}
	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)
	addrDec := assetsMgr.GetAddressDecoderV2()

	flag := addrDec.AddressVerify("0x4402a2969da0689a0e6f5fbad8be930430b4ad63af25f3c93dbd03bb40908d08")
	log.Infof("flag: %v, expect: false", flag)

	flag = addrDec.AddressVerify("6541a59bd17cf20f058e8b5377f034a32843410f")
	log.Infof("flag: %v, expect: false", flag)

	flag = addrDec.AddressVerify("0x562ff43493d6a2baf88358b38cbc268b7cbb8a89")
	log.Infof("flag: %v, expect: true", flag)

}