package login

import (
	"encoding/base64"
	"math/rand"
	"net/http"
	"GoFiberAPI/platform/authenticator"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store = session.New()

func Handler(auth *authenticator.Authenticator) fiber.Handler {
    return func(ctx *fiber.Ctx) error {
        state, err := generateRandomState()
        if err != nil {
            return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
        }

        // Save the state inside the session.
        sess, err := store.Get(ctx)
        if err != nil {
            return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
        }
        sess.Set("state", state)
        if err := sess.Save(); err != nil {
            return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
        }

        return ctx.Redirect(auth.AuthCodeURL(state), http.StatusTemporaryRedirect)
    }
}

func generateRandomState() (string, error) {
    b := make([]byte, 32)
    _, err := rand.Read(b)
    if err != nil {
        return "", err
    }

    state := base64.StdEncoding.EncodeToString(b)

    return state, nil
}

// func main() {
//     app := fiber.New()

//     auth := &Authenticator{}
//     app.Get("/login", Handler(auth))

//     app.Listen(":3000")
// }