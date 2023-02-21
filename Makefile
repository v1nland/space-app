.DEFAULT_GOAL := help
.SILENT:

help:
	echo "Usage: make <command>\n"
	echo "Available commands:"
	echo "    gqlgen                Create auto-generated graphql files"
	echo "    help                  This help :P"

gqlgen:
	go run github.com/99designs/gqlgen generate
