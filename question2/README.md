# Additional Note on the bug where open lead to an unexpected output 
- During the trouble shooting on the open function issue, I want to explain more on the process where i conclude issue is from line 8 from the `explain.py`

```python
fp = "./file.pin"
# using high level API and works perfectly
with open(fp, 'r') as f:
    for line in f:
        print(line.rstrip())

# override the high level API to low level API 
from os import *
# This will failed and still confused us we are using 
# high level API. 
with open(fp, 'r') as f:
    for line in f:
        print(line.rstrip())

```
> output of running `python3 explain.py`
```
Here is some text 1.2.3.4
Here is some text 123.41.2.0
SOME_TEXT 3.5.110.229
SOME_TEXT 83.5.3.229 SOME_TEXT 83.5.11000.229
Traceback (most recent call last):
  File "/Users/rockliang/RBC/interview/question2/root/devops/explain.py", line 11, in <module>
    with open(fp, 'r') as f:
         ^^^^^^^^^^^^^
TypeError: 'str' object cannot be interpreted as an integer
```
- The exception message suggested the line that is problemtic is open functions. And the google search suggested otherwise.

- It ticks me on we might be using a wrong functions as we are casting a string as an integer from line `'str' object cannot be interpreted as an integer`. 

- So i suspect is coming from `os` package where define another function with same name `open` and override with line 8 from the `explain.py`

## Show case where both methods exists but same name

- By default the `open` is a [high level API](https://docs.python.org/3/library/functions.html#open) which returns the file object with a stream of data. 

- The bug is caused by importing the [low level API](https://docs.python.org/3/library/os.html#os.open) from os and cause the confusion during trouble shooting. 


## Lession Learned
- We should always import the exact methods we are using that override defined functions.

- In addition, always distinguish the functions name using `Polymorphism` or Name to avoid unexpected blocker during the code.

# Golang Solutions
- Simply run `go run *.go` you should be able to view the result
```
1.2.3.0
1.2.3.4
1.2.3.4
1.2.30.4
1.22.3.4
3.1.110.229
3.5.110.29
3.5.110.229
8.5.0.229
8.5.110.229
76.8.17.126
83.5.1.229
83.5.1.229
83.5.3.229
83.5.11.2
83.5.110.2
83.5.110.229
123.4.3.0
123.4.3.0
123.41.2.0
123.41.3.1
123.42.3.0
123.45.3.0
255.247.133.43
```

