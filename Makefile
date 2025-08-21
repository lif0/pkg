tidy:
	find . -name 'go.mod' -print -execdir sh -c 'echo "→ $$PWD"; go mod tidy' \;

generate-badge:
	sh .github/scripts/pre-push