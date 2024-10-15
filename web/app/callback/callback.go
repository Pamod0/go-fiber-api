package callback

import (
	"GoFiberAPI/platform/authenticator"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store = session.New()

// type Authenticator struct {
//     // Add necessary fields and methods
// }

// func (auth *Authenticator) Exchange(ctx context.Context, code string) (*Token, error) {
//     // Implement this method to exchange the authorization code for a token
//     return &Token{}, nil
// }

// func (auth *Authenticator) VerifyIDToken(ctx context.Context, token *Token) (*IDToken, error) {
//     // Implement this method to verify the ID token
//     return &IDToken{}, nil
// }

// type Token struct {
//     AccessToken string
// }

// type IDToken struct {
//     // Add necessary fields and methods
// }

// func (idToken *IDToken) Claims(v interface{}) error {
//     // Implement this method to extract claims from the ID token
//     return nil
// }

func Handler(auth *authenticator.Authenticator) fiber.Handler {
    return func(ctx *fiber.Ctx) error {
        sess, err := store.Get(ctx)
        if err != nil {
            return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
        }

        if ctx.Query("state") != sess.Get("state") {
            return ctx.Status(http.StatusBadRequest).SendString("Invalid state parameter.")
        }

        // Exchange an authorization code for a token.
        token, err := auth.Exchange(ctx.Context(), ctx.Query("code"))
        if err != nil {
            return ctx.Status(http.StatusUnauthorized).SendString("Failed to convert an authorization code into a token.")
        }

        idToken, err := auth.VerifyIDToken(ctx.Context(), token)
        if err != nil {
            return ctx.Status(http.StatusInternalServerError).SendString("Failed to verify ID Token.")
        }

        var profile map[string]interface{}
        if err := idToken.Claims(&profile); err != nil {
            return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
        }

        sess.Set("access_token", token.AccessToken)
        sess.Set("profile", profile)
        if err := sess.Save(); err != nil {
            return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
        }

        // Redirect to logged in page.
        return ctx.Redirect("/user", http.StatusTemporaryRedirect)
    }
}

// func main() {
//     app := fiber.New()

//     auth := &Authenticator{}
//     app.Get("/callback", Handler(auth))

//     app.Listen(":3000")
// }