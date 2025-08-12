package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Load menginisialisasi Viper untuk membaca konfigurasi dari Env dan file.
// Prioritas konfigurasi: Env > file.
func Load() *viper.Viper {
	v := viper.New()

	// -- Preferensi format file --
	v.SetConfigName("config") // nama file tanpa ekstensi
	v.SetConfigType("yaml")
	v.AddConfigPath(".")	// cari di root server

	// -- Default value agar aman saat dev --
	v.SetDefault("app.port", 8080)
	v.SetDefault("app.env", "development")
	v.SetDefault("db.driver", "mysql")
	v.SetDefault("security.jwt_secret", "CHANGE_ME_DEV_ONLY") // ganti di prod!

	// -- Bind ENV: gunakan prefix BLOGAPP_, contoh: BLOGAPP_PORT
	v.SetEnvPrefix("BLOGAPP")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// Baca file config.yaml jika ada (opsional)
	if err := v.ReadInConfig(); err != nil {
		log.Printf("[config] skip read file: %v", err)
	}
	
	return v
} 