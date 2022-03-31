<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">

▒█▀▀█ ▒█░░░ ▀█▀ <br/>
▒█░░░ ▒█░░░ ▒█░ <br/>
▒█▄▄█ ▒█▄▄█ ▄█▄ <br/>

**URL** [multiverse-os.org](https://multiverse-os.org)
The `cli` framework aims to provide a security focused, and easy-to-use
toolbox for creating command-line interfaces for simple scripts, to full
featured TUI applications. Not just the standard command-processor model
(commands, flags, params) but also shell interfaces.

This framework also seeks to establish precedent that `cli` frameworks 
should not just provide a help output, and register commands and flags
but provide the tools necessary for not jsut secure user input but rich
user input, high-quality and customizable ascii/text generation for
tables, banners (using figlet fonts), sparkline graphs and more. Loading
bars and spinners that are easy to customize and full featured. 

#### CLI Framework
**Features** 
As this software is in its pre-alpha stages, not all the features below are
completed, some are complete, some are in-progress, and some are in planning
stages. 

  * **Full VT100 support** providing ANSI coloring and styling through several 
    sub-packages providing different levels of sophistication to provide
    functionality for simple scripts with little overhead, or robust full CLI
    applications with full SGR/CSI functionality with helpers, grid system, and
    other features required for complete TUI applications.  
  
  * **Sophisticated user input** including secure password input, list/menu, 
    multiselect, shell, and input validaiton for all basic types.

  * **Command-line interfaces with commands, subcommands, flags, and params**
    with full support for stacked flags, flag param separation using both " "
    and "=" for maximum compatibility.  

  * **Loading Bars & Spinners** easy to customize, and included as indepedent
    subpackages for minimal overhead. 

  * **ASCII/Text helpers** in the form of *Tables*, *Graphs/Histograms*, 
    *QR Codes*, *Banners* (using figlet fonts), symbol sets (using unicode) for 
    a variety of purposes. 

  * **Localization support**
