/*
 * Copyright 2018 Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package origins

import (
	"net/http"

	"github.com/tricksterproxy/trickster/pkg/cache"
	oo "github.com/tricksterproxy/trickster/pkg/proxy/origins/options"
	po "github.com/tricksterproxy/trickster/pkg/proxy/paths/options"
)

// Client is the primary interface for interoperating with Trickster and upstream TSDB's
type Client interface {
	// Handlers returns a map of the HTTP Handlers the client has registered
	Handlers() map[string]http.Handler
	// DefaultPathConfigs returns the default PathConfigs for the given OriginType
	DefaultPathConfigs(*oo.Options) map[string]*po.Options
	// Configuration returns the configuration for the Proxy Client
	Configuration() *oo.Options
	// Name returns the name of the origin the Proxy Client is handling
	Name() string
	// HTTPClient will return the HTTP Client for this Origin
	HTTPClient() *http.Client
	// SetCache sets the Cache object the client will use when caching origin content
	SetCache(cache.Cache)
	// Router returns a Router that handles HTTP Requests for this client
	Router() http.Handler
	// Cache returns a handle to the Cache instance used by the Client
	Cache() cache.Cache
}
