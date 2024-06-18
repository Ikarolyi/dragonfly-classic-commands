# Dragonfly Classic Commands
A WIP implementation of the classic Minecraft commands


## Usage:

Put this line in your `main()` function before the `srv.Listen()` line:

```
classicCommands.Init()
```

And have this in your accept handler function:

```
for srv.Accept(func(p *player.Player) {<br/>
  // Your code<br/>
  classicCommands.PassAccept(p)<br/>
  // Also your code<br/>
}) {}<br/>
```