## wrkspc completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(wrkspc completion bash)

To load completions for every new session, execute once:

#### Linux:

	wrkspc completion bash > /etc/bash_completion.d/wrkspc

#### macOS:

	wrkspc completion bash > $(brew --prefix)/etc/bash_completion.d/wrkspc

You will need to start a new shell for this setup to take effect.


```
wrkspc completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --config string    config file (default is $HOME/.wrkspc.yml) (default ".wrkspc.yml")
  -p, --port string      The port the service should run on. (default "8090")
  -s, --service string   The service that should be constructed. (default "gateway")
      --viper            use Viper for configuration (default true)
```

### SEE ALSO

* [wrkspc completion](wrkspc_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 16-Oct-2022
