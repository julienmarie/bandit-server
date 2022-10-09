package strategies

import "github.com/julienmarie/bandit-server/repository"

type Strategy interface {
	Choose(repo repository.Repository, context string, experiments []string) string
}
