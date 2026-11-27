package oci

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"testing"
)

func TestDockerConfigDir(t *testing.T) {
	homeEnv := "HOME"
	switch runtime.GOOS {
	case "windows":
		homeEnv = "USERPROFILE"
	case "plan9":
		homeEnv = "home"
	}

	t.Run("home directory", func(t *testing.T) {
		home := t.TempDir()
		t.Setenv(homeEnv, home)
		t.Setenv("DOCKER_CONFIG", "")
		if got, want := dockerConfigDir(), filepath.Join(home, ".docker"); got != want {
			t.Fatalf("config directory = %q, want %q", got, want)
		}
	})

	t.Run("explicit override", func(t *testing.T) {
		dir := t.TempDir()
		t.Setenv("DOCKER_CONFIG", dir)
		if got := dockerConfigDir(); got != dir {
			t.Fatalf("config directory = %q, want %q", got, dir)
		}
	})

	t.Run("account fallback", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("Docker does not use account lookup on Windows")
		}
		current, err := user.Current()
		if err != nil {
			t.Skipf("account lookup unavailable: %v", err)
		}
		t.Setenv(homeEnv, "")
		t.Setenv("DOCKER_CONFIG", "")
		if got, want := dockerConfigDir(), filepath.Join(current.HomeDir, ".docker"); got != want {
			t.Fatalf("config directory = %q, want %q", got, want)
		}
	})
}

func TestLoadDockerConfig(t *testing.T) {
	for _, tc := range []struct {
		name    string
		content string
		missing bool
		wantErr bool
	}{
		{name: "credentials", content: `{"auths":{"registry.example":{"auth":"dXNlcjpwYXNz"}}}`},
		{name: "missing config", missing: true},
		{name: "invalid config", content: "{", wantErr: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			t.Setenv("DOCKER_CONFIG", dir)
			filename := filepath.Join(dir, "config.json")
			if !tc.missing {
				if err := os.WriteFile(filename, []byte(tc.content), 0600); err != nil {
					t.Fatal(err)
				}
			}
			cfg, err := loadDockerConfig()
			if (err != nil) != tc.wantErr {
				t.Fatalf("load error = %v, want error: %v", err, tc.wantErr)
			}
			if tc.wantErr {
				return
			}
			if cfg.Filename != filename {
				t.Fatalf("loaded %q, want %q", cfg.Filename, filename)
			}
			if tc.missing {
				if len(cfg.AuthConfigs) != 0 {
					t.Fatal("missing config returned credentials")
				}
				return
			}
			auth, err := cfg.GetAuthConfig("registry.example")
			if err != nil {
				t.Fatal(err)
			}
			if auth.Username != "user" || auth.Password != "pass" {
				t.Fatal("registry credentials were not loaded")
			}
		})
	}
}
