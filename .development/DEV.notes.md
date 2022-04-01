# Developer Notes
Key details tracking the progress to first release candidate and freezing of the
API. We are focusing on a simple CLI program first but eventually should provide
everything we need for any CLI interaction for all Multiverse software. 




## Functionality Missing
Core functionality that is considered to be requried for the first release
candidate to be ready, and for the API to freeze. 

  * Need to add configuration via JSON, YAML, TOML, and XML. Then support using
    environmental viariables. 
  * Basic validations 
  * Ability to initialize config/local-data folders; then build default data in
    each one. Including a PID file 
  * Support daemonization
  * Support console 
  * Add ability to write to log files 
  * Merge terminal code for a consistent TUI style system with no dependencies
    and small as possible. 
  * Support a **few** specialized types that are common, like IP adddress, Port
    Number, Web URL, File path. 
  * build default help documentation served as a [webpage](https://github.com/mwitkow/go-flagz), but also produce man page, and other standard
  * build debian package
  * sign builds, support multisig

At the end we should be able to replicate the INSTALL and RESCUE modes provided
by a debian installation disk. 

