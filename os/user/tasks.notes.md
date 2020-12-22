# Tasks For Drop-in User Library Replacement
A replacement for the standard library `user` library (Multiverse OS will be 
providing a complete `os` package, and all subpackages). The primary motivation
is that the current standard libraries make signficant preformance sacrafices on
Unix machines for the benefit of preformance increase on Windows machines (the
use of string to hold UID, and GID is a primary example). 

## General Tasks 

  * Get rid of any and all string fields, when an int should/could be used; for
    example with `User.Gid` and `User.Uid`. 
  * Extend the functionality of the `user` library to provide more complete
    functionality after focusing solely on Unix/Linux support, dropping support
    for Windows entirely. 
    * Functionality to check the password
    * Functionality to generate, encrypt, sign, and manipulate SSH
      configurations files and key pairs. 
    * Load all configuration values defined/represented as files and
      configuration files in `.config` or `.local` or `.{{APPLICATION NAME}}`.
      After loading all of this data into a consistent and straight forward
      key/value store. Then make it easy to modify these values by changing the
      values and writing the changes. 
    * Build a new configuration and local data storage as defined in the
      Multiverse OS specification. 
    * Make it simple to backup all the user settings into an auto-extracting
      archive that can be easily deployed on a new machine 
    * Build a new virtual filesystem showing only the files and folders
      accessible to the user. 
    * Combine all configurations defined as default by `/etc/*`, then show which
      of those are overriden by the files in the home folder. 
    * Build a home folder based bin folder and add it to the path. Make it
      simple to add folders to path. Divide binaries available by each path.
      Show the system executables, and the user executables. 
    * Track all connections, processes, memory, processor, and similar
      information used by the 
