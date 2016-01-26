# quiz


Q: Given a list of words like https://github.com/NodePrime/quiz/blob/master/word.list find the longest compound-word in the list, which is also a concatenation of other sub-words that exist in the list. The program should allow the user to input different data. The finished solution shouldn't take more than one hour. Any programming language can be used, but Go is preferred.


Fork this repo, add your solution and documentation on how to compile and run your solution, and then issue a Pull Request. 

Obviously, we are looking for a fresh solution, not based on others' code.

---

# Usage

```bash
go get github.com/jatinn/quiz
cd $GOPATH/src/github.com/jatinn/quiz
go build
./quiz word.list
```

## Notes
1. Supports concurrency. By default spawns workers equal to the number of cores available.
2. Can modify number of worker with a cpu multiplier flag.

	```bash
	./quiz -m=10 word.list
	```

3. Processes the complete list so that if there are multiple compoun words with the same length, it will list all of them.
4. List does no need to be sorted and handles unicode as well.

	```bash
	./quiz sample.list
	Found 2 longest compund words of length 5:
		 abcde
		 ↴↴↴↴↴
	Time to completion: 359.208µs
	```
