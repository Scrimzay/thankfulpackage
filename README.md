ever wanted to thank the packages you use because you use em so much but too lazy to write em all down cause theres so much cause youre a package noob? this make it easy, just import it and call it like so:

```
package main

import (
	"fmt"
	"github.com/Scrimzay/thankfulpackage"
	"os"
)

func main() {
	dir := "." // Current directory
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	err := githubthanks.GenerateThanks(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("README.md generated successfully")
}
```
