# Overview
This is a CLI tool to search all bookmarks with a keyword over user Profiles in Google Chrome.\
Now for Windows only.\
Using [Bubbletea](https://github.com/charmbracelet/bubbletea), thank you!.

# Usage
1. Write your user name in settings_example.json. In Windows11, it's the first 5 character of your user name. Please check __C:\Users:\Your username__
2. Change a file name from ```settings_example.json``` to ```settings.json```.
2. Run this command, or compile and run exe file.
```
go run main.go
```
4. Put a keyword you want to search, and enter. Ctrl + c terminate the program.

# Testing
## util package
Before ```go test```, you should set username as environmental variable.\
```$Env:UserNameForTest="username"```\
And put settings.json you want to use in the same folder.

# Valid version
Windows : 11\
Chrome :  129.0.6668.90\
Go : 1.22.3
