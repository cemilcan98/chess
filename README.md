### Run
```
Applications used to run the chess application is Docker, Goland (application which Go runs on), and terminal. 
Make sure to have Go (programming language) downloaded to your computer.

The rest API was written in Golang.

Once everything is downloaded follow the next steps:
1) Open Goland
2) cd chess/api
3) Use the command: make env-dev-start
4) go run main.go import -f jj_games_pgn_updated.pgn
5) go run main.go api
6) python -m SimpleHTTPServer 5000
 
             OR
 
   python -m http.server 5000

   http://localhost:5000/ 


 ```

