

/cmd
Main applications for this project.

The directory name for each application should match the name of the executable you want to have (e.g., /cmd/myapp).

Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects, then it should live in the /pkg directory. If the code is not reusable or if you don't want others to reuse it, put that code in the /internal directory. 



/internal
Private application and library code. This is the code you don't want others importing in their applications or libraries. Note that this layout pattern is enforced by the Go compiler itself. 


/pkg
Library code that's ok to use by external applications (e.g., /pkg/mypubliclib).


/vendor
Application dependencies (managed manually or by your favorite dependency management tool like the new built-in Go Modules feature).
The go mod vendor command will create the /vendor directory for you.

Note: Don't commit your application dependencies if you are building a library.


/api
OpenAPI/Swagger specs, JSON schema files, protocol definition files.