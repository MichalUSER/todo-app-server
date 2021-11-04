curl --header "Content-Type: application/json" \
	--request POST \
	--data '{"title": "Learn EE", "description": "use rust", "completed": false, "height": 50}' \
	http://localhost:8080/add/michal

curl http://localhost:8080/todos/michal
