package routers

import (
	"test/pkg/router"

	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRouter(router *router.Router, db *mongo.Database) *router.Router {

	// noteRouter := router.Group("/note")
	// router *router.Router,.POST("/login", publicuserlogin.New(db, sendgrid).Login)
	// authRouter.PATCH("/login/verification/:email", publicuserloginverification.New(db, authModule, tokenModule).VerificationOTP)
	// authRouter.POST("/login/resend", publicuserloginresend.New(db, sendgrid).Login)

	return router
}
