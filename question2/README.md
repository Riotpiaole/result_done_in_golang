# Additional Note on the bug where open lead to an unexpected output 
- During the trouble shooting on the open function issue, we spent too much tme investigating the why is the solution from google search wasn't applicable. This raise the concern that we need to enforce some coding practice when coding. 
  - The issue was reproduced in [explain.py](https://github.com/Riotpiaole/result_done_in_golang/blob/main/question2/root/devops/explain.py) 

> And the output is
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

- It ticks me on we might be using a wrong functions from the error message ( `'str' object cannot be interpreted as an integer`) after Ben print the program execute prior on `open`.

- So i suspect the function was used is coming from `os` package and override it similar to [explain.py line8](https://github.com/Riotpiaole/result_done_in_golang/blob/main/question2/root/devops/explain.py#L8)

- To reinforce my assumption, i ran both function with `help` and found them are completely different
> `help(open)`
```
open(file, mode='r', buffering=-1, encoding=None, errors=None, newline=None, closefd=True, opener=None)
    Open file and return a stream.  Raise OSError upon failure.

    file is either a text or byte string giving the name (and the path
    if the file isn't in the current working directory) of the file to
    be opened or an integer file descriptor of the file to be
    wrapped. (If a file descriptor is given, it is closed when the
    returned I/O object is closed, unless closefd is set to False.)
...
```

> `help(os.open)`
```
open(path, flags, mode=511, *, dir_fd=None)
    Open a file for low level IO.  Returns a file descriptor (integer).

    If dir_fd is not None, it should be a file descriptor open to a directory,
      and path should be relative; path will then be relative to that directory.
    dir_fd may not be implemented on your platform.
      If it is unavailable, using it will raise a NotImplementedError.
```

## Open vs os.open

- By default the `open` is a [high level API](https://docs.python.org/3/library/functions.html#open) which returns the file object with a stream of data. 

- The bug is caused by importing the [low level API](https://docs.python.org/3/library/os.html#os.open) from os and cause the confusion during trouble shooting. 


## Lession Learned
- We should always import the exact methods we are using that override defined functions.

- In addition, always distinguish the functions name using `Polymorphism` or Name to avoid unexpected blocker during the code.

# Solutions output for question2
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

