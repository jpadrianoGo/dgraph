// Copyright 2019 ChainSafe Systems (ON) Corp.
// This file is part of gossamer.
//
// The gossamer library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gossamer library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gossamer library. If not, see <http://www.gnu.org/licenses/>.

package dot

import (
	"testing"

	"github.com/ChainSafe/gossamer/lib/utils"
	"github.com/stretchr/testify/require"
)

// TestLoadConfig tests loading a toml configuration file
func TestLoadConfig(t *testing.T) {
	cfg, cfgFile := NewTestConfigWithFile(t)
	require.NotNil(t, cfg)

	genFile := NewTestGenesisFile(t, cfg)
	require.NotNil(t, genFile)

	defer utils.RemoveTestDir(t)

	cfg.Global.Genesis = genFile.Name()

	err := InitNode(cfg)
	require.Nil(t, err)

	err = LoadConfig(cfg, cfgFile.Name())
	require.Nil(t, err)

	// TODO: improve dot config tests
	require.NotNil(t, cfg)
}

// TestExportConfig tests exporting a toml configuration file
func TestExportConfig(t *testing.T) {
	cfg, cfgFile := NewTestConfigWithFile(t)
	require.NotNil(t, cfg)

	genFile := NewTestGenesisFile(t, cfg)
	require.NotNil(t, genFile)

	defer utils.RemoveTestDir(t)

	cfg.Global.Genesis = genFile.Name()

	err := InitNode(cfg)
	require.Nil(t, err)

	file := ExportConfig(cfg, cfgFile.Name())

	// TODO: improve dot config tests
	require.NotNil(t, file)
}

// Gssmr Node

// TestLoadConfigGssmr tests loading the toml configuration file for ksmcc
func TestLoadConfigGssmr(t *testing.T) {
	cfg := GssmrConfig()
	require.NotNil(t, cfg)

	cfg.Global.DataDir = utils.NewTestDir(t)
	cfg.Global.Genesis = "../node/gssmr/genesis.json"

	defer utils.RemoveTestDir(t)

	err := InitNode(cfg)
	require.Nil(t, err)

	err = LoadConfig(cfg, "../node/gssmr/config.toml")
	require.Nil(t, err)

	// TODO: improve dot config tests
	require.NotNil(t, cfg)
}

// TestExportConfigGssmr tests exporting the toml configuration file
func TestExportConfigGssmr(t *testing.T) {
	cfg := GssmrConfig()
	require.NotNil(t, cfg)

	gssmrConfig := cfg.Global.Config
	gssmrGenesis := cfg.Global.Genesis
	gssmrDataDir := cfg.Global.DataDir
	cfg.Global.DataDir = utils.NewTestDir(t)
	cfg.Global.Genesis = "../node/gssmr/genesis.json"

	defer utils.RemoveTestDir(t)

	err := InitNode(cfg)
	require.Nil(t, err)

	cfg.Global.Config = gssmrConfig
	cfg.Global.Genesis = gssmrGenesis
	cfg.Global.DataDir = gssmrDataDir

	file := ExportConfig(cfg, "../node/gssmr/config.toml")

	// TODO: improve dot config tests
	require.NotNil(t, file)
}

// Ksmcc Node

// TestLoadConfigKsmcc tests loading the toml configuration file for ksmcc
func TestLoadConfigKsmcc(t *testing.T) {
	cfg := KsmccConfig()
	require.NotNil(t, cfg)

	cfg.Global.DataDir = utils.NewTestDir(t)
	cfg.Global.Genesis = "../node/ksmcc/genesis.json"

	defer utils.RemoveTestDir(t)

	err := InitNode(cfg)
	require.Nil(t, err)

	err = LoadConfig(cfg, "../node/ksmcc/config.toml")

	// TODO: improve dot config tests
	require.Nil(t, err)
}

// TestExportConfigKsmcc tests exporting the toml configuration file
func TestExportConfigKsmcc(t *testing.T) {
	cfg := KsmccConfig()
	require.NotNil(t, cfg)

	ksmccConfig := cfg.Global.Config
	ksmccGenesis := cfg.Global.Genesis
	ksmccDataDir := cfg.Global.DataDir
	cfg.Global.DataDir = utils.NewTestDir(t)
	cfg.Global.Genesis = "../node/ksmcc/genesis.json"

	defer utils.RemoveTestDir(t)

	err := InitNode(cfg)
	require.Nil(t, err)

	cfg.Global.Config = ksmccConfig
	cfg.Global.Genesis = ksmccGenesis
	cfg.Global.DataDir = ksmccDataDir

	file := ExportConfig(cfg, "../node/ksmcc/config.toml")

	// TODO: improve dot config tests
	require.NotNil(t, file)
}