## wrkspc completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(wrkspc completion zsh); compdef _wrkspc wrkspc

To load completions for every new session, execute once:

#### Linux:

	wrkspc completion zsh > "${fpath[1]}/_wrkspc"

#### macOS:

	wrkspc completion zsh > $(brew --prefix)/share/zsh/site-functions/_wrkspc

You will need to start a new shell for this setup to take effect.


```
wrkspc completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
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
