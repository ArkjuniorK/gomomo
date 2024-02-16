This directory contains database models that would be shared across the moduels/packages. The main reason why models are defined in seperate package
and not inside the module to avoid duplications, for example if somehow two modules require a same model then it would be hassle to duplicate
the model and make sure the definition of model on each module is identical. The most common use case is for `auth` module, where this
module need to access the user model that might lived at another module.

Thus, the decision to move the model outside the module would introduces coupling between each modules, but since this project is a monorepo
it would be a good trade-off.
