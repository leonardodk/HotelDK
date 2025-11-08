package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache      bool                          // bool to say whether to use template cache
	TemplateCache map[string]*template.Template // a map -it's the template cache
	InfoLog       log.Logger                    // no idea yet
	InProduction  bool                          // bool to say whether to be in in production mode
	Session       *scs.SessionManager           // a session manager
}
