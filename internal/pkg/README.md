Package that act as module would live at this directory, each package would contain
at least 5 directory and a one main file named after the package name. Those directory is `/routes`, `/services`, `/repositories`, `/handlers`, `/models`.
It's not limited to these dirs so dev could add new directory if required, for example if the
package/module had too many middleware, then it possible to create new directory called `/middlewares` to
hold the middlewares.

One thing to note that only code from main file that would be exported and used outside the package.
So code from `/services` shouldn't be used outside the package. 