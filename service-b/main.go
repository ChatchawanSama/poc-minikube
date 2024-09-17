package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

type InternalConfigResponse struct {
	SecretData   string `json:"secretData"`
	ConfigMapVal string `json:"configMapVal"`
}

func main() {
	viper.AutomaticEnv()
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	http.HandleFunc("/internal/config", func(w http.ResponseWriter, r *http.Request) {
		// Read SECRET_KEY and CONFIG_MAP_KEY from environment variables
		secretData := viper.GetString("SECRET_KEY")
		configMapVal := viper.GetString("CONFIG_MAP_KEY")

		fmt.Println("Secret Data: ", secretData)
		fmt.Println("Config Map Value: ", configMapVal)

		config := InternalConfigResponse{
			SecretData:   secretData,
			ConfigMapVal: configMapVal,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(config)
	})

	fmt.Println("Service running on port 8081")
	http.ListenAndServe(":8081", nil)
}
