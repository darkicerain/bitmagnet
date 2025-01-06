package httpserver

import (
	"bitmagnet-io/bitmagnet/internal/boilerplate/httpserver"
	"bitmagnet-io/bitmagnet/internal/boilerplate/lazy"
	"bitmagnet-io/bitmagnet/internal/database/dao"
	"bitmagnet-io/bitmagnet/internal/protocol"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In
	Dao    lazy.Lazy[*dao.Query]
	Logger *zap.SugaredLogger
}

type Result struct {
	fx.Out
	Option httpserver.Option `group:"http_server_options"`
}

func New(p Params) Result {
	return Result{
		Option: &builder{
			dao:    p.Dao,
			logger: p.Logger.Named("downloader"),
		},
	}
}

type builder struct {
	dao    lazy.Lazy[*dao.Query]
	logger *zap.SugaredLogger
}

func (b *builder) Key() string {
	return "downloader"
}

func (b *builder) Apply(e *gin.Engine) error {

	e.GET("/downloader", func(c *gin.Context) {

		hash := c.Query("hash")

		if hash != "" {

			infoHash, err := protocol.ParseID(hash)
			if err != nil {
				c.Status(500)
				c.Writer.WriteString(err.Error())
				return
			}
			c.Status(200)
			c.Header("Content-Type", "application/x-bittorrent; charset=utf-8")
			c.Header("Content-Disposition", "attachment; filename=\""+hash+".torrent\"")
			get, _ := b.dao.Get()
			dotFile := get.TorrentTorrentDotFile
			take, err := dotFile.Where(dotFile.InfoHash.Eq(infoHash)).Take()
			if err == nil {
				_, err := c.Writer.Write(take.BinaryFile)
				if err != nil {
					c.Status(500)
					_, err := c.Writer.WriteString(err.Error())
					if err != nil {
						return
					}
					return
				}
				return
			}
		}
		c.Status(500)
		_, err := c.Writer.WriteString("hash is empty")
		if err != nil {
			return
		}
	})
	return nil
}
