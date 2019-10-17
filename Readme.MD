1. Use a CLI (command line interface) that allows any YAML file and output that data in JSON.
For this program we need to use go run fileparser.go -config=employee.yaml to pass the input
output will be displayed in the console.

2. The program will not need a REST API and should use libraries specific for yaml and json
I have not created any Rest API.

3. Determine how large of a YAML file can be given to the program before it has an issue.
I have gone through the library documentation. I couldn't find the capacity.

4. How can different schema versions be used with this program
Added go modules I hope that is what expected from this problem.

5. The unit tests will only validate that the output given yaml input is formatted correctly."

I am not able to write unit tests this time since I am not able to use structs to unmarshal for the given program.

