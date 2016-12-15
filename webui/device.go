package webui

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ranjib/reefer/controller"
	"github.com/ranjib/reefer/controller/raspi"
	"log"
	"net/http"
)

func (h *APIHandler) ListDevices(w http.ResponseWriter, r *http.Request) {
	devices, err := h.controller.Devices().List()
	if err != nil {
		errorResponse(http.StatusInternalServerError, "Failed to retrieve device list", w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(devices); err != nil {
		errorResponse(http.StatusInternalServerError, "Failed to json decode. Error: "+err.Error(), w)
	}
}

func (h *APIHandler) CreateDevice(w http.ResponseWriter, r *http.Request) {
	var dd raspi.DeviceDetails
	if err := json.NewDecoder(r.Body).Decode(&dd); err != nil {
		errorResponse(http.StatusBadRequest, err.Error(), w)
		return
	}
	if err := h.controller.Devices().Create(dd); err != nil {
		errorResponse(http.StatusBadRequest, err.Error(), w)
		return
	}

}
func (h *APIHandler) GetDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["id"]
	d, err := h.controller.Devices().Get(name)
	if err != nil {
		errorResponse(http.StatusBadRequest, "Failed to retrieve specified device", w)
		return
	}
	dev, ok := d.(controller.Device)
	if !ok {
		errorResponse(http.StatusInternalServerError, "Failed to type cast device", w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(dev.Config()); err != nil {
		errorResponse(http.StatusInternalServerError, "Failed to json decode. Error: "+err.Error(), w)
	}
}
func (h *APIHandler) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["id"]
	if err := h.controller.Devices().Delete(name); err != nil {
		errorResponse(http.StatusBadRequest, "Failed to delete specified device", w)
		return
	}
}
func (h *APIHandler) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["id"]
	log.Println("TODO update device", name)
}
func (h *APIHandler) ConfigureDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["id"]
	type State struct {
		On bool `json:"on"`
	}
	d, err := h.controller.Devices().Get(name)
	if err != nil {
		errorResponse(http.StatusBadRequest, "Failed to retrieve specified device", w)
		return
	}
	dev, ok := d.(controller.Device)
	if !ok {
		errorResponse(http.StatusBadRequest, "Failed to typecast to device", w)
		return
	}
	var s State
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		errorResponse(http.StatusBadRequest, err.Error(), w)
		return
	}
	if s.On {
		if err := dev.On(); err != nil {
			errorResponse(http.StatusBadRequest, err.Error(), w)
			return
		}
		return
	}
	if err := dev.Off(); err != nil {
		errorResponse(http.StatusBadRequest, err.Error(), w)
		return
	}

}