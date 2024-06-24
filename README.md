# Dragonfly Classic Commands
A WIP implementation of the classic Minecraft commands


## Usage:

Put this line in your `main()` function before the `srv.Listen()` line:

```
  classicCommands.Init(srv)
```

And this on the end of the `main()` function:

```
  classicCommands.Save()
```