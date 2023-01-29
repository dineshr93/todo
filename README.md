# todo

TODO CLI APP

## Description

This is intended for managing the tasks locally.
By default uses _$HOME/.todo.json_ file to store the data. (you can change this using _--config_ option any time)

## Getting Started

Contains following commands

1. Add
2. Mark Complete
3. Delete
4. List

### Dependencies

- Cobra
- Viper
- Simple table

### Installing

- Add the binary to your environment path and use it

### Executing program

- How to run the program

```
COMMAND FORMAT:
todo add "project - task - jiraLink"
todo add "YOUR_PROJECT_NAME - TASK_NAME - http://www.yourcompany.jira.com"

todo complete TASKID
<!-- you can also mark multiple tasks to mark complete by passing space separated ids  -->
todo complete 1 2 ..

todo pending TASKID
<!-- you can also mark multiple tasks back to mark pending by passing space separated ids  -->
todo pending 1 2 ..

todo delete TASKID
<!-- you can also mark multiple tasks for deletion by passing space separated ids  -->
todo delete 1 2 ..
```

## Help

any issues please raise an bug reguest

```
>todo -h
TODO CLI APP

1. Add
2. Mark Complete
3. Delete
4. List

Usage:
  todo [command]

Available Commands:
  add         add a todo Format: Project - task - JIRA link
  complete    Mark a task as completed
  delete      Delete a task from the list
  help        Help about any command
  ls          List all todos
  pending     Mark a task as pending

Flags:
      --config string   config file (default is $HOME/.tdv.json)
  -h, --help            help for tdv
  -t, --toggle          Help message for toggle

Use "todo [command] --help" for more information about a command.

```

## Authors

Dinesh Ravi

## Version History

- 1.0.0
  - Initial Release

## License

This project is licensed under the Apache License 2.0 - see the [Apache-2.0](LICENSE) file for details

## Acknowledgments

- [cobra](https://www.github.com/spf13/cobra)
- [viper](https://www.github.com/spf13/viper)
- [simpletable](https://www.github.com/alexeyco/simpletable)
- [joefazee](https://www.github.com/joefazee)