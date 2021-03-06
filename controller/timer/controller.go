package timer

import (
	"github.com/reef-pi/reef-pi/controller/equipments"
	"github.com/reef-pi/reef-pi/controller/utils"
	"gopkg.in/robfig/cron.v2"
	"log"
)

type Controller struct {
	store      utils.Store
	runner     *cron.Cron
	cronIDs    map[string]cron.EntryID
	telemetry  *utils.Telemetry
	equipments *equipments.Controller
}

func New(store utils.Store, telemetry *utils.Telemetry, e *equipments.Controller) *Controller {
	return &Controller{
		runner:     cron.New(),
		cronIDs:    make(map[string]cron.EntryID),
		telemetry:  telemetry,
		store:      store,
		equipments: e,
	}
}

func (c *Controller) Setup() error {
	return c.store.CreateBucket(Bucket)
}

func (c *Controller) Start() {
	if err := c.loadAllJobs(); err != nil {
		log.Println("ERROR: Failed to load timer jobs. Error:", err)
	}
	c.runner.Start()
}

func (c *Controller) Stop() {
	c.runner.Stop()
}
