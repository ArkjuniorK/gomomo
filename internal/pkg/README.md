### Package as Business Logic
Each package should contain business logic and consist of `service`, `repository`, `model` and main file
which use to initialize the package. If somehow a package is growing too big, 
then breaking each file into sub package could also be done to help 
reduce complexity and manage the code as long as main file 
initiate and return the interface of Service which 
would be used by application API.