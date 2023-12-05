package main

import (
	"fmt"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/intrnls/vhandler"
	"github.com/intrnls/vhandler/priority"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel

	cfg := server.DefaultConfig()
	srvCfg, err := cfg.Config(log)
	if err != nil {
		log.Fatal(err)
	}
	srv := srvCfg.New()
	srv.Listen()
	for {
		srv.Accept(func(p *player.Player) {
			setupHandler(p).Set(p)
		})
	}
}

func setupHandler(p *player.Player) *vhandler.Player {
	v := vhandler.NewPlayer()

	v.OnMove(priority.Normal, func(ctx *event.Context, newPos mgl64.Vec3, newYaw, newPitch float64) {
		p.SendTip(fmt.Sprintf("X: %.2f Y: %.2f Z: %.2f\nYaw: %.0f Pitch: %.0f", newPos.X(), newPos.Y(), newPos.Z(), newYaw, newPitch))
	})

	return v
}
