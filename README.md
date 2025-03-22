# Gonion

The idea of this project is to have a tool that generates project structure
following the Onion/Hexagonal/Ports & Adapters architecture principles.

The tool is configured using YAML config that is then parsed and templated.
The goal is not to create a completely functional project but to generate most
of the boilerplate code and files that you then can go and easily edit.

Apart from generating things like Domain objects, Application services,
API controllers etc., the tool can also bootstrap project files like Makefile,
GolangCI linter configuration, gitignore etc.
