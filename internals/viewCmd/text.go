package viewCmd

const Help = `
Purpose:
This is a simple command line application that allows running multiple commands from a single name defined in commands.txt this can include multiple or chained commands.
Usage:
Now the most important part is the commands.txt file, defines the commands that can be run, the format is as follows: 'test: | ls | echo hi |' as you can see everything before the first colon is
the command name, everything between |  | is the commands that is runned on the name call every command must be separated by a pipe with a space before and after it, the command name can be anything you want but must be unique.
Special commands:
Special commands are commands that are in the specialCommands.txt file, also shown in the help special commands tab. What these do are commands that are supported by the application itself and they simplfie the commands and organization fo the related information. 
This can of course be ignored by the user if they want by simply running the command as is, but it is recommended to use the special commands as they are more organized and easier to use for most users.
`