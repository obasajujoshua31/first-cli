install:
	go install practCLI
call-add:
	practCLI add 5.4 6.6 7.6 -f
add-even:
	practCLI add even 1 3 4 5 6 7 8 9
format:
	go fmt ./...
add-odd:
	practCLI add odd 1 3 4 5 8 9 4 6
login:
	practCLI login -e example@example.com -p example