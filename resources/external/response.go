package external

import (
	"github.com/ravielze/Labpro-BE/config"
	"github.com/ravielze/oculi/response"
	"github.com/ravielze/oculi/validator"
)

func NewResponder(v validator.Validator, config *config.Env) response.Responder {
	return response.New(v, config.IsDevelopment())
}
