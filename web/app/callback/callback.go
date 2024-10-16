package callback

import (
	"GoFiberAPI/platform/authenticator"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
)

// Define a session store for Fiber
var store = session.New()

// Handler for the callback.
func Handler(auth *authenticator.Authenticator) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Get the session
		sess := store.Get(ctx)

		// State validation to prevent CSRF attacks
		if ctx.Query("state") != sess.Get("state") {
			return ctx.Status(fiber.StatusBadRequest).SendString("Invalid state parameter.")
		}

		// Exchange the authorization code for a token
		token, err := auth.Exchange(ctx.Context(), ctx.Query("code"))
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).SendString("Failed to exchange authorization code for a token.")
		}

		// Verify the ID token
		idToken, err := auth.VerifyIDToken(ctx.Context(), token)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to verify ID Token.")
		}

		// Extract user profile from the ID token
		var profile map[string]interface{}
		if err := idToken.Claims(&profile); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Store the access token and profile in the session
		sess.Set("access_token", token.AccessToken)
		sess.Set("profile", profile)

		// Save session
		if err := sess.Save(); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to save session.")
		}

		// Redirect to the logged-in page
		return ctx.Redirect("/health", fiber.StatusTemporaryRedirect)
	}
}
