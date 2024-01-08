package web

import (
	"embed"
	"fmt"
	"log"
	"mime"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

//go:generate pnpm install
//go:generate pnpm build

//go:embed all:dist/*
var dist embed.FS

func RunDevCmd() {
	cmd := exec.Command("pnpm", "dev")
	cmd.Dir = "./web"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
}

func ServeDev(c *fiber.Ctx) error {
	path := c.Path() // like `/`
	url := "http://localhost:3001" + path
	return proxy.Forward(url)(c)
}

func ServeDist(c *fiber.Ctx) error {
	path := c.Path() // like `/`
	path = fmt.Sprintf("dist%s", path)
	if strings.HasSuffix(path, "/") {
		path += "index.html"
	}

	f, err := dist.ReadFile(path)
	if err != nil {
		return err
	}
	fileExt := filepath.Ext(path)
	mimeType := mime.TypeByExtension(fileExt)
	c.Set(fiber.HeaderContentType, mimeType)

	return c.SendString(string(f))
}