# Testing Command Line Applications

To test a command line application you can use an expect script.

- https://github.com/Netflix/go-expect
- https://github.com/google/goexpect

It runs an external program and compares the output with a set of expected output.
For this it created a virtual PTY and runs the program in it. The PTY is used to capture the output of the program.

To also capture the FS we could use a virtual FS (https://github.com/spf13/afero), which then would provide the input and output files for the program.
The other option would be to use a read only part of local FS and a temp dir for the output.

Another option would be to use directly the cobra commands without launching an external program. This would allow us to use the same tests for the CLI and the API.
