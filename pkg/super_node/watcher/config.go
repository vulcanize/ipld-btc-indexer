// VulcanizeDB
// Copyright © 2019 Vulcanize

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package watcher

import (
	"github.com/vulcanize/vulcanizedb/pkg/config"
	"github.com/vulcanize/vulcanizedb/pkg/eth/core"
	"github.com/vulcanize/vulcanizedb/pkg/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/super_node/shared"
)

// Config holds all of the parameters necessary for defining and running an instance of a watcher
type Config struct {
	// Subscription settings
	SubscriptionConfig shared.SubscriptionSettings
	// Database settings
	DBConfig config.Database
	// DB itself
	DB *postgres.DB
	// Subscription client
	Client core.RPCClient
	// WASM instantiation paths and namespaces
	WASMInstances [][2]string
	// Path and names for trigger functions (sql files) that (can) use the instantiated wasm namespaces
	TriggerFunctions [][2]string
}