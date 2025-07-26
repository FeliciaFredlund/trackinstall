# Track Installs

A CLI to track manually installed programs/apps, their dependencies, and their install/uninstall instructions.

This tool solves the problem of keeping track of when manually installed packages in your package manager are no longer needed. This tool allows you to track why you manually installed a package (aka dependency). When you tell the tool you're uninstalling a program it can tell you which dependencies that none of your programs use anymore, so you can safely move them to auto (installed) and your package manager can now tell you if it is safe to remove or not.

This is to avoid having your OS slowly bloat by not removing dependencies that nothing uses anymore. A package manager can only do that for dependencies installed by it, not when you manually install them through the package manager.

## Usage

TO BE ADDED: Instructions on how to install/run project

ADD SCREENSHOTS

ADD AN EXAMPLE FOLDER WITH A DATA FILE (SO IT IS EASIER TO TEST) AND AN INSTALLATION SCRIPT THAT CAN BE IMPORTED

## About this project

This project was started and an intial MVP (I hope!) was completed during a [boot.dev](https://boot.dev) hackathon in July 2025. I created it so I could track manual installations on my own machines to ensure that my OS doesn't slowly bloat.

## Future plans

- Add the option to pick the path and filename for the data file. This should also allow placing the binary file wherever. Option would be to add a config file (with the file path) or ???
- While adding a project, if a name is missing or invalid, change it to a for loop so they can keep trying picking a name. Same name as another program should as if they want to overwrite the existing one.
- Also add to edit project, to change missing/invalid name to a for loop, and the option to overwrite if another program with the same name exists
- Add what package manager is used (needs some way to save that, config file like filepath?), and have the tool give the commands for installing the dependencies and removing/changing them to auto when no longer needed.
- List dependencies and what programs are dependent on them.

## Change log

TO BE ADDED when MVP is done