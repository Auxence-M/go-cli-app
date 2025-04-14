## This a small cli app project made in golang

The app name is doli. doli is a todo list cli app made in golang using the [spf13 cobra package](https://github.com/spf13/cobra)

## How to run

1. Modify the _.doli.yaml_ file and provide a path for where you would like to store the datafile for this app. 
2. Run `go build` to build the app, a _doli.exe_ file should appear in the directory 
3. You can now start using the app by using the following `./doli [command] --flag`

## List of commands with examples

### add command

```
./doli add "todo 1" "todo 2" "todo 3"
./doli add "todo 1" -p1
./doli add "todo 1" --priority3
./doli add --help
```

### list command

By default, the list command only show items in the todo list that are not yet set to done

```
./doli list
./doli list --all
./doli list --done
./doli list -p2
./doli list -p2 --all
./doli list -p3 --done
```

### edit command

```
./doli edit [todo number] [new todo value] [new todo priority]
./doli edit 2 "todo 5" -p3
./doli edit 1 -p1
```
### done command

```
./doli done [todo number] 
./doli done 2 
```