package config

import (
	"fmt"

	"github.com/konstellation-io/kli/internal/render"
	"github.com/konstellation-io/kli/text"
)

func (c *Config) RenderServerList(r render.Renderer) {
	r.SetHeader([]string{"Server", "URL"})

	for _, s := range c.ServerList {
		defaultMark := ""
		isDefault := text.Normalize(s.Name) == c.DefaultServer

		if isDefault {
			defaultMark = "*"
		}

		r.Append([]string{
			fmt.Sprintf("%s%s", s.Name, defaultMark),
			s.URL,
		})
	}

	r.Render()
}
