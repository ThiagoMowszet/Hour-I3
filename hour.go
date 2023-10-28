package main

import (
	"fmt"
	"os/exec"
	"encoding/json"
	"net/http"
	"time"
	_"strings"
)

type WorldTimeResponse struct {
	Datetime string `json:"datetime"`
}

func obtenerHoraArgentina() (string, error) {
	resp, err := http.Get("http://worldtimeapi.org/api/timezone/America/Argentina/Cordoba")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data WorldTimeResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	return data.Datetime, nil
}

func main() {
	hora, err := obtenerHoraArgentina()
	if err != nil {
		fmt.Println("Error al obtener la hora de Argentina:", err)
		return
	}

	// Parsear la hora obtenida
	t, err := time.Parse(time.RFC3339Nano, hora)
	if err != nil {
		fmt.Println("Error al parsear la hora:", err)
		return
	}

	// Formatear la hora en el formato correcto
	formattedTime := t.Format("2006-01-02 15:04:05")

	// Ejecutar el comando para actualizar la hora
	cmd := exec.Command("sudo", "timedatectl", "set-time", formattedTime)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error al cambiar la hora:", err)
	} else {
		fmt.Println("Hora actualizada a:", formattedTime)
	}

	// Cambiar la zona horaria
	cmd = exec.Command("sudo", "timedatectl", "set-timezone", "America/Argentina/Cordoba")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error al cambiar la zona horaria:", err)
	} else {
		fmt.Println("Zona horaria cambiada a America/Argentina/Cordoba")
	}
}
