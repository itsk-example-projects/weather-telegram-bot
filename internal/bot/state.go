package bot

import (
	"sync"

	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type state struct {
	rwMux    sync.RWMutex
	userData map[int64]map[string]any
}

func (c *state) getUserData(ctx *ext.Context, key string) (any, bool) {
	c.rwMux.RLock()
	defer c.rwMux.RUnlock()

	if c.userData == nil {
		return nil, false
	}

	userData, ok := c.userData[ctx.EffectiveUser.Id]
	if !ok {
		return nil, false
	}

	v, ok := userData[key]
	return v, ok
}

func (c *state) setUserData(ctx *ext.Context, key string, val any) {
	c.rwMux.Lock()
	defer c.rwMux.Unlock()

	if c.userData == nil {
		c.userData = map[int64]map[string]any{}
	}

	_, ok := c.userData[ctx.EffectiveUser.Id]
	if !ok {
		c.userData[ctx.EffectiveUser.Id] = map[string]any{}
	}
	c.userData[ctx.EffectiveUser.Id][key] = val
}
